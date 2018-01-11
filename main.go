package main

import (
	"fmt"
	"os"

	"github.com/harikiriboy/nifcloud-ssh-config/commands"
)

func main() {
	if err := commands.RootCmd.Execute(); err != nil {
		fmt.Printf("nifcloud-ssh-config command is error: %s", err)
		os.Exit(1)
	}
}
