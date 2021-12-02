package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ServeVersionCmd = &cobra.Command{
	Use:   "version",
	Short: "version number",
	RunE:  serveVersion,
}

func serveVersion(cmd *cobra.Command, args []string) error {
	fmt.Print("version 0.1")
	return nil
}
