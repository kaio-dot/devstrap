package ui

import "fmt"

type ProgressBar struct {
	total int64
	width int
}

func NewProgressBar(total int64, width int) *ProgressBar {
	return &ProgressBar{
		total: total,
		width: 30,
	}
}

func (p *ProgressBar) Render(current int64) {
	percent := float64(current) / float64(p.total)
	filled := int(percent * float64(p.width))

	bar := ""

	for i := 0; i < p.width; i++ {
		if i < filled {
			bar += "█"
		} else {
			bar += "░"
		}
	}

	fmt.Printf("\r[%s] %3.0f%%", bar, percent*100)
}

func (p *ProgressBar) Complete() {
	fmt.Println()
}
