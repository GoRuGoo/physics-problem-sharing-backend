package usecase

import (
	"errors"
	"mime/multipart"
	"path/filepath"

	"github.com/google/uuid"
)

type WriteInteractor struct {
	Repository WriteRepository
}

func (w WriteInteractor) UniqueFileNameAssignment(problenNumber string, fileData multipart.File) error {
	uuid, err := generateUUID()
	if err != nil {
		return errors.New(err.Error())
	}
	objectData := filepath.Join(problenNumber, uuid)
	err = w.Repository.WriteHandler(objectData, fileData)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func generateUUID() (string, error) {
	uuid, err := uuid.NewUUID()
	if err != nil {
		return uuid.String(), errors.New(err.Error())
	}
	return uuid.String(), nil
}
