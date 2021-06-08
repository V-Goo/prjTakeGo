//usr/bin/env go run "$0" "$@"; exit

package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) < 1 {
		fmt.Println("usage: entr new dir name")
		os.Exit(1)
	}

}
