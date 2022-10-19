package controllers

import (
	"go-todo-api/models"
	"go-todo-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func BookAdd(c *gin.Context) {
	// Book構造体をインスタンス化
	book := models.Book{}
	// リクエストのボディをBook構造体にデコードする
	err := c.Bind(&book)
	// エラーハンドリング
	if err != nil {
		c.String(http.StatusBadRequest, "Bad Request")
		return
	}
	// BookService構造体をインスタンス化(フィールド何も持ってない)
	bookService := service.BookService{}
	err = bookService.SetBook(&book)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server ERROR")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "OK",
	})
}

func BookList(c *gin.Context) {
	bookService := service.BookService{}
	BookLists := bookService.GetBookList()
	c.JSONP(http.StatusOK, gin.H{
		"message": "ok",
		"data":    BookLists,
	})
}

func BookUpdate(c *gin.Context) {
	book := models.Book{}
	err := c.Bind(&book)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	bookService := service.BookService{}
	err = bookService.UpdateBook(&book)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}

func BookDelete(c *gin.Context) {
	id := c.PostForm("id")
	intId, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	bookService := service.BookService{}
	err = bookService.DeleteBook(int(intId))
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}
