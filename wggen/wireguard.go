package wggen

import (
	"fmt"
	"io"
	"os/exec"
	"strings"
)

func isWGInstalled() bool {
	cmd := exec.Command("wg", "--version")
	err := cmd.Run()
	return err == nil
}

func genPrivKey() (string, error) {
	if !isWGInstalled() {
		return "", fmt.Errorf("WireGuard may not be installed or in your PATH variable")
	}
	cmd := exec.Command("wg", "genkey")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(output)), nil
}

func genPubKey(privatekey string) (string, error) {
	if !isWGInstalled() {
		return "", fmt.Errorf("WireGuard may not be installed or in your PATH variable")
	}
	cmd := exec.Command("wg", "pubkey")
	stdin, err := cmd.StdinPipe()

	if err != nil {
		return "", err
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, privatekey)
	}()

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

func genPSK() (string, error) {
	if !isWGInstalled() {
		return "", fmt.Errorf("WireGuard may not be installed or in your PATH variable")
	}
	cmd := exec.Command("wg", "genpsk")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil

}
