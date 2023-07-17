package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/xaner4/wggen/wggen"
)

var dir string
var endpoint string

var rootCmd = &cobra.Command{
	Use:   "wgconf",
	Short: "Easy Wireguard Configuration generation",
	Long:  `wgconf is a tool for generating Wireguard Configuration and keeping track of key pairs in one place`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&dir, "path", "P", "wg-config", "Directory of the config files")
	rootCmd.PersistentFlags().StringVarP(&endpoint, "endpoint", "e", "", "IP or DNS name for the server instance")
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
	}
}
