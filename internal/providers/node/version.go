package node

import (
	"encoding/json"
	"errors"
	"net/http"
)

const nodeReleasesURL = "https://nodejs.org/dist/index.json"

type NodeRelease struct {
	Version string      `json:"version"`
	Date    string      `json:"date"`
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

func getSpecificVersion(version string) (string, error) {
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
		if r.Version == version {
			return r.Version, nil
		}
	}

	return "", errors.New("Versão não foi encontrada, tem certeza dessa versão aí?")
}

// ListVersions retorna uma lista de versões disponíveis do Node.js
// limit define quantas versões retornar (0 = todas)
// onlyLTS filtra apenas versões LTS
func ListVersions(limit int, onlyLTS bool) ([]NodeRelease, error) {
	resp, err := http.Get(nodeReleasesURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var releases []NodeRelease
	if err := json.NewDecoder(resp.Body).Decode(&releases); err != nil {
		return nil, err
	}

	var filtered []NodeRelease
	for _, r := range releases {
		if onlyLTS && r.LTS == nil {
			continue
		}
		filtered = append(filtered, r)

		if limit > 0 && len(filtered) >= limit {
			break
		}
	}

	return filtered, nil
}
