package command

import (
	"github.com/spf13/cobra"

	"github.com/kimbbakar/yatt/service"
)

func AddCreateNoteCommand(rootCommand *cobra.Command) {
	noteSrv := &service.NoteService{}
	createNoteCmd := &cobra.Command{
		Use:   "create",
		Short: "create note",
		RunE:  noteSrv.CreateCommand,
	}

	createNoteCmd.PersistentFlags().StringP("title", "t", "", "add title")
	rootCommand.AddCommand(createNoteCmd)
}

func AddFlashStorageCommand(rootCommand *cobra.Command) {
	noteSrv := &service.NoteService{}
	flashStorageCmd := &cobra.Command{
		Use:   "flash",
		Short: "flash all note/config",
		RunE:  noteSrv.FlashStorageCommand,
	}

	rootCommand.AddCommand(flashStorageCmd)
}
