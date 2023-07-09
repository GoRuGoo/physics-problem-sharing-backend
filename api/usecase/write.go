package usecase

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

type WriteInteractor struct {
	repository WriteRepository
}

func (w WriteInteractor) UniqueFileNameAssignment(problenNumber string, fileData os.File) error {
	uuid, err := generateUUID()
	if err != nil {
		return errors.New(err.Error())
	}
	objectData := filepath.Join(problenNumber, uuid)
	err = w.repository.WriteHandler(objectData, fileData)
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
