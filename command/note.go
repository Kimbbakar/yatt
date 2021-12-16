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

	createNoteCmd.PersistentFlags().StringP("note", "n", "", "add note")
	rootCommand.AddCommand(createNoteCmd)
}

func AddListNoteCommand(rootCommand *cobra.Command) {
	noteSrv := &service.NoteService{}
	listNoteCmd := &cobra.Command{
		Use:   "list",
		Short: "list note",
		RunE:  noteSrv.ListCommand,
	}

	listNoteCmd.PersistentFlags().IntP("tail", "t", 0, "add title")
	rootCommand.AddCommand(listNoteCmd)
}

func AddDeleteNoteCommand(rootCommand *cobra.Command) {
	noteSrv := &service.NoteService{}
	deleteNoteCmd := &cobra.Command{
		Use:   "delete",
		Short: "delete note",
		RunE:  noteSrv.DeleteCommand,
	}

	rootCommand.AddCommand(deleteNoteCmd)
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
