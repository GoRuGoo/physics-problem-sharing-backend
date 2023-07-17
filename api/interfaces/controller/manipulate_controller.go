package controller

import (
	"net/http"
	"path/filepath"
	"physics/interfaces/bucket"
	"physics/usecase"

	"github.com/gin-gonic/gin"
)

type ManipulateController struct {
	Interactor    usecase.ManipulateInteractor
	BucketHandler bucket.BucketHandler
}

func NewManipulateController(bucketHandler bucket.BucketHandler) *ManipulateController {
	return &ManipulateController{
		Interactor: usecase.ManipulateInteractor{
			Repository: bucket.BucketRepository{
				BucketHandler: bucketHandler,
			},
		},
		BucketHandler: bucketHandler,
	}
}

func (w ManipulateController) WriteWithProblemNumberController(c *gin.Context, problem_num_query string) {
	problem_num := c.Param(problem_num_query)
	file, err := c.FormFile("problem_file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	problem_src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer problem_src.Close()

	err = w.Interactor.UniqueFileNameAssignment(problem_num, problem_src)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ok": "ok"})
	return
}

func (w ManipulateController) DeleteWithFileNumberController(c *gin.Context) {
	problem_dir := c.Param("problem_dir")
	problem_num := c.Param("problem_num")

	object_path := filepath.Join(problem_dir, problem_num)

	err := w.BucketHandler.DeleteExecute(object_path)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ok": "ok"})
	return
}
func (w ManipulateController) GetAllObjectsController(c *gin.Context) {
	problem_lists, err := w.BucketHandler.SelectAllObjectsExecute()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ok": problem_lists})
	return
}

func (w ManipulateController) GetSpecificObjectsController(c *gin.Context) {
	problem_num := c.Param("problem_num")
	problem_lists, err := w.BucketHandler.SelectSpecificObjectsExecute(problem_num)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ok": problem_lists})
	return
}
