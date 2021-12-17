package repository

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/xuri/excelize/v2"
)

type localStorageRepo struct {
	client *excelize.File
}

func getStorage() *excelize.File {
	f, err := excelize.OpenFile(filePath + fileName)
	if err != nil {
		if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
			log.Fatal(err)
		}

		f = excelize.NewFile()
		if err := f.SaveAs(filePath + fileName); err != nil {
			log.Fatal(err)
		}
	}

	f.NewSheet(configSheet)
	return f
}

func GetNewLocalStorage() *localStorageRepo {
	once.Do(func() {
		lStorage = &localStorageRepo{getStorage()}
	})

	return lStorage
}

func (l *localStorageRepo) getConfig(key string) string {
	v, err := l.client.GetCellValue(configSheet, "B"+configDetails[key]["row"])
	if err != nil || v == "" {
		v = configDetails[key]["default"]
	}

	return v
}

func (l *localStorageRepo) setConfig(key string, value interface{}) error {
	if err := l.client.SetCellValue(configSheet, "A"+configDetails[key]["row"], key); err != nil {
		return err
	}

	if err := l.client.SetCellValue(configSheet, "B"+configDetails[key]["row"], value); err != nil {
		return err
	}

	return l.client.Save()
}

func (l *localStorageRepo) getNewRow() (string, error) {
	curRow, err := strconv.Atoi(l.getConfig("currentRow"))
	if err != nil {
		return "", err
	}

	l.setConfig("currentRow", curRow+1)
	return "A" + strconv.Itoa(curRow+1), nil
}

func (l *localStorageRepo) getNoteSheet() (string, error) {
	curRow, err := strconv.Atoi(l.getConfig("currentRow"))
	if err != nil {
		return "", err
	}

	curSheet, err := strconv.Atoi(l.getConfig("currentNoteSheet"))
	if err != nil {
		return "", err
	}

	if curRow >= rowLimit {
		curSheet++
		l.setConfig("currentRow", 0)
	}

	sheet := noteSheet + "-" + strconv.Itoa(curSheet)
	l.client.NewSheet(sheet)
	l.setConfig("currentNoteSheet", curSheet)

	return sheet, nil
}

func (l *localStorageRepo) AddNote(note string) error {
	sheet, err := l.getNoteSheet()
	if err != nil {
		return err
	}
	row, err := l.getNewRow()
	if err != nil {
		return err
	}
	key := appName + "-" + strings.Split(sheet, "-")[1] + "-" + row
	id := getUniqueID()
	date := time.Now().Format(time.RFC1123)

	// key - id - date - note - deleted
	if err := l.client.SetSheetRow(sheet, row, &[]interface{}{key, id, date, note, false}); err != nil {
		return err
	}

	return l.client.Save()
}

func (l *localStorageRepo) UpdateNote(sheet, row string, value []interface{}) error {
	if err := l.client.SetSheetRow(sheet, row, &value); err != nil {
		return err
	}

	return l.client.Save()
}

func (l *localStorageRepo) ListNotes(sheetName string) ([][]string, error) {
	return l.client.GetRows(sheetName)
}

func (l *localStorageRepo) FlashStorage() error {
	return os.RemoveAll(filePath)
}

func (l *localStorageRepo) NextSheet(sheetName string) (string, error) {
	if sheetName == "" {
		return l.getNoteSheet()
	}

	data := strings.Split(sheetName, "-")
	if data[1] == "0" {
		return "", nil
	}

	curSheet, err := strconv.Atoi(data[1])
	if err != nil {
		return "", err
	}
	return data[0] + "-" + strconv.Itoa(curSheet-1), nil
}
