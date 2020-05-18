package main

import (
	"github.com/gin-gonic/gin"
	"github.com/varid68/rest-api/controller"
	"github.com/varid68/rest-api/database"
)

func main() {
	r := gin.Default()

	database.InitDB()

	v1 := r.Group("/v1")
	{
		v1.GET("/user", controller.GetUsers)
		v1.GET("/user/:id", controller.GetUser)
		v1.POST("/user", controller.CreateUser)
		v1.DELETE("/user/:id", controller.DeleteUser)
		v1.PUT("/user/:id", controller.UpdateUser)
	}
	r.Run()
}
