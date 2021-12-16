package service

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/Kimbbakar/yatt/repository"
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

	repo := repository.GetNewLocalStorage()
	curSheet := repo.NextSheet("")
	for {
		if tail <= 0 || curSheet == "" {
			break
		}

		notes := repo.ListNotes(curSheet)
		for i := len(notes) - 1; i >= 0 && tail > 0; i-- {
			if deleted, err := strconv.Atoi(notes[i][4]); err != nil {
				log.Fatal(err)
			} else if deleted == 1 {
				continue
			}

			fmt.Printf("ID: %s\n", notes[i][1])
			fmt.Printf("Date: %s\n\n", notes[i][2])
			fmt.Print("    ")
			fmt.Printf("Note: %s\n", notes[i][3])
			tail--

			fmt.Println()
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

func (n *NoteService) DeleteCommand(cmd *cobra.Command, args []string) error {
	id := args[0]

	repo := repository.GetNewLocalStorage()
	curSheet := repo.NextSheet("")
	for {
		if curSheet == "" {
			break
		}

		notes := repo.ListNotes(curSheet)
		for i := len(notes) - 1; i >= 0; i-- {
			if strings.HasPrefix(notes[i][1], id) {
				row := strings.Split(notes[i][0], "-")[2]

				updateValue := make([]interface{}, len(notes[i]))
				for idx, v := range notes[i] {
					updateValue[idx] = v
				}

				updateValue[4] = true
				repo.UpdateNote(curSheet, row, updateValue)
				response("Note has been deleted successfully", false, false, true)
				return nil
			}
		}

		curSheet = repo.NextSheet(curSheet)
	}

	response("No note found with given ID", false, false, true)
	return nil
}
