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

	"github.com/4rchr4y/bpm/iostream"
	"github.com/4rchr4y/godevkit/v3/env"
	"github.com/4rchr4y/godevkit/v3/syswrap"
	dummy_component "github.com/4rchr4y/goray/example/dummy-component"
	noop_driver "github.com/4rchr4y/goray/example/noop-driver"
	"github.com/4rchr4y/goray/exp"
	"github.com/4rchr4y/goray/interface/component"
	"github.com/4rchr4y/goray/interface/driver"

	"github.com/4rchr4y/goray/internal/bundlepkg"
	"github.com/4rchr4y/goray/internal/config/confload"
	"github.com/4rchr4y/goray/internal/grpcplugin"
	"github.com/4rchr4y/goray/internal/grpcplugin/detection"
	"github.com/4rchr4y/goray/internal/grpcwrap"
	"github.com/4rchr4y/goray/internal/hclutl"
	"github.com/4rchr4y/goray/internal/kernel"
	"github.com/4rchr4y/goray/internal/proto/protocomponent"
	"github.com/4rchr4y/goray/internal/proto/protodriver"

	"github.com/g10z3r/ason"
	pluginHCL "github.com/hashicorp/go-plugin"
	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/rego"
	"github.com/open-policy-agent/opa/topdown"
	"github.com/spf13/afero"
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
	bpm := bundlepkg.NewBundlePkgManager(bundlepkg.BundlePkgManagerConf{
		RootDir:  env.MustGetString("BPM_PATH"),
		OSWrap:   new(syswrap.OSWrap),
		IOWrap:   new(syswrap.IOWrap),
		IOStream: iostream.NewIOStream(),
	})
	linker := bundlepkg.NewBundlePkgLinker(bundlepkg.LinkerConf{
		Fetcher:    bpm.Fetcher(),
		Inspector:  bpm.Inspector(),
		Manifester: bpm.Manifester(),
	})

	parser := confload.NewParser(afero.OsFs{})
	loader := confload.NewLoader(parser)

	conf, diags := loader.LoadConf(".ray")
	fmt.Println(diags)
	if diags.HasErrors() {
		fmt.Fprintf(os.Stderr, "errors encountered while parsing HCL file: %s", diags.Error())
		return
	}

	app := &exp.RayApp{
		Drivers: processDriversList(detection.FindDriverPaths(".ray/driver")),
		Config:  conf,
	}

	confParams := exp.ConfParams{
		Scope: hclutl.NewScope(),
	}

	scope := hclutl.NewScope()
	evaluator := kernel.NewEvaluate(conf, scope)
	s, diags := evaluator.ExpandVariables()
	fmt.Println(diags)
	if diags.HasErrors() {
		panic(diags.Error())
	}
	confParams.Modules = s.Modules

	app.Configure(&confParams)

	b, err := bpm.Storage().LoadFromAbs("./testdata", nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	modules, err := linker.Link(context.Background(), b)
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

func processDriversList(paths []string) map[string]component.Interface {
	result := make(map[string]component.Interface, len(paths))

	result["dummy"] = startComponent(".ray/component/github.com/4rchr4y/ray-dummy-component@v0.0.1")

	return result
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
		SyncStdout:       os.Stdout,
		Stderr:           os.Stderr,
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
		SyncStdout:       os.Stdout,
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
