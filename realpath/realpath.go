package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, "Usage: realpath <path>")
		os.Exit(1)
	}
	newPath, err := filepath.EvalSymlinks(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Bad path: %s\n", newPath)
		os.Exit(2)
	} else {
		fmt.Println(newPath)
	}
}
