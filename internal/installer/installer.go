package installer

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/kaio-dot/devstrap/internal/platform"
	"github.com/kaio-dot/devstrap/internal/utils"
)

// InstallTool instala uma ferramenta após extração, adicionando ao PATH
func InstallTool(extractedDir string, p platform.Platform) error {
	// Verificar se o diretório existe
	if _, err := os.Stat(extractedDir); os.IsNotExist(err) {
		return fmt.Errorf("diretório não encontrado: %s", extractedDir)
	}

	// Encontrar o diretório bin dentro da pasta extraída
	binDir, err := findBinDirectory(extractedDir, p)
	if err != nil {
		return fmt.Errorf("não encontrei o diretório bin: %w", err)
	}

	fmt.Printf("Diretório bin encontrado: %s\n", binDir)

	// Adicionar ao PATH baseado no sistema operacional
	switch p.OS {
	case platform.Windows:
		if err := utils.AddToPathWindows(binDir); err != nil {
			return fmt.Errorf("erro ao adicionar ao PATH no Windows: %w", err)
		}
		fmt.Println("✓ Adicionado ao PATH do Windows")

	case platform.Linux:
		fmt.Printf("Para usar a ferramenta, adicione ao seu PATH:\n")
		fmt.Printf("export PATH=\"%s:$PATH\"\n", binDir)
		fmt.Println("✓ Ferramenta instalada (adicione ao PATH manualmente)")

	default:
		return fmt.Errorf("sistema operacional não suportado: %s", p.OS)
	}

	return nil
}

// findBinDirectory procura pelo diretório com executáveis dentro do diretório extraído
func findBinDirectory(extractedDir string, p platform.Platform) (string, error) {
	// No Windows, procurar por node.exe na raiz ou subdiretórios
	// No Linux/Mac, procurar pela pasta bin/

	execName := "node"
	if p.OS == platform.Windows {
		execName = "node.exe"
	}

	// Tentar na raiz do diretório extraído
	if fileExists(filepath.Join(extractedDir, execName)) {
		fmt.Printf("✓ Executável encontrado na raiz: %s\n", extractedDir)
		return extractedDir, nil
	}

	// Tentar diretório bin/ direto (Linux/Mac)
	binDir := filepath.Join(extractedDir, "bin")
	if fileExists(filepath.Join(binDir, execName)) {
		fmt.Printf("✓ Executável encontrado em bin: %s\n", binDir)
		return binDir, nil
	}

	// Procurar em subdiretórios (caso o zip tenha uma pasta pai)
	entries, err := os.ReadDir(extractedDir)
	if err != nil {
		return "", err
	}

	fmt.Printf("Procurando executável em subdiretórios de: %s\n", extractedDir)

	for _, entry := range entries {
		if entry.IsDir() {
			subDir := filepath.Join(extractedDir, entry.Name())
			fmt.Printf("  Checando subdiretório: %s\n", entry.Name())

			// Verificar se o executável está na raiz do subdiretório (Windows)
			if fileExists(filepath.Join(subDir, execName)) {
				fmt.Printf("✓ Executável encontrado em: %s\n", subDir)
				return subDir, nil
			}

			// Verificar se /subdir/bin existe (Linux/Mac)
			binInSubDir := filepath.Join(subDir, "bin")
			if fileExists(filepath.Join(binInSubDir, execName)) {
				fmt.Printf("✓ Executável encontrado em: %s\n", binInSubDir)
				return binInSubDir, nil
			}
		}
	}

	// Se não encontrou, retornar erro
	return "", fmt.Errorf("executável '%s' não encontrado em %s", execName, extractedDir)
}

// fileExists verifica se um arquivo existe
func fileExists(filePath string) bool {
	info, err := os.Stat(filePath)
	return err == nil && !info.IsDir()
}

// GetExecutablePath retorna o caminho completo do executável da ferramenta
func GetExecutablePath(binDir, toolName string, p platform.Platform) string {
	execName := toolName
	if p.OS == platform.Windows {
		execName += ".exe"
	}
	return filepath.Join(binDir, execName)
}
