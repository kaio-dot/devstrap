package node

import (
	"encoding/json"
	"errors"
	"net/http"
)

const nodeReleasesURL = "https://nodejs.org/dist/index.json"

type NodeRelease struct {
	Version string      `json:"version"`
	LTS     interface{} `json:"lts"`
}

func getLatestLTS() (string, error) {
	resp, err := http.Get(nodeReleasesURL)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var releases []NodeRelease
	if err := json.NewDecoder(resp.Body).Decode(&releases); err != nil {
		return "", err
	}

	for _, r := range releases {
		if r.LTS != nil {
			return r.Version, nil
		}
	}

	return "", errors.New("Nenhuma versão encontrada, tem certeza dessa versão aí?")
}
