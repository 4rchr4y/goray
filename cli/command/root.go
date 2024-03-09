package command

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"go/parser"
	"go/token"
	"log"
	"os"
	"os/exec"

	"github.com/4rchr4y/bpm/bundleutil/encode"
	"github.com/4rchr4y/bpm/bundleutil/inspect"
	"github.com/4rchr4y/bpm/bundleutil/manifest"
	"github.com/4rchr4y/bpm/fetch"
	"github.com/4rchr4y/bpm/iostream"
	"github.com/4rchr4y/bpm/pkg/linker"
	"github.com/4rchr4y/bpm/storage"
	"github.com/4rchr4y/godevkit/v3/env"
	"github.com/4rchr4y/godevkit/v3/syswrap"
	"github.com/4rchr4y/goray/builtin/rck/rcschema"
	dummy_component "github.com/4rchr4y/goray/example/dummy-component"
	noop_driver "github.com/4rchr4y/goray/example/noop-driver"
	"github.com/4rchr4y/goray/interface/component"
	"github.com/4rchr4y/goray/interface/driver"

	"github.com/4rchr4y/goray/internal/domain/grpcwrap"
	"github.com/4rchr4y/goray/internal/grpcplugin"
	"github.com/4rchr4y/goray/internal/proto/protocomponent"
	"github.com/4rchr4y/goray/internal/proto/protodriver"

	"github.com/g10z3r/ason"
	pluginHCL "github.com/hashicorp/go-plugin"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/rego"
	"github.com/open-policy-agent/opa/topdown"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "goray",
	Short: "",
	Long:  "",
	Run:   runRootCmd,
}

type failCase struct {
	Msg string `json:"msg"`
	Pos int    `json:"pos"`
	Sev string `json:"sev"`
}

func runRootCmd(cmd *cobra.Command, args []string) {
	dir := env.MustGetString("BPM_PATH")

	io := iostream.NewIOStream()

	osWrap := new(syswrap.OSWrap)
	ioWrap := new(syswrap.IOWrap)
	encoder := &encode.Encoder{
		IO: io,
	}

	s := &storage.Storage{
		Dir:     dir,
		IO:      io,
		OSWrap:  osWrap,
		IOWrap:  ioWrap,
		Encoder: encoder,
	}

	content, err := osWrap.ReadFile(".ray/main.ray")
	if err != nil {
		log.Fatal(err)
		return
	}

	parser := hclparse.NewParser()
	file, diags := parser.ParseHCL(content, "filename.hcl")
	if diags.HasErrors() {
		fmt.Fprintf(os.Stderr, "Errors encountered while parsing HCL file: %s", diags.Error())
		return
	}

	f, diags := rcschema.DecodeFile(file.Body)
	for _, d := range diags {
		fmt.Println(d.Summary)
	}
	if f != nil {
		for _, v := range f.Ray.RequiredProviders {
			fmt.Println(v.Source)
		}
	}

	// client := startDriver(".ray/driver/github.com/4rchr4y/ray-noop-driver@v0.0.1")
	client := startComponent(".ray/component/github.com/4rchr4y/ray-dummy-component@v0.0.1")
	defer client.Stop()
	defer client.Shutdown()

	schema := client.DescribeSchema()
	fmt.Println("--------------")
	log.Printf("Received schema: %+v\n", schema.Schema.Root.Description)
	fmt.Println("--------------")

	b, err := s.LoadFromAbs("./testdata", nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	inspector := &inspect.Inspector{
		IO: io,
	}

	fetcher := &fetch.Fetcher{
		IO:        io,
		Storage:   s,
		Inspector: inspector,
		GitHub: &fetch.GithubFetcher{
			IO:      io,
			Client:  nil,
			Encoder: encoder,
		},
	}

	manifester := &manifest.Manifester{
		IO:      io,
		OSWrap:  osWrap,
		Storage: s,
		Encoder: encoder,
		Fetcher: fetcher,
	}

	l := linker.Linker{
		Fetcher:    fetcher,
		Manifester: manifester,
		Inspector:  inspector,
	}

	modules, err := l.Link(context.Background(), b)
	if err != nil {
		log.Fatal(err)
		return
	}

	policies := make(map[string]string)
	for path, f := range modules {
		policies[path] = f.String()
	}

	compiler, err := ast.CompileModulesWithOpt(policies, ast.CompileOpts{
		EnablePrintStatements: true,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	fileMap, err := tmpGetFileAstAsMap("testdata/main.go")
	if err != nil {
		log.Fatal(err)
		return
	}

	var buf bytes.Buffer
	r := rego.New(
		rego.Query("data.testbundle.policy1"),
		rego.Compiler(compiler),
		rego.Input(fileMap),
		rego.EnablePrintStatements(true),
		rego.PrintHook(topdown.NewPrintHook(&buf)),
	)

	rs, err := r.Eval(context.Background())
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, result := range rs {
		for _, r := range result.Expressions {
			fmt.Println(r.Value)
		}
	}

	fmt.Println(buf.String())
}

func tmpGetFileAstAsMap(filePath string) (map[string]interface{}, error) {
	fset := token.NewFileSet()

	f, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	pass := ason.NewSerPass(fset)
	fileAstJson := ason.SerializeFile(pass, f)

	jsonData, err := json.Marshal(fileAstJson)
	if err != nil {
		return nil, err
	}

	var fileMap map[string]interface{}
	if err := json.Unmarshal(jsonData, &fileMap); err != nil {
		return nil, err
	}

	return fileMap, nil
}

func startComponent(pluginPath string) component.Interface {
	plugins := map[int]pluginHCL.PluginSet{
		1: {
			"component": &grpcplugin.GRPCComponentPlugin{
				ServeFn: func() protocomponent.ComponentServer {
					return grpcwrap.ComponentWrapper(dummy_component.Component())
				},
			},
		},
	}

	client := pluginHCL.NewClient(&pluginHCL.ClientConfig{
		HandshakeConfig:  grpcplugin.Handshake,
		VersionedPlugins: plugins,
		Cmd:              exec.Command(pluginPath),
		AllowedProtocols: []pluginHCL.Protocol{pluginHCL.ProtocolGRPC},
		Managed:          true,
	})

	rpcClient, err := client.Client()
	if err != nil {
		log.Fatalf("Failed to connect to plugin: %s", err)
	}

	raw, err := rpcClient.Dispense("component")
	if err != nil {
		log.Fatalf("Failed to dispense plugin: %s", err)
	}

	p := raw.(*grpcplugin.GRPCComponent)
	p.PluginClient = client

	return p
}

func startDriver(pluginPath string) driver.Interface {
	plugins := map[int]pluginHCL.PluginSet{
		1: {
			"driver": &grpcplugin.GRPCDriverPlugin{
				ServeFn: func() protodriver.DriverServer {
					return grpcwrap.DriverWrapper(noop_driver.Driver())
				},
			},
		},
	}

	client := pluginHCL.NewClient(&pluginHCL.ClientConfig{
		HandshakeConfig:  grpcplugin.Handshake,
		VersionedPlugins: plugins,
		Cmd:              exec.Command(pluginPath),
		AllowedProtocols: []pluginHCL.Protocol{pluginHCL.ProtocolGRPC},
		Managed:          true,
	})

	rpcClient, err := client.Client()
	if err != nil {
		log.Fatalf("Failed to connect to plugin: %s", err)
	}

	raw, err := rpcClient.Dispense("driver")
	if err != nil {
		log.Fatalf("Failed to dispense plugin: %s", err)
	}

	p := raw.(*grpcplugin.GRPCDriver)
	p.PluginClient = client

	return p
}
