package main

import (
	og "Og/src"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	for _, arg := range args {
		fmt.Println(og.Compile(arg))
	}
}
