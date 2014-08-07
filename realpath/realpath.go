package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	stripOnly := flag.Bool("s", false, "only strip the path, don't resolve symlinks")

	flag.Parse()
	args := flag.Args()

	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, "Usage: realpath <path>")
		os.Exit(1)
	}

	newPath, err := filepath.Abs(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Bad path: %s\n", newPath)
		os.Exit(2)
	}

	fmt.Println(*stripOnly)

	if !*stripOnly {
		fmt.Println("Resolving symlinks")
		newPath, err = filepath.EvalSymlinks(newPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Bad path: %s\n", newPath)
			os.Exit(2)
		}
	}

	fmt.Println(newPath)
}
