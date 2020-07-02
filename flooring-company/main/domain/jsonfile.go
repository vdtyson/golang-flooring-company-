package domain

import (
	"os"
)

type JsonContainer interface {
	WriteJson() error
	WasFileCreated() bool
	FilePath() string
	createFile() (*os.File, error)
	RemoveFile() error
	openAndGetFile() (*os.File, error)
	GetOrCreateFile() (*os.File, error)
}
