package repository

import "sync"

var (
	once     sync.Once
	lStorage *localStorageRepo
	appName  = "YATT"
	filePath = "/.yatt/"
	fileName = "storage.xlsx"
	rowLimit = 20
)

const (
	noteSheet   = "note"
	configSheet = "config"
)

var configDetails = map[string]map[string]string{
	"currentRow": {
		"default": "0",
		"row":     "2",
	},
	"currentNoteSheet": {
		"default": "0",
		"row":     "3",
	},
}
