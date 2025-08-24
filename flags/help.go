package flags

import (
	"flag"
	"fmt"
)

func Help() {
	fmt.Println("\nWelcome to Pomo — a CLI-based Pomodoro timer!\n")
	fmt.Println("Usage:")
	fmt.Println("  pomo [flags]\n")
	fmt.Println("Options:")
	flag.PrintDefaults()
	fmt.Println("")
}
