package main

import (
	"fmt"
	"os"
)

func main() {
	for index, parm := range os.Args[1:] {
		fmt.Printf("index: %d, parm: %s\n", index, parm)
	}
}
