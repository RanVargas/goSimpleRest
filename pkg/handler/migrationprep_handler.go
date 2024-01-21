package handler

import (
	"booklib/internal/seeders"

	"github.com/gin-gonic/gin"
)

func UpdateSeederHandler(c *gin.Context) {
	err := seeders.CaptureDbData()
	if err != nil {
		c.String(500, "Operation failed")
	}
	c.String(200, "Seed has been updated")
}

func InjectSeedHanlder(c *gin.Context) {
	err := seeders.SeedBooks()
	if err != nil {
		c.String(500, "An error has ocurred"+err.Error())
		return
	}
	c.String(200, "Seed has been injected to DB")
}

func InjectUpdatedSeed(c *gin.Context) {
	seeders.SeedBooksFromFile()
	c.String(200, "Seed has been planted from files")
}

func InjectSeedFromFilesHandler(c *gin.Context) {
	err := seeders.SeedBooksFromFile()
	if err != nil {
		c.String(500, "Operation failed")
		return
	}
	c.String(201, "Seed injected successfully")
}
