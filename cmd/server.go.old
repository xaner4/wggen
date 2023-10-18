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

var cmdServer = &cobra.Command{
	Use:   "server",
	Short: "Generate a new server config instance",
	Long:  `Server will generate a Yaml configuration for the server instance of Wireguard. This will generate new public and private keys for the server instance automatically`,
	Run:   func(cmd *cobra.Command, args []string) { servers() },
}

func init() {
	cmdServer.Flags().StringVarP(&srvname, "name", "n", "wg0", "Name of the server instance")
	cmdServer.Flags().IntVarP(&port, "port", "p", 51820, "Listening port for the server instance")
	cmdServer.Flags().StringVarP(&subnets, "subnet", "s", "172.16.16.1/24", "Comma-separated list of IP ranges that clients will connect from")

}

func servers() {

	// Split the comma-separated list of subnets into individual subnets
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
