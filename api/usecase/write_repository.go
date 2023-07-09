package usecase

import "os"

type WriteRepository interface {
	WriteHandler(string, os.File) error
}
