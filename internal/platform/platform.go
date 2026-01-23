package platform

import "runtime"

type OS string
type Arch string

const (
	Windows OS = "Windows"
	Linux   OS = "Linux"
)

const (
	AMD64 Arch = "amd64"
	AMD32 Arch = "amd32"
	ARM64 Arch = "arm64"
	ARM32 Arch = "arm32"
)

type Platform struct {
	OS   OS
	Arch Arch
}

func DetectPlatform() Platform {
	return Platform{
		OS:   OS(runtime.GOOS),
		Arch: Arch(runtime.GOARCH),
	}
}

func (p Platform) IsSupported() bool {
	switch p.OS {
	case Windows, Linux:
		return true
	default:
		return false
	}
}
