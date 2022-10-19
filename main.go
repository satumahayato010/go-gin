package main

import (
	"html/template"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

func nl2br(text string) template.HTML {
	return template.HTML(strings.Replace(template.HTMLEscapeString(text), "\n", "<br />", -1))
}

func main() {
	router := gin.Default()
	router.SetFuncMap(template.FuncMap{
		"nl2br": nl2br,
	})

	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	router.POST("/result", func(c *gin.Context) {
		result, ok := c.GetPostForm("result")
		if ok != true {
			log.Fatal("ERROR")
		}
		c.HTML(200, "result.html", result)
	})

	router.Run(":8080")
}
