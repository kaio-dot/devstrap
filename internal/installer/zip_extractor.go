package installer

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func ExtractZip(zipPath, destDir string) error {
	if err := os.MkdirAll(destDir, 0755); err != nil {
		return err
	}

	fmt.Println("Extraindo pacote:", zipPath)
	fmt.Println("Destino da extração:", destDir)

	r, err := zip.OpenReader(zipPath)
	if err != nil {
		return err
	}

	defer r.Close()

	for _, f := range r.File {
		fpath := filepath.Join(destDir, f.Name)

		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, 0755)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(fpath), 0755); err != nil {
			return err
		}

		in, err := f.Open()
		if err != nil {
			return err
		}
		defer in.Close()

		out, err := os.Create(fpath)
		if err != nil {
			return err
		}

		defer out.Close()

		if _, err := io.Copy(out, in); err != nil {
			return err
		}
	}

	fmt.Println("Extração concluída para:", destDir)
	return nil
}
