package node

import (
	"fmt"

	"github.com/kaio-dot/devstrap/internal/installer"
	"github.com/kaio-dot/devstrap/internal/platform"
	paths "github.com/kaio-dot/devstrap/internal/utils"
)

func (n *NodeProvider) Install(version string, p platform.Platform) error {
	var url string

	switch p.OS {
	case platform.Linux:
		url = fmt.Sprintf(
			"https://nodejs.org/dist/%s/node-%s-linux-x64.tar.xz",
			version, version,
		)

	case platform.Windows:
		url = fmt.Sprintf(
			"https://nodejs.org/dist/%s/node-%s-win-x64.zip",
			version, version,
		)

	default:
		return fmt.Errorf("Ainda não temos suporte pra esse teu OS, te acalma aí: %s",
			p.OS)
	}

	zipPath := paths.ZipDownload("node", version)
	destDir := paths.ToolsDir("node", version)

	_, err := installer.DownloadTool(url, zipPath)
	if err != nil {
		return err
	}
	fmt.Println("Download concluído:", zipPath)

	if err := installer.ExtractZip(zipPath, destDir); err != nil {
		return err
	}

	fmt.Println("Instalando versão do Node:", url)

	return nil
}
