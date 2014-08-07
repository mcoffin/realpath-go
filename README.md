# realpath-go

A minimal implementation of the `realpath(1)` command. Originally created because OSX lacks the command, and many installation scripts depend upon it.

## Installation

To install, simply run `go get github.com/mcoffin/realpath-go/realpath`

## Usage

`realpath [-s] filename`
* `-s`: Only strip the path, don't evaluate symlinks
