package service

import (
	"fmt"
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
		response("empty note not allowed", true, false, true)
		return nil
	}

	repo := repository.GetNewLocalStorage()
	repo.AddNote(title)

	return nil
}

func (n *NoteService) FlashStorageCommand(cmd *cobra.Command, args []string) error {
	var confirm string
	repo := repository.GetNewLocalStorage()

	response("This will remove all preset note/config", false, true, true)
	response("Are you sure you want to continue? [Y/n] ", false, true, false)
	fmt.Scanf("%s", &confirm)
	if !(confirm == "Y" || confirm == "n") {
		response("Wrong input", true, false, true)
	} else if confirm == "Y" {
		repo.FlashStorage()
		response("Storage flashed", false, false, true)
	}

	return nil
}
