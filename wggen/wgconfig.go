package wggen

import (
	"fmt"
	"net/netip"
	"strings"
)

type WGSrv struct {
	Name       string   `yaml:"name"`
	Endpoint   string   `yaml:"endpoint"`
	ListenPort int      `yaml:"listenPort"`
	IPAddress  []string `yaml:"IPAddress"`
	PrivateKey string   `yaml:"privateKey"`
	PublicKey  string   `yaml:"publicKey"`
	Peers      []Peers  `yaml:"peers"`
}

type Peers struct {
	Name                string   `yaml:"name"`
	IPAddress           []string `yaml:"IPAddress"`
	AllowedIPs          []string `yaml:"allowedIPs"`
	PrivateKey          string   `yaml:"privateKey"`
	PublicKey           string   `yaml:"publicKey"`
	PresharedKey        string   `yaml:"presharedKey"`
	PersistentKeepalive int      `yaml:"persistentKeepalive"`
	DNS                 []string `yaml:"dns"`
}

func GenServerConf(Name string, endpoint string, listen int, IPSubnet []string) (serverConf WGSrv, err error) {
	if listen == 0 {
		listen = 51820
	}

	ips := make([]string, 0)
	subnet := []string{}
	for _, v := range IPSubnet {
		ip, err := GetGateway(v)

		if err != nil {
			return WGSrv{}, err
		}
		a := strings.Split(v, "/")
		a[0] = ip
		ip = strings.Join(a, "/")
		ips = append(ips, ip)

		subnet = append(subnet, a[1])
		if len(subnet) > 0 {
			if subnet[0] != a[1] {
				return WGSrv{}, fmt.Errorf("missmatch in subnet lenght between networks. Change to same subnets lengths")
			}
		}
	}

	priv, err := genPrivKey()
	if err != nil {
		return WGSrv{}, err
	}

	pub, err := genPubKey(priv)
	if err != nil {
		return WGSrv{}, err
	}

	serverConf = WGSrv{Name: Name, Endpoint: endpoint, IPAddress: ips, ListenPort: listen, PrivateKey: priv, PublicKey: pub}

	return serverConf, err
}

func (wg *WGSrv) GenPeerConf(Name string, AllowedIPs []string, DNS []string, PresharedKey bool, PersistentKeepalive bool) (Peers, error) {
	var psk string
	var pkl int
	var err error

	for _, v := range wg.Peers {
		if v.Name == Name {
			return Peers{}, fmt.Errorf("peer with name \"%s\" already exists", Name)
		}
	}

	if PresharedKey {
		psk, err = genPSK()
		if err != nil {
			return Peers{}, err
		}
	}

	if PersistentKeepalive {
		pkl = 30
	}
	a, err := wg.allocatedIPs()
	if err != nil {
		return Peers{}, err
	}

	ips := make([]string, 0)
	for i := range wg.IPAddress {
		ip, err := getIP(wg.IPAddress[i], a)
		if err != nil {
			return Peers{}, err
		}
		if IsIPv4(ip) {
			ip = fmt.Sprintf("%s/32", ip)
		}

		if IsIPv6(ip) {
			ip = fmt.Sprintf("%s/64", ip)
		}

		ips = append(ips, ip)
	}

	for i, v := range AllowedIPs {
		ipa := strings.Split(v, "/")
		IPNet, err := netip.ParsePrefix(v)
		if err != nil {
			return Peers{}, fmt.Errorf("not valid AllowedIP address")
		}
		IPNet = IPNet.Masked()
		ipa[0] = IPNet.Addr().String()

		if len(ipa) != 2 {
			ipa = append(ipa, fmt.Sprintf("%s/32", v))
		}
		AllowedIPs[i] = strings.Join(ipa, "/")

	}

	if err != nil {
		return Peers{}, err
	}

	priv, err := genPrivKey()
	if err != nil {
		return Peers{}, err
	}

	pub, err := genPubKey(priv)
	if err != nil {
		return Peers{}, err
	}

	Peer := Peers{Name: Name, IPAddress: ips, AllowedIPs: AllowedIPs, DNS: DNS, PresharedKey: psk, PrivateKey: priv, PublicKey: pub, PersistentKeepalive: pkl}
	return Peer, nil
}
