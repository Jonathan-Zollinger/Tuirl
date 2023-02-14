package main

import (
	"os"

	"github.com/Jonathan-Zollinger/Tuirl/internal/cmd"
)

var version = "dev"

func main() {
	cmd.Execute(
		version,
		os.Exit,
		os.Args[1:],
	)
}
