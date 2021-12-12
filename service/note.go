package service

import (
	"github.com/spf13/cobra"

	"github.com/kimbbakar/yatt/repository"
)

type NoteService struct {
}

func (n *NoteService) CreateCommand(cmd *cobra.Command, args []string) error {
	title, _ := cmd.Flags().GetString("title")
	repo := repository.GetNewLocalStorage()
	repo.AddNote(title)

	return nil
}
