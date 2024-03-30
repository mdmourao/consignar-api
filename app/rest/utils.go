package rest

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetLimit(c *gin.Context, defaultValue int) (int, error) {
	var err error
	limitStr := c.Query("limit")
	fmt.Println(limitStr)
	limit := defaultValue

	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			return 0, fmt.Errorf("invalid limit")
		}
		if limit > defaultValue {
			limit = defaultValue
		}
	}
	return limit, nil
}

func GetOffset(c *gin.Context) (int, error) {
	var err error
	offsetStr := c.Query("offset")
	offset := 0

	if offsetStr != "" {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			return 0, fmt.Errorf("invalid offset")
		}
		if offset < 0 {
			return 0, fmt.Errorf("invalid offset")
		}
	}
	return offset, nil
}

func GetSeach(c *gin.Context) string {
	return c.Query("search")
}

func GetIdFromParams(c *gin.Context) (int, error) {
	var err error
	id := 0
	idStr := c.Param("id")
	fmt.Println(idStr)
	if idStr == "" {
		return 0, fmt.Errorf("not a valid id")
	} else {
		id, err = strconv.Atoi(idStr)
		if err != nil {
			return 0, fmt.Errorf("not a valid id")
		}
		return id, nil
	}
}
