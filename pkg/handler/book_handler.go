package handler

import (
	"booklib/internal/database"
	"booklib/pkg/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func BookListHandler(c *gin.Context) {
	page, err := strconv.Atoi(c.Param("page"))
	books := make([]model.Book, 0)
	if err != nil || page < 0 {
		result, rtvErr := database.GetBooks(1)
		if rtvErr != nil {
			c.String(500, "The Server has unexpectedly no respond")
			return
		}
		books = append(books, result...)
	} else {
		result, rtvErr := database.GetBooks(page)
		if rtvErr != nil {
			c.String(500, "The Server has unexpectedly no respond")
			return
		}
		books = append(books, result...)
	}

	c.JSON(200, books)
}

func BookByIdHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 0 {
		c.String(400, "This is a bad request")
		return
	}
	result, rtvErr := database.GetBookByID(id)
	if rtvErr != nil {
		c.String(500, "The Server has unexpectedly no respond")
		return
	}
	c.JSON(200, result)
}
