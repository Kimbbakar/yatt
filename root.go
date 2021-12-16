package main

import (
	"fmt"
	"os"

	"github.com/Kimbbakar/yatt/command"
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
	command.AddListNoteCommand(RootCmd)
	command.AddDeleteNoteCommand(RootCmd)
}

func main() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
