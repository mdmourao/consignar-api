package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mdmourao/irs-api/db"
	"github.com/mdmourao/irs-api/rest"
)

func main() {
	db.InitConnection()
	// db.Migrate()
	// db.PopulateLocalities()
	// db.PopulateEntities()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})

	})

	router.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})

	})

	router.GET("/localities", rest.GetLocalities)
	router.GET("/localities/:id", rest.GetLocalityById)
	router.GET("/localities/:id/entities", rest.GetEntitiesFromLocalityId)

	router.GET("/entities/:id", rest.GetEntityById)

	router.Run(":50007")
}
