package infrastructure

import (
	"physics/interfaces/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func InitializeRouter() {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	router.Use(cors.New(config))

	controller := controller.NewManipulateController(NewBucketHandler())

	router.GET("/object_lists/all", func(c *gin.Context) { controller.GetAllObjectsController(c) })
	router.GET("/object_lists/:problem_num", func(c *gin.Context) { controller.GetSpecificObjectsController(c) })
	router.POST("/write/:problem_num", func(c *gin.Context) { controller.WriteWithProblemNumberController(c, "problem_num") })
	router.DELETE("/delete/:problem_dir/:problem_num", func(c *gin.Context) { controller.DeleteWithFileNumberController(c) })
	Router = router
}
