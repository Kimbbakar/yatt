package repository

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
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

	index := f.NewSheet(noteSheet)
	f.SetActiveSheet(index)
	f.NewSheet(configSheet)

	return f
}

func GetNewLocalStorage() *localStorageRepo {
	once.Do(func() {
		lStorage = &localStorageRepo{getStorage()}
	})

	return lStorage
}

func (l *localStorageRepo) getNewRow() string {
	v, err := l.client.GetCellValue(configSheet, "B"+configDetails["currentCell"]["row"])
	if err != nil || v == "" {
		v = configDetails["currentCell"]["default"]
	}

	curRow, err := strconv.Atoi(v)
	if err != nil {
		log.Fatal(err)
	}

	if err := l.client.SetCellValue(configSheet, "A"+configDetails["currentCell"]["row"], "currentCell"); err != nil {
		log.Fatal(err)
	}

	if err := l.client.SetCellValue(configSheet, "B"+configDetails["currentCell"]["row"], curRow+1); err != nil {
		log.Fatal(err)
	}

	if err := l.client.Save(); err != nil {
		log.Fatal(err)
	}

	return "A" + strconv.Itoa(curRow+1)
}

func (l *localStorageRepo) AddNote(note string) {
	id := uuid.New()

	row := l.getNewRow()
	if err := l.client.SetSheetRow(noteSheet, row, &[]interface{}{id, time.Now(), note}); err != nil {
		log.Fatal(err)
	}

	if err := l.client.Save(); err != nil {
		log.Fatal(err)
	}
}
