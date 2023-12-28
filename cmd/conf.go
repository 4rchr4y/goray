package main

import (
	"fmt"
	"log"

	"github.com/4rchr4y/goray/rayfile"
	"github.com/spf13/cobra"
)

var confCmd = &cobra.Command{
	Use:   "conf",
	Short: "",
	Long:  "",
	Run:   runConfCmd,
}

func init() {
	rootCmd.AddCommand(confCmd)
}

func runConfCmd(cmd *cobra.Command, args []string) {
	const filepath = "./testdata/test_conf.toml"

	conf, err := rayfile.NewConfigFromFile(filepath)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Printf("%#v", conf.Workspace)
}
