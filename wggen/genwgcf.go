package wggen

import (
	"embed"
	"fmt"
	"strings"
	"text/template"
)

//go:embed templates/*.tmpl
var templateFiles embed.FS

func renderTemplate(tmplStr string, data interface{}) (string, error) {
	tmpl, err := template.New("").Parse(tmplStr)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w", err)
	}

	var sb strings.Builder
	err = tmpl.Execute(&sb, data)
	if err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}

	return sb.String(), nil
}

func (wg *WGSrv) GeneratePeerConfig(name string) (wgcfg string, err error) {
	wgsrv := *wg
	wgsrv.Peers = nil
	// Read the template files
	peerTmpl, err := templateFiles.ReadFile("templates/peer.tmpl")
	if err != nil {
		return "", fmt.Errorf("failed to read peer.tmpl: %w", err)
	}

	for _, p := range wg.Peers {
		if name != p.Name {
			continue
		}
		wgsrv.Peers = append(wgsrv.Peers, p)
	}
	wgcfg, err = renderTemplate(string(peerTmpl), wgsrv)
	if err != nil {
		return "", fmt.Errorf("failed to render peer: %w", err)
	}

	return wgcfg, nil
}

func (wg *WGSrv) GenerateSrvConfig() (wgcfg string, err error) {
	// Read the template files
	peerTmpl, err := templateFiles.ReadFile("templates/server.tmpl")
	if err != nil {
		return "", fmt.Errorf("failed to read server.tmpl: %w", err)
	}

	wgcfg, err = renderTemplate(string(peerTmpl), wg)
	if err != nil {
		return "", fmt.Errorf("failed to render Server: %w", err)
	}

	return wgcfg, nil
}
