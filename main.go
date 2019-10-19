package main

import (
	//"github.com/k-kinzal/pr/logger"
	"fmt"
	"github.com/k-kinzal/pr/cmd"
	"os"
	//"github.com/spf13/cobra"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, fmt.Sprintf("pr: %s", err))
		os.Exit(1)
	}
}
