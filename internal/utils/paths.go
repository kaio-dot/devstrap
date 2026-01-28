package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/sys/windows/registry"
)

// ZipDownload constrói ao caminho para o arquivo zip baixado
func ZipDownload(tool, version string) string {
	if version == "" {
		version = "latest"
	}

	fileName := fmt.Sprintf("%s-%s.zip", tool, version)
	return filepath.Join(os.TempDir(), "devstrap", "downloads", fileName)
}

func ToolsDir(tool, version string) string {
	if version == "" {
		version = "latest"
	}
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return filepath.Join(home, ".devstrap", "tools", tool, version)
}

func AddToPathWindows(dir string) error {
	fmt.Printf("Tentando adicionar ao PATH: %s\n", dir)

	key, _, err := registry.CreateKey(
		registry.CURRENT_USER,
		"Environment",
		registry.SET_VALUE,
	)

	if err != nil {
		return fmt.Errorf("erro ao abrir chave do registro: %w", err)
	}

	defer key.Close()

	path, _, err := key.GetStringValue("Path")
	if err != nil {
		fmt.Printf("Aviso: erro ao ler PATH atual: %v\n", err)
		path = "" // Começar vazio se não conseguir ler
	}

	if !containsPath(path, dir) {
		newPath := fmt.Sprintf("%s;%s", path, dir)
		fmt.Printf("Adicionando ao PATH do registro...\n")
		if err := key.SetStringValue("Path", newPath); err != nil {
			return fmt.Errorf("erro ao escrever PATH no registro: %w", err)
		}
		fmt.Printf("✓ PATH atualizado no registro!\n")
	} else {
		fmt.Printf("✓ Diretório já está no PATH\n")
	}

	return nil
}

func containsPath(path, dir string) bool {
	return strings.Contains(
		strings.ToLower(path),
		strings.ToLower(dir),
	)
}
