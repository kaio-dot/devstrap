package installer

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/kaio-dot/devstrap/internal/ui"
)

func DownloadTool(url, destPath string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Erro no download: %s", resp.Status)
	}

	if err := os.MkdirAll(filepath.Dir(destPath), 0755); err != nil {
		return "", err
	}

	out, err := os.Create(destPath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	progressBar := ui.NewProgressBar(resp.ContentLength, 30)
	counter := &WriteCounter{Bar: progressBar}

	_, err = io.Copy(out, io.TeeReader(resp.Body, counter))
	if err != nil {
		return "", err
	}

	progressBar.Complete()
	fmt.Println("Download concluído:", destPath)
	return destPath, nil
}

type WriteCounter struct {
	Bar        *ui.ProgressBar
	Downloaded int64
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Downloaded += int64(n)
	wc.Bar.Render(wc.Downloaded)
	return n, nil
}
