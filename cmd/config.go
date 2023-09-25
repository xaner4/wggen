package cmd

import (
	"fmt"
	"os"

	"github.com/mdp/qrterminal/v3"
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

var cmdCfgPeerQR = &cobra.Command{
	Use:   "qr",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		warmUp()
		qrPeer()
	},
}

func init() {
	cmdConfig.AddCommand(cmdCfgSrv)
	cmdConfig.AddCommand(cmdCfgPeer)
	cmdConfig.AddCommand(cmdCfgPeerQR)

	cmdCfgPeer.Flags().StringVarP(&cfgname, "name", "n", "", "Name of the peer you want to show")
	cmdCfgPeerQR.Flags().StringVarP(&cfgname, "name", "n", "", "Name of the peer you want to show")
}

func cfgSrv() {
	cfg, err := srv.GenerateSrvConfig()
	if err != nil {
		fmt.Printf("Error generating server config: %v\n", err)
	}
	fmt.Println(cfg)
}

func cfgPeer() {
	if cfgname == "" {
		fmt.Println("No peer specified")
		os.Exit(1)
	}
	cfg, err := srv.GeneratePeerConfig(cfgname)
	if err != nil {
		fmt.Printf("Error generating peer config: %v\n", err)
	}
	fmt.Println(cfg)
}

func qrPeer() {
	if cfgname == "" {
		fmt.Println("No peer specified")
		os.Exit(1)
	}
	cfg, err := srv.GeneratePeerConfig(cfgname)
	if err != nil {
		fmt.Printf("Error generating peer config: %v\n", err)
	}

	config := qrterminal.Config{

		Level:     qrterminal.L,
		Writer:    os.Stdout,
		BlackChar: qrterminal.BLACK,
		WhiteChar: qrterminal.WHITE,
		QuietZone: 1,
	}

	qrterminal.GenerateWithConfig(cfg, config)
}
