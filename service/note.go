package service

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/kimbbakar/yatt/repository"
)

type NoteService struct {
}

func (n *NoteService) RunCommand(cmd *cobra.Command, args []string) error {
	title, _ := cmd.Flags().GetString("title")
	n.run(title)
	return nil
}

func (n *NoteService) run(title string) error {
	repo := repository.GetNewLocalStorage()

	log.Println(repo.GetNewRow())
	return nil
}
