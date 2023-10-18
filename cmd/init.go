package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
	wggen "github.com/xaner4/wggen/wggen"
)

var srvname string
var port int
var subnets string

var cmdInit = &cobra.Command{
	Use:   "init [endpoint]",
	Short: "Initialize a new wireguard server",
	Long:  `Automatically initialize a new wireguard server with the given endpoint and a private public key pair.`,
	Args:  cobra.ExactArgs(1),
	Run:   func(cmd *cobra.Command, args []string) { initwg(args[0]) },
}

func init() {
	cmdInit.Flags().StringVarP(&srvname, "name", "n", "wg0", "Name of the server instance")
	cmdInit.Flags().IntVarP(&port, "port", "p", 51820, "Listening port for the server instance")
	cmdInit.Flags().StringVarP(&subnets, "subnet", "s", "172.16.16.1/24", "Comma-separated list of IP ranges that clients will connect from")

}

func initwg(endpoint string) {
	// Split a comma-separated list of subnets into individual subnets
	subnetList := strings.Split(subnets, ",")

	srv, err := wggen.GenServerConf(srvname, endpoint, port, subnetList)
	if err != nil {
		fmt.Printf("Something went wrong with generating server config\nError: %v\n", err)
		os.Exit(1)
	}

	err = srv.SaveWGConfig(dir)
	if err != nil {
		fmt.Printf("Something went wrong with saving the server config\nError: %v\n", err)
		os.Exit(1)
	}

	log.Printf("Server %s was succesfully generated and saved to %s\n", endpoint, dir)
}
