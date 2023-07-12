package bucket

import (
	"errors"
	"mime/multipart"
)

type BucketRepository struct {
	BucketHandler
}

func (b BucketRepository) WriteHandler(objectData string, fileData multipart.File) error {
	err := b.WriteExecute(objectData, fileData)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}
