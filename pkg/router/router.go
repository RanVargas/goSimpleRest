package router

import (
	"booklib/pkg/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/books/:page", handler.BookListHandler)
	r.GET("/book/:id", handler.BookByIdHandler)
	r.GET("/book/:id/content/:page", handler.BookContentHandler)
	r.GET("/seed", handler.InjectSeedHanlder)
	//r.GET("/seed/update", handler.UpdateSeederHandler)
	//r.GET("/seed/files", handler.InjectSeedFromFilesHandler)

	return r
}
