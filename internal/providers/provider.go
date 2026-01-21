package providers

type Provider interface {
	Name() string
	Install(version string) error
	Update(version string) error
	Version(version string) (string, error)
	Uninstall(version string) error
}