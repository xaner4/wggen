package cmd

import (
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
	_ = srv.GeneratePeerConfig(cfgname)
}

func cfgSrv() {
	_ = srv.GenerateSrvConfig()
}
