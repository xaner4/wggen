package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/xaner4/wggen/wggen"
)

var peername string
var allowedIPs string
var DNS string
var PresharedKeys bool
var PersistentKeepalives bool
var srv *wggen.WGSrv

var cmdPeer = &cobra.Command{
	Use:   "peer",
	Short: "List peers in a Wireguard network",
	Long:  `peer will `,
	Run:   func(cmd *cobra.Command, args []string) { warmUp(); listPeers() },
}

var cmdPeerAdd = &cobra.Command{
	Use:   "add",
	Short: "Add a new peer to the wireguard network",
	Long:  `peer will generate a Yaml configuration for the peer instance of Wireguard. This will generate new public and private keys for the peer instance automatically`,
	Run:   func(cmd *cobra.Command, args []string) { warmUp(); addPeer() },
}

var cmdPeerDel = &cobra.Command{
	Use:   "del",
	Short: "Del a peer to the wireguard network",
	Long:  `peer will generate a Yaml configuration for the peer instance of Wireguard. This will generate new public and private keys for the peer instance automatically`,
	Run:   func(cmd *cobra.Command, args []string) { warmUp(); delPeer() },
}

func init() {
	cmdPeerAdd.Flags().StringVarP(&peername, "name", "n", "", "Name of the new peer")
	cmdPeerAdd.Flags().StringVarP(&allowedIPs, "allowedip", "a", "", "Allowed IPs that the peer is sending through to wireguard")
	cmdPeerAdd.Flags().StringVarP(&DNS, "DNS", "d", "1.1.1.1", "DNS servers of the peer instance")
	cmdPeerAdd.Flags().BoolVarP(&PresharedKeys, "presharedKeys", "K", true, "Generate new preshared keys for the peer instance")
	cmdPeerAdd.Flags().BoolVarP(&PersistentKeepalives, "persistentKeepalives", "k", true, "Use persistent keepalives for the peer instance (Defaults to 30 sec)")
	cmdPeerAdd.MarkFlagRequired("name")

	cmdPeerDel.Flags().StringVarP(&peername, "name", "n", "", "Name of the peer you want to delete")
	cmdPeerDel.MarkFlagRequired("name")

	cmdPeer.AddCommand(cmdPeerAdd)
	cmdPeer.AddCommand(cmdPeerDel)

}

func listPeers() {
	if len(srv.Peers) == 0 {
		fmt.Println("No peers found")
		return
	}

	fmt.Printf("Peers for %s:\n", endpoint)
	fmt.Println("Name\t\tIP\t\t\t\t\tAllowedIPs")
	fmt.Println("------------------------------------------------------------------------------------------------------------------------------------------------")
	for _, wg := range srv.Peers {
		fmt.Printf("%s\t\t%s\t\t\t%s \n", wg.Name, wg.IPAddress, wg.AllowedIPs)
	}
}

func addPeer() {
	aip := strings.Split(allowedIPs, ",")
	dns := strings.Split(DNS, ",")

	if aip[0] == "" {
		aip = nil
	}
	aip = append(aip, srv.IPAddress...)

	peer, err := srv.GenPeerConf(peername, aip, dns, PresharedKeys, PersistentKeepalives)
	if err != nil {
		fmt.Printf("Error generating peer %v\nError: %v\n", peername, err)
		os.Exit(1)
	}
	srv.Peers = append(srv.Peers, peer)
	err = srv.UpdateWGConfig(dir)
	if err != nil {
		fmt.Printf("Error adding peer %v to %v\nError: %v\n", peername, endpoint, err)
		os.Exit(1)
	}
	fmt.Printf("Peer \"%s\" has been added successfully to %v\n", peername, endpoint)
}

func delPeer() {
	if len(srv.Peers) == 0 {
		fmt.Printf("no peer found")
		return
	}

	// Find the index of the peer with the given name
	index := -1
	for i, peer := range srv.Peers {
		if peer.Name == peername {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Printf("no such peer")
		return
	}

	// Swap the peer to delete with the last peer in the slice
	srv.Peers[index] = srv.Peers[len(srv.Peers)-1]
	srv.Peers = srv.Peers[:len(srv.Peers)-1]

	err := srv.UpdateWGConfig(dir)
	if err != nil {
		fmt.Printf("failed to save server config\nError: %v\n", err)
	}
	fmt.Printf("Peer %s has been deleted\n", peername)
}
