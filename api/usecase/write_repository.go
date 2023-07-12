package usecase

import (
	"mime/multipart"
)

type WriteRepository interface {
	WriteHandler(string, multipart.File) error
}
