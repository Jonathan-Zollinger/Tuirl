package cmd

import (
	"fmt"
)

func Execute(version string, exit func(int), args []string) {
	fmt.Printf("log.Logger: %v\n", version)
}
