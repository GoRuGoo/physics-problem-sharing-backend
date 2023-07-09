package controller

import (
	"net/http"
	"os"
	"physics/interfaces/bucket"
	"physics/usecase"

	"github.com/gin-gonic/gin"
)

type WriteController struct {
	Interactor usecase.WriteInteractor
}

func NewWriteController(bucketHandler bucket.BucketHandler) *WriteController {
	return &WriteController{
		Interactor: usecase.WriteInteractor{
			Repository: bucket.BucketRepository{
				BucketHandler: bucketHandler,
			},
		},
	}
}

func (w WriteController) WriteWithProblemNumberController(c Context, query string) {
	problem_num := c.Param(query)
	var gazou os.File
	err := w.Interactor.UniqueFileNameAssignment(problem_num, gazou)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"ok": "ok"})
}
