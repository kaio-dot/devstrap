package providers

import _ "github.com/kaio-dot/devstrap/internal/providers/node"

var registry = make(map[string]Provider)

func RegisterProvider(p Provider) {
	registry[p.Name()] = p
}

func GetProvider(name string) (Provider, bool) {
	p, ok := registry[name]
	return p, ok
}
