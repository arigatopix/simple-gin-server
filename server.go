package main

import (
	"github.com/gin-gonic/gin"

	db "gin-webservice/db"
	rg "gin-webservice/routes"
)

func main() {
	db.ConnectDB()
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
