package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"flag"
)

var (
	varName  = flag.String("var", "x", "variable name")
	packName = flag.String("pack", "main", "package name")
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Fprintln(os.Stderr, "please specify a file")
		os.Exit(1)
	}

	bytes, e := ioutil.ReadFile(os.Args[1])
	if e != nil {
		fmt.Fprintln(os.Stderr, e)
		os.Exit(1)
	}

	fmt.Printf("package %s\n", *packName)
	fmt.Println()
	fmt.Printf("var %s = []byte{", *varName)

	for i, b := range bytes {
		if i%8 == 0 {
			fmt.Println()
			fmt.Print("\t")
		} else {
			fmt.Print(" ")
		}

		fmt.Printf("0x%02x,", b)
	}
	fmt.Println()
	fmt.Println("}")
}
