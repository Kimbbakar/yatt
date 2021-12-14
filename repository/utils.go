package repository

import "sync"

var (
	once     sync.Once
	lStorage *localStorageRepo
	filePath = "/.yatt/"
	fileName = "storage.xlsx"
)

const (
	noteSheet   = "note"
	configSheet = "config"
)

var configDetails = map[string]map[string]string{
	"currentRow": {
		"default": "0",
		"value":   "2",
	},
}
