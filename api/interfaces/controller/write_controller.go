package controller

import (
	"fmt"
	"net/http"
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

func (w WriteController) WriteWithProblemNumberController(c *gin.Context, query string) {
	problem_num := c.Param(query)
	file, err := c.FormFile("problem_file")
	if err != nil {
		fmt.Println("file erro")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	problem_src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	defer problem_src.Close()

	err = w.Interactor.UniqueFileNameAssignment(problem_num, problem_src)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"ok": "ok"})
}
