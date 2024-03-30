package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mdmourao/irs-api/db"
)

func GetLocalities(c *gin.Context) {
	limit, err := GetLimit(c, 1000)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
	}

	offset, err := GetOffset(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
	}

	search := GetSeach(c)

	localities, err := db.GetLocalities(limit, offset, search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "server error",
		})
	}

	c.JSON(http.StatusOK, localities)
}

func GetLocalityById(c *gin.Context) {
	id, err := GetIdFromParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
	}

	locality, err := db.GetLocalityById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "server error",
		})
	}

	c.JSON(http.StatusOK, locality)
}
