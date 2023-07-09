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

	controller := controller.NewWriteController(NewBucketHandler())

	router.PUT("/write/:problem_num", func(c *gin.Context) { controller.WriteWithProblemNumberController(c, "problem_num") })
}
