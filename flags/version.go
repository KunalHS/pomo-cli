package flags

import (
	"fmt"
	"os"
)

const versionString = "v0.1.2-beta"

// Implementation for the version flag.
func Version() {
	fmt.Println(versionString)
	os.Exit(0)
}
