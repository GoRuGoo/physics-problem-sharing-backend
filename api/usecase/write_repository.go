package usecase

import (
	"mime/multipart"
)

type ManipulateRepository interface {
	WriteHandler(string, multipart.File) error
}
