package handler

import (
	"booklib/internal/database"
	"strconv"

	"github.com/gin-gonic/gin"
)

func BookContentHandler(c *gin.Context) {

	bookID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(400, "Invalid book ID")
		return
	}

	page, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		c.String(400, "Invalid page number")
		return
	}
	content, err := database.GetBookContentByID(bookID, page)
	if err != nil {
		c.String(400, "Non existant book ID number")
		return
	}
	c.JSON(200, content)

}
