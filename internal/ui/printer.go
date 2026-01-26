package ui

import "fmt"

func Step(msg string) {
	fmt.Printf("→ %s\n", msg)
}

func Success(msg string) {
	fmt.Printf("✔ %s\n", msg)
}

func Error(msg string) {
	fmt.Printf("✖ %s\n", msg)
}

func Info(msg string) {
	fmt.Printf("ℹ %s\n", msg)
}
