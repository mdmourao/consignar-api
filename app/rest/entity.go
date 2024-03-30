package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mdmourao/irs-api/db"
)

func GetEntitiesFromLocalityId(c *gin.Context) {
	id, err := GetIdFromParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
	}

	limit, err := GetLimit(c, 50)
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

	entities, err := db.GetEntitiesFromLocalityId(id, limit, offset)

	count := db.GetEntitiesCount(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "server error",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"count": count,
		"data":  entities,
	})

}

func GetEntityById(c *gin.Context) {
	id, err := GetIdFromParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
	}

	entity, err := db.GetEntityById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "server error",
		})
	}

	c.JSON(http.StatusOK, entity)
}
