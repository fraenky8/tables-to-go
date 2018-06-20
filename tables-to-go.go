package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fraenky8/tables-to-go/internal/cli"
)

// main function to run the transformations
func main() {

	cmdArgs := cli.NewCmdArgs()

	if cmdArgs.Help {
		flag.Usage()
		os.Exit(0)
	}

	if err := cmdArgs.Verify(); err != nil {
		fmt.Printf("settings verification error: %v", err)
		os.Exit(1)
	}

	if err := cli.Run(cmdArgs.Settings); err != nil {
		fmt.Printf("run error: %v", err)
		os.Exit(1)
	}
}
