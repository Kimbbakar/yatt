package main

import (
	"fmt"
	"os"

	"github.com/kimbbakar/yatt/command"
	"github.com/spf13/cobra"
)

// RootCmd of the binary
var (
	RootCmd = &cobra.Command{
		Short: "Tell me your thoughts",
	}
)

func init() {
	RootCmd.AddCommand(command.ServeVersionCmd)
	command.AddCreateNoteCommand(RootCmd)
	command.AddFlashStorageCommand(RootCmd)
}

// Execute executes the root command
func Execute() {

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	Execute()
}
