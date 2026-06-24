package utils

import "fmt"

func Display(args ...any) {
	fmt.Print(args...)
}

func DisplayLine(args ...any) {
	fmt.Println(args...)
}

func DisplayFormatter(pattern string, args ...any) {
	fmt.Printf(pattern, args...)
}
