package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gotoolkit/cms/pkg/database"
	"github.com/gotoolkit/cms/pkg/model"
)

func Get(v interface{}) gin.HandlerFunc {
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

func GetAll(v interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func Create(v interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		file, _ := c.FormFile("file")
		if err := c.SaveUploadedFile(file, file.Filename); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "file uploaded",
		})
	}
}

func Delete(v interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
