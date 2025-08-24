package main

import (
	"flag"
	"fmt"
	"pomo-cli/flags"
)

func main() {
	versionFlag := flag.Bool("version", false, "Print version")
	flag.Parse()

	switch {
	case *versionFlag:
		flags.Version()
	}
	fmt.Println("This is the help screen, I will populate it later!")
}
