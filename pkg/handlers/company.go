package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gotoolkit/cms/pkg/database"
	"github.com/gotoolkit/cms/pkg/model"
)

type company struct {
}

func (c *company) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			abortWithStatus(c, http.StatusBadRequest, fmt.Sprintf("params id format not correct: %v", err))
			return
		}
		data := model.PurchaseCompany{Id: id}
		has, err := database.GetDB().Get(&data)
		if err != nil {
			abortWithStatus(c, http.StatusBadRequest, fmt.Sprintf("failed to get company with %d: %v", id, err))
			return
		}

		if !has {
			abortWithStatus(c, http.StatusBadRequest, fmt.Sprintf("not have company with %d: %v", id, err))
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})

	}
}

func (c *company) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (c *company) Create() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (c *company) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
