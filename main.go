package main

import (
	"os"

	"github.com/k-kinzal/pr/cmd"
	//"github.com/spf13/cobra"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(255)
	}
}
