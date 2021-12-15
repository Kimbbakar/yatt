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

func (l *localStorageRepo) setConfig(key string, value interface{}) {
	if err := l.client.SetCellValue(configSheet, "A"+configDetails[key]["row"], key); err != nil {
		log.Fatal(err)
	}

	if err := l.client.SetCellValue(configSheet, "B"+configDetails[key]["row"], value); err != nil {
		log.Fatal(err)
	}

	if err := l.client.Save(); err != nil {
		log.Fatal(err)
	}
}

func (l *localStorageRepo) getNewRow() string {
	curRow, err := strconv.Atoi(l.getConfig("currentRow"))
	if err != nil {
		log.Fatal(err)
	}

	l.setConfig("currentRow", curRow+1)
	return "A" + strconv.Itoa(curRow+1)
}

func (l *localStorageRepo) getNoteSheet() string {
	curRow, err := strconv.Atoi(l.getConfig("currentRow"))
	if err != nil {
		log.Fatal(err)
	}

	curSheet, err := strconv.Atoi(l.getConfig("currentNoteSheet"))
	if err != nil {
		log.Fatal(err)
	}

	if curRow >= rowLimit {
		curSheet++
		l.setConfig("currentRow", 0)
	}

	sheet := noteSheet + "-" + strconv.Itoa(curSheet)
	l.client.NewSheet(sheet)
	l.setConfig("currentNoteSheet", curSheet)

	return sheet
}

func (l *localStorageRepo) AddNote(note string) {
	sheet := l.getNoteSheet()
	row := l.getNewRow()
	key := appName + "-" + strings.Split(sheet, "-")[1] + "-" + row
	id := getUniqueID()
	date := time.Now().Format(time.RFC1123)

	if err := l.client.SetSheetRow(sheet, row, &[]interface{}{key, id, date, note}); err != nil {
		log.Fatal(err)
	}

	if err := l.client.Save(); err != nil {
		log.Fatal(err)
	}
}

func (l *localStorageRepo) ListNotes(sheetName string) [][]string {
	notes, err := l.client.GetRows(sheetName)
	if err != nil {
		log.Fatal(err)
	}

	return notes
}

func (l *localStorageRepo) FlashStorage() {
	if err := os.RemoveAll(filePath); err != nil {
		log.Fatal(err)
	}
}

func (l *localStorageRepo) NextSheet(sheetName string) string {
	if sheetName == "" {
		return l.getNoteSheet()
	}

	data := strings.Split(sheetName, "-")
	if data[1] == "0" {
		return ""
	}

	curSheet, err := strconv.Atoi(data[1])
	if err != nil {
		log.Fatal(err)
	}
	return data[0] + "-" + strconv.Itoa(curSheet-1)
}
