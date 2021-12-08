package command

import (
	"log"

	"github.com/spf13/cobra"
)

func AddCreateNoteCommand(rootCommand *cobra.Command) {
	createNoteCmd := &cobra.Command{
		Use:   "create",
		Short: "create note",
		RunE: func(cmd *cobra.Command, args []string) error {
			title, _ := cmd.Flags().GetString("title")
			log.Print(title)

			return nil
		},
	}

	createNoteCmd.PersistentFlags().StringP("title", "t", "", "add title")

	rootCommand.AddCommand(createNoteCmd)
}
