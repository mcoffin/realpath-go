// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const (
	MajorVersion = 0
	MinorVersion = 0
)

func main() {
	printVersion := flag.Bool("v", false, "print version info then exit")
	stripOnly := flag.Bool("s", false, "only strip the path, don't resolve symlinks")

	flag.Parse()
	args := flag.Args()

	if *printVersion {
		fmt.Printf("realpath %d.%d\n", MajorVersion, MinorVersion)
		return
	}

	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, "Usage: realpath [-s] <path>")
		os.Exit(1)
	}

	newPath, err := filepath.Abs(args[0])
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}

	if !*stripOnly {
		newPath, err = filepath.EvalSymlinks(newPath)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println(newPath)
}
