package main

import (
	"flag"
	"pomo-cli/flags"
)

func main() {
	versionFlag := flag.Bool("version", false, "Print version")
	shorterVersionFlag := flag.Bool("v", false, "Print version")
	flag.Parse()

	switch {
	case *versionFlag || *shorterVersionFlag:
		flags.Version()
	}
	flags.Help()
}
