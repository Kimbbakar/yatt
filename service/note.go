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
	note, _ := cmd.Flags().GetString("note")
	note = strings.Trim(note, " ")
	if note == "" {
		response("empty note not allowed", true, false, true)
		return nil
	}

	repo := repository.GetNewLocalStorage()
	repo.AddNote(note)

	return nil
}

func (n *NoteService) ListCommand(cmd *cobra.Command, args []string) error {
	tail, _ := cmd.Flags().GetInt("tail")
	if tail == 0 {
		tail = 20
	}

	fmt.Println("ID | Date | Note")
	repo := repository.GetNewLocalStorage()
	curSheet := repo.NextSheet("")
	for {
		if tail <= 0 || curSheet == "" {
			break
		}

		notes := repo.ListNotes(curSheet)
		for i := len(notes) - 1; i >= 0 && tail > 0; i-- {
			fmt.Printf("%s | %s | %s\n", notes[i][0], notes[i][1], notes[i][2])
			tail--
		}

		curSheet = repo.NextSheet(curSheet)
	}

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
