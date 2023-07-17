package bucket

import (
	"mime/multipart"
)

type BucketHandler interface {
	WriteExecute(string, multipart.File) error
	DeleteExecute(string) error
	SelectSpecificObjectsExecute(string) ([]string, error)
	SelectAllObjectsExecute() ([]string, error)
}
