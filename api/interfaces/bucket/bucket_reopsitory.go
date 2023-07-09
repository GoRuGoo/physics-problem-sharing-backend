package bucket

import (
	"errors"
	"os"
)

type BucketRepository struct {
	BucketHandler
}

func (b BucketRepository) WriteHandler(objectData string, fileData os.File) error {
	err := b.WriteExecute(objectData, fileData)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}
