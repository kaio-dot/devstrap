package node

import (
	"fmt"

	"github.com/kaio-dot/devstrap/internal/providers"
)

type NodeProvider struct{}

// Garantir conformidade com a interface Provider
var _ providers.Provider = (*NodeProvider)(nil)

func (n *NodeProvider) Name() string {
	return "node"
}

func (n *NodeProvider) Install(version string) error {
	if version == "latest" || version == "lts" || version == "" {
		v, err := getLatestLTS()
		if err != nil {
			return err
		}
		version = v
	}

	fmt.Println("Versão do Node encontrada: ", version)
	return nil
}

func (n *NodeProvider) Update(version string) error {
	fmt.Println("Atualizando Node.js para a versão", version)
	return nil
}

func (n *NodeProvider) Version(version string) (string, error) {
	return version, nil
}

func (n *NodeProvider) Uninstall(version string) error {
	fmt.Println("Desisntalando Node.js", version)
	return nil
}

func init() {
	providers.RegisterProvider(&NodeProvider{})
}
