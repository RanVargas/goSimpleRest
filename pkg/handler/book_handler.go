package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RespondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func BooksHandler(c *gin.Context) {

	c.String(200, "Listing books")
}

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

	c.String(200, fmt.Sprintf("Getting content for book ID %d, page %d", bookID, page))
}
