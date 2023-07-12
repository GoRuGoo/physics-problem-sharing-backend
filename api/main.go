package main

import (
	"physics/infrastructure"
)

func main() {
	infrastructure.InitializeRouter()
	infrastructure.Router.Run(":8080")
}
