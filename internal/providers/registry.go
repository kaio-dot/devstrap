package providers

import "github.com/kaio-dot/devstrap/internal/providers/node"

var registry = map[string]Provider{
	"node": node.Provider,
}

func RegisterProvider(p Provider) {
	registry[p.Name()] = p
}

func GetProvider(name string) (Provider, bool) {
	p, ok := registry[name]
	return p, ok
}

func DebugRegistry() map[string]Provider {
	return registry
}
