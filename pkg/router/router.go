package router

import (
	"booklib/pkg/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/books", handler.BooksHandler)
	r.GET("/books/:id/content/:page", handler.BookContentHandler)

	return r
}
