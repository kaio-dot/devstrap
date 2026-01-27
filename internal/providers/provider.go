package providers

import (
	"github.com/kaio-dot/devstrap/internal/platform"
	"github.com/kaio-dot/devstrap/internal/providers/node"
)

type Provider interface {
	Name() string
	GetLatestVersion() (string, error)
	Install(version string, p platform.Platform) error
	Update(version string, p platform.Platform) error
	Version(version string) (string, error)
	Uninstall(version string, p platform.Platform) error
	ListAvailableVersions(limit int, onlyLTS bool) ([]node.NodeRelease, error)
}
