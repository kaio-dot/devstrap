package node

import (
	"fmt"
	"path/filepath"

	"github.com/kaio-dot/devstrap/internal/platform"
	paths "github.com/kaio-dot/devstrap/internal/utils"
)

type NodeProvider struct{}

func (n *NodeProvider) Name() string {
	return "node"
}

func (n *NodeProvider) GetLatestVersion() (string, error) {
	return getLatestLTS()
}

func (n *NodeProvider) getSpecificVersion(version string) (string, error) {
	return getSpecificVersion(version)
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

func (n *NodeProvider) ListAvailableVersions(limit int, onlyLTS bool) ([]NodeRelease, error) {
	return ListVersions(limit, onlyLTS)
}

func (n *NodeProvider) ExecutableDir(version string, p platform.Platform) (string, error) {
	return filepath.Join(paths.ToolsDir("node", version), "bin"), nil
}

func (n *NodeProvider) Validate() error {
	if version, err := getLatestLTS(); err != nil {
		return fmt.Errorf("não consegui validar Node.js: %w", err)
	} else {
		fmt.Printf("Node.js validado, versão LTS disponível: %s\n", version)
		return nil
	}
}

var Provider = &NodeProvider{}
