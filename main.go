package main

import (
	"go-todo-api/controllers"
	"go-todo-api/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.Use(middleware.RecordUaAndTime)

	bookEngine := engine.Group("/book")
	{
		v1 := bookEngine.Group("/v1")
		{
			v1.POST("/add", controllers.BookAdd)
			v1.GET("/list", controllers.BookList)
			v1.PUT("/update", controllers.BookUpdate)
			v1.DELETE("/delete", controllers.BookDelete)
		}
	}
	engine.Run(":8080")
}
