package main

import (
	"fmt"
	"os"
)

func main() {
	os.Exit(run())
}

func run() int {
	for i, v := range os.Args[1:] {
		t := " "
		if i == 0 {
			t = ""
		}
		fmt.Printf("%s%s", t, v)
	}
	print("\n")

	return 0
}
