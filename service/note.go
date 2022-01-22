package service

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/Kimbbakar/yatt/repository"
)

var (
	// key - id - date - note - description - deleted
	prefixIndent2 = "  "
	prefixIndent4 = "    "
	lineDevider   = "|yatt@yatt|"
)

const (
	KEY     = 0
	ID      = 1
	DATE    = 2
	NOTE    = 3
	DESC    = 4
	DELETED = 5
)

type NoteService struct {
}

func (n *NoteService) CreateCommand(cmd *cobra.Command, args []string) error {
	repo := repository.GetNewLocalStorage()

	if note, err := cmd.Flags().GetString("note"); err != nil {
		response(err.Error(), true, false, true)
	} else if note = strings.TrimSpace(note); note != "" {
		return repo.AddNote(note, "")
	}

	if note, err := cmd.Flags().GetString("note-with-description"); err != nil {
		response(err.Error(), true, false, true)
	} else if note = strings.TrimSpace(note); note != "" {
		desc, err := n.inputDescription()
		if err != nil {
			response(err.Error(), true, false, true)
		}

		desc = strings.TrimSpace(desc)

		return repo.AddNote(note, desc)
	}

	response("empty note not allowed", true, false, true)

	return nil
}

func (n *NoteService) ListCommand(cmd *cobra.Command, args []string) error {
	tail, _ := cmd.Flags().GetInt("tail")
	if tail == 0 {
		tail = 20
	}

	repo := repository.GetNewLocalStorage()
	curSheet, err := repo.NextSheet("")
	if err != nil {
		response(err.Error(), true, false, true)
	}
	for {
		if tail <= 0 || curSheet == "" {
			break
		}

		notes, err := repo.ListNotes(curSheet)
		if err != nil {
			response(err.Error(), true, false, true)
		}

		for i := len(notes) - 1; i >= 0 && tail > 0; i-- {
			if deleted, err := strconv.Atoi(notes[i][4]); err != nil {
				response(err.Error(), true, false, true)
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

		curSheet, err = repo.NextSheet(curSheet)
		if err != nil {
			response(err.Error(), true, false, true)
		}
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
	curSheet, err := repo.NextSheet("")
	if err != nil {
		response(err.Error(), true, false, true)
	}

	for {
		if curSheet == "" {
			break
		}

		notes, err := repo.ListNotes(curSheet)
		if err != nil {
			response(err.Error(), true, false, true)
		}

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

		curSheet, err = repo.NextSheet(curSheet)
		if err != nil {
			response(err.Error(), true, false, true)
		}
	}

	response("No note found with given ID", false, false, true)
	return nil
}

func (n *NoteService) inputDescription() (string, error) {
	fmt.Println("\nAdd the description[entry empty line to terminate]")
	in := bufio.NewReader(os.Stdin)
	details := ""
	for {
		fmt.Print(prefixIndent4)
		str, err := in.ReadString('\n')
		str = strings.Trim(str, " ")

		if err != nil {
			return "", err
		} else if str == "\n" {
			break
		}

		if len(details) > 0 {
			details += lineDevider
		}
		details += str
	}

	return details, nil
}
