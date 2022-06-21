package main

import (
	"github.com/gin-gonic/gin"

	rg "gin-webservice/routes"
)

func main() {
	RunServer()
}

func RunServer() {
	r := gin.New()

	apiGroups := r.Group("/api")
	{
		rg.Routes(apiGroups)
	}

	r.Run(":5000")
}
