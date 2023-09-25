package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/xaner4/wggen/wggen"
)

var dir string
var endpoint string
var version string = "0.0.1"

var rootCmd = &cobra.Command{
	Use:   "wggen",
	Short: "Easy Wireguard Configuration generation",
	Long:  `wggen is a tool for generating Wireguard Configuration and keeping track of key pairs in one place`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print current version and exit",
	Long:  `Print current version and exit`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("wggen version: %s\n", version)
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		flags := cmd.InheritedFlags()
		flags.SetAnnotation("endpoint", cobra.BashCompOneRequiredFlag, []string{"false"})
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&dir, "path", "P", "wg-config", "Directory of the config files")
	rootCmd.PersistentFlags().StringVarP(&endpoint, "endpoint", "e", "", "IP or DNS name for the server instance")
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(cmdServer)
	rootCmd.AddCommand(cmdPeer)
	rootCmd.AddCommand(cmdConfig)

	rootCmd.MarkPersistentFlagRequired("endpoint")
}

func Execute() error {
	return rootCmd.Execute()
}

func warmUp() {
	var err error
	srv, err = wggen.GetWGConfig(dir, endpoint)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
