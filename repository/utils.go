package repository

import "sync"

var (
	once     sync.Once
	lStorage *localStorageRepo
	filePath = "./static/"
	fileName = "yatt.xlsx"
)

const (
	noteSheet   = "note"
	configSheet = "config"
)

var configDetails = map[string]map[string]string{
	"currentCell": {
		"default": "0",
		"row":     "2",
	},
}
