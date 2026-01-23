package installer

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func DownloadTool(url, destDir string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Tá difícil baixar essa porcaria, mas vai dar certo: %s", resp.Status)
	}

	if err := os.MkdirAll(filepath.Dir(destDir), 0755); err != nil {
		return "", err
	}

	out, err := os.Create(destDir)

	if err != nil {
		return "", err
	}

	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return destDir, err
}
