package service

import (
	"strings"

	"github.com/spf13/cobra"

	"github.com/kimbbakar/yatt/repository"
)

type NoteService struct {
}

func (n *NoteService) CreateCommand(cmd *cobra.Command, args []string) error {
	title, _ := cmd.Flags().GetString("title")
	title = strings.Trim(title, " ")
	if title == "" {
		response("empty note not allowed", true)
		return nil
	}

	repo := repository.GetNewLocalStorage()
	repo.AddNote(title)

	return nil
}
