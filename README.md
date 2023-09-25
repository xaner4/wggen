# WGGEN - Wireguard Configuration Manager and Generator

WGGEN is a powerful tool designed to simplify the management and generation of Wireguard configurations. It allows you to effortlessly set up and manage your Wireguard VPN network, providing an easy-to-use command-line interface and keeping track of key pairs in one central location.

## Features
* Generate Wireguard server and client configurations with ease
* Manage key pairs and IP addresses for each client securely
* Cross-platform compatibility (Linux, macOS, Windows)
* Customizable template-based configuration generation
* Intuitive command-line interface for seamless configuration management

## Installation

To install WGGEN, you need to have Go installed on your system. You can install WGGEN using the following command:

```bash
go install github.com/xaner4/wggen
```

Make sure you have your Go workspace bin directory added to your system's PATH for easy access to the wggen command.
## Usage


WGGEN offers several subcommands that enable you to perform different tasks. Here are the available subcommands:
### server

The server subcommand allows you to generate a new server configuration for your Wireguard instance. By running the following command, you can create a server configuration:

```bash
wggen server --name=<name> --endpoint=<endpoint> --port=<port> --subnet=<subnet>
```

* \<name>: Name of the server instance (default: "wg0")
* \<endpoint>: IP or DNS name for the server instance
* \<port>: Listening port for the server instance (default: 51820)
* \<subnet>: Comma-separated list of IP ranges that clients will connect from (default: "172.16.16.1/24")

### peer

The peer subcommand allows you to manage peers in your Wireguard network. It provides two subcommands:

* add: Add a new peer to the Wireguard network.
* del: Delete a peer from the Wireguard network.

To add a new peer, use the following command:

```bash
wggen peer add --name=<name> --server=<server_endpoint> --allowed-ips=<allowed_ips>
```

To delete a peer, use the following command:

```bash
wggen peer del --name=<name> --server=<server_endpoint>
```

* \<name>: Name of the peer
* \<server_endpoint>: IP or DNS name of the server instance to which the peer belongs
* \<allowed_ips>: Comma-separated list of allowed IP addresses for the peer

### config

The config subcommand enables you to generate specific wireguard configurations and print them to stdout. It provides two subcommands:

* server: Generates the server Wireguard configuration.
* peer: Generates the peer Wireguard configuration.

To generate the server configuration and print it to stdout, use:

```bash
wggen config server
```

To generate the peer configuration and print it to stdout, use:

```bash
wggen config peer --name=<name> --server=<server_endpoint>
```
