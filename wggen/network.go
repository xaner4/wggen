package wggen

import (
	"fmt"
	"net/netip"
	"strings"
)

/*
getIP takes in a subnet and a slice of allocated IP addresses and returns a valid unused IP address.

The method loops through the IP addresses in the given subnet and checks if each IP address is not already allocated. It returns the first available IP address that is not present in the allocated slice.

Parameters:

	subnet (string): The subnet in CIDR notation (e.g., "192.168.0.0/24").
	allocated ([]string): A slice of strings representing the allocated IP addresses.

Returns:

	addr (string): The available IP address within the subnet.
	err (error): An error object if no available IP address is found or if there is an error in the provided subnet.

Example usage:
subnet := "192.168.0.0/24"
allocated := []string{"192.168.0.2", "192.168.0.5", "192.168.0.10"}
ip, err := getIP(subnet, allocated)
if err != nil {
// Handle error
}
fmt.Println("Available IP:", ip)
*/
func getIP(subnet string, allocated []string) (addr string, err error) {
	IPSubnet, err := netip.ParsePrefix(subnet)
	if err != nil {
		return "", fmt.Errorf("invalid subnet: %s, error %v", subnet, err)
	}

	IPSubnet = IPSubnet.Masked()

	for IP := IPSubnet.Addr().Next(); IPSubnet.Contains(IP); IP = IP.Next() {
		if stringInSlice(allocated, IP.String()) {
			continue
		}
		addr = IP.String()
		break
	}

	if addr == "" {
		return "", fmt.Errorf("no available IP address found")
	}

	if err = ValidateIP(addr, subnet); err != nil {
		return "", err
	}
	return addr, nil
}

/*
GetGateway takes in a subnet and returns the first IP address after the network address.
The function expects the subnet to be in CIDR notation (e.g., "192.168.0.0/24") and returns the IP address immediately following the network address.

Parameters:

	subnet (string): The subnet in CIDR notation (e.g., "192.168.0.0/24").

Returns:

	    gateway (string): The IP address of the gateway.
		err (error): An error object if the provided subnet is not a valid.

Example usage:
subnet := "192.168.0.0/24"
gateway := GetGateway(subnet)
fmt.Println("Gateway IP:", gateway)
*/
func GetGateway(subnet string) (string, error) {
	IP, err := netip.ParsePrefix(subnet)
	if err != nil {
		return "", fmt.Errorf("invalid subnet: %s, error %v", subnet, err)
	}

	IP = IP.Masked()
	addr := IP.Addr()
	if !IP.Contains(addr.Next()) {
		return "", fmt.Errorf("invalid subnet: %s, error %v", subnet, err)
	}
	return addr.Next().String(), nil

}

/*
ValidateIP takes an IP address and the subnet that the IP address belongs to and performs various checks to validate the IP address.

Parameters:
- IP (string): The IP address to validate.
- subnet (string): The subnet in CIDR notation to which the IP address belongs.

Returns:
- err (error): An error object if the IP address is not valid or does not belong to the specified subnet.

Description:
The ValidateIP function checks the validity of the provided IP address within the context of the given subnet. It performs the following checks:

1. Empty IP or Subnet: The function checks if either the IP or subnet is empty. If either value is empty, it returns an error indicating the missing value.

2. Invalid Subnet: The function checks if the provided subnet is valid by parsing it using the netip.ParsePrefix function. If the subnet is invalid, an error is returned.

3. Network Address: The function compares the IP address with the network address of the subnet. If they match, it implies that the IP address is the network address and is considered invalid.

4. Subnet Containment: The function verifies if the IP address falls within the specified subnet by checking if it is contained within the network range. If the IP address is not within the subnet, an error is returned.

Example Usage:
err := ValidateIP("192.168.0.10", "192.168.0.0/24")

	if err != nil {
	    fmt.Println("Invalid IP address:", err)
	}

Please note that this function provides basic IP address validation within a subnet and may require additional checks based on specific use cases and requirements.
*/
func ValidateIP(IP string, subnet string) error {
	if IP == "" || subnet == "" {
		return fmt.Errorf("subnet or IP is empty:\nIP: %s\nsubnet: %s", IP, subnet)
	}

	IPPrefix, err := netip.ParsePrefix(subnet)
	IPNet := IPPrefix.Masked()
	if err != nil {
		return fmt.Errorf("invalid subnet: %s, error: %v", subnet, err)
	}

	IPAddr, err := netip.ParseAddr(IP)
	if err != nil {
		return err
	}

	if !IPPrefix.IsValid() {
		return fmt.Errorf("invalid subnet: %s", subnet)
	}

	if IPNet.Addr().String() == IP {
		return fmt.Errorf("IP address is the network address: %s", IP)
	}

	if !IPNet.Contains(IPAddr) {
		return fmt.Errorf("IP address is not within the specified subnet: %s", IP)
	}

	return nil
}

func IsIPv4(address string) bool {
	return strings.Count(address, ":") < 2
}

func IsIPv6(address string) bool {
	return strings.Count(address, ":") >= 2
}

/*
allocatedIPs returns a list of allocated IP addresses from the Wireguard server configuration.

This method loops through the IP addresses used for the Wireguard network, removes the subnet information, and appends the resulting IP addresses to the allocated slice. It takes into account both the IP addresses defined for the server and the IP addresses defined for each peer.

Parameters:

	None

Returns:

	allocated ([]string): A slice containing the allocated IP addresses without the subnet.
	err (error): An error object if any error occurs during the process.

Example usage:
allocated, err := wg.allocatedIPs()
if err != nil {
// Handle error
}

for _, ip := range allocated {
fmt.Println(ip)
}
*/
func (wg *WGSrv) allocatedIPs() (allocated []string, err error) {

	if wg.IPAddress != nil {
		for _, IP := range wg.IPAddress {
			IP = strings.Split(IP, "/")[0]
			allocated = append(allocated, IP)
		}

	}

	if len(wg.Peers) == 0 {
		return allocated, nil
	}

	for i := range wg.Peers {
		if wg.Peers[i].IPAddress != nil {
			for _, IP := range wg.Peers[i].IPAddress {
				IP = strings.Split(IP, "/")[0]
				allocated = append(allocated, IP)
			}
		}
	}

	return allocated, nil
}
