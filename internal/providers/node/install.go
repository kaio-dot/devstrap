package node

import (
	"fmt"

	"github.com/kaio-dot/devstrap/internal/platform"
)

func (n *NodeProvider) Install(version string, p platform.Platform) error {
	var url string

	switch p.OS {
	case platform.Linux:
		url = fmt.Sprintf(
			"https://nodejs.org/dist/v%s/node-v%s-linux-x64.tar.xz",
			version, version,
		)

	case platform.Windows:
		url = fmt.Sprintf(
			"https://nodejs.org/dist/v%s/node-v%s-win-x64.zip",
			version, version,
		)

	default:
		return fmt.Errorf("Ainda não temos suporte pra esse teu OS, te acalma aí: %s",
			p.OS)
	}

	fmt.Println("Instalando versão do Node:", url)

	return nil
}
