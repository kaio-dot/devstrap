package cli

import "strings"

// ParseToolVersion separa tool@version em (tool, version)
// Exemplos:
//   - "node" -> ("node", "latest")
//   - "node@20.11.0" -> ("node", "v20.11.0")
//   - "node@lts" -> ("node", "lts")
func ParseToolVersion(input string) (tool string, version string) {
	parts := strings.Split(input, "@")

	if len(parts) == 1 {
		return parts[0], "latest"
	}

	tool = parts[0]
	version = parts[1]

	// Adiciona o prefixo 'v' se não tiver e não for 'latest' ou 'lts'
	if version != "latest" && version != "lts" && !strings.HasPrefix(version, "v") {
		version = "v" + version
	}

	return tool, version
}
