package node

import (
	"fmt"

	"github.com/kaio-dot/devstrap/internal/installer"
	"github.com/kaio-dot/devstrap/internal/platform"
	paths "github.com/kaio-dot/devstrap/internal/utils"
)

func (n *NodeProvider) Install(version string, p platform.Platform) error {
	// Resolve a versão
	var resolvedVersion string
	var err error

	switch version {
	case "latest", "lts":
		resolvedVersion, err = getLatestLTS()
		if err != nil {
			return fmt.Errorf("erro ao buscar versão LTS: %w", err)
		}
	default:
		// Valida se a versão específica existe
		resolvedVersion, err = getSpecificVersion(version)
		if err != nil {
			return fmt.Errorf("versão não encontrada: %w", err)
		}
	}

	fmt.Printf("Versão resolvida: %s\n", resolvedVersion)

	var url string

	switch p.OS {
	case platform.Linux:
		url = fmt.Sprintf(
			"https://nodejs.org/dist/%s/node-%s-linux-x64.tar.xz",
			resolvedVersion, resolvedVersion,
		)

	case platform.Windows:
		url = fmt.Sprintf(
			"https://nodejs.org/dist/%s/node-%s-win-x64.zip",
			resolvedVersion, resolvedVersion,
		)

	default:
		return fmt.Errorf("Ainda não temos suporte pra esse teu OS, te acalma aí: %s",
			p.OS)
	}

	zipPath := paths.ZipDownload("node", resolvedVersion)
	destDir := paths.ToolsDir("node", resolvedVersion)

	_, err = installer.DownloadTool(url, zipPath)

	if err != nil {
		return err
	}
	fmt.Println("Download concluído:", zipPath)

	if err := installer.ExtractZip(zipPath, destDir); err != nil {
		return err
	}

	fmt.Println("Instalando versão do Node:", url)

	return installer.InstallTool(destDir, p)
}
