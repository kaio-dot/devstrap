package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

// ZipDownload builds a download path for a tool version under the temp directory.
func ZipDownload(tool, version string) string {
	if version == "" {
		version = "latest"
	}

	fileName := fmt.Sprintf("%s-%s.zip", tool, version)
	return filepath.Join(os.TempDir(), "devstrap", "downloads", fileName)
}
