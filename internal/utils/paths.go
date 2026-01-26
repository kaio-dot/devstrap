package utils

import (
	"fmt"
	"os"
	"path/filepath"
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
