package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for _, parm := range os.Args[:] {
		s += sep + parm
		sep = " "
	}
	fmt.Println(s)
}
