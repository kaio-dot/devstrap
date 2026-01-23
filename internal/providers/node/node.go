package node

import (
	"fmt"

	"github.com/kaio-dot/devstrap/internal/platform"
)

type NodeProvider struct{}

func (n *NodeProvider) Name() string {
	return "node"
}

func (n *NodeProvider) GetLatestVersion() (string, error) {
	return getLatestLTS()
}

func (n *NodeProvider) Update(version string, p platform.Platform) error {
	fmt.Println("Atualizando Node.js para a versão", version)
	return nil
}

func (n *NodeProvider) Version(version string) (string, error) {
	return version, nil
}

func (n *NodeProvider) Uninstall(version string, p platform.Platform) error {
	fmt.Println("Desisntalando Node.js", version)
	return nil
}

var Provider = &NodeProvider{}
