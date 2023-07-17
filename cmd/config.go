package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cfgname string

var cmdConfig = &cobra.Command{
	Use:   "config",
	Short: "",
	Long:  ``,
	Run:   func(cmd *cobra.Command, args []string) { cmd.Help() },
}

var cmdCfgSrv = &cobra.Command{
	Use:   "server",
	Short: "Generates server wireguard configuration and print to stdout",
	Long:  `config server will generate the wireguard configuration put on the server, it will print the complete configuration with clients configured`,
	Run: func(cmd *cobra.Command, args []string) {
		warmUp()
		cfgSrv()
	},
}

var cmdCfgPeer = &cobra.Command{
	Use:   "peer",
	Short: "Generates peer wireguard configuration and print to stdout",
	Long:  `config peer will generate the wireguard configuration that is used by the peer's`,
	Run: func(cmd *cobra.Command, args []string) {
		warmUp()
		cfgPeer()
	},
}

func init() {
	cmdConfig.AddCommand(cmdCfgSrv)
	cmdConfig.AddCommand(cmdCfgPeer)

	cmdCfgPeer.Flags().StringVarP(&cfgname, "name", "n", "", "Name of the peer you want to show")
}

func cfgPeer() {
	cfg, err := srv.GeneratePeerConfig(cfgname)
	if err != nil {
		fmt.Printf("Error generating peer config: %v\n", err)
	}
	fmt.Println(cfg)
}

func cfgSrv() {
	cfg, err := srv.GenerateSrvConfig()
	if err != nil {
		fmt.Printf("Error generating server config: %v\n", err)
	}
	fmt.Println(cfg)
}
