package api

import (
	"assignment3/database"
	"assignment3/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	var db = database.GetDB()

	var data model.Weather
	err := db.Find(&data).Error
	if err != nil {
		fmt.Println("Error getting weather datas : ", err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
	fmt.Printf("%d", data)
}

func Update(c *gin.Context) {
	var db = database.GetDB()

	var data model.Weather
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := db.Model(&model.Weather{}).Where("id = ?", 1).Updates(&data).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})

}
