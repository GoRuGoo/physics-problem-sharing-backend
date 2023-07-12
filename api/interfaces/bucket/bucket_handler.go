package bucket

import (
	"mime/multipart"
)

type BucketHandler interface {
	WriteExecute(string, multipart.File) error
}
