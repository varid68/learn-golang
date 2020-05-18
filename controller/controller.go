package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/varid68/rest-api/database"
	"github.com/varid68/rest-api/model"
)

func GetUsers(c *gin.Context) {
	db := database.GetDB()
	users := []model.User{}

	var sortby string = "name"
	var order string = "asc"
	var limit int = 10
	var offset int = 0
	var count int

	if v := c.Query("sortby"); v != "" {
		sortby = v
	}

	if v := c.Query("order"); v != "" {
		order = v
	}

	_offset := c.Query("offset")
	if v, err := strconv.Atoi(_offset); err == nil {
		offset = v
	}

	_limit := c.Query("limit")
	if v, err := strconv.Atoi(_limit); err == nil {
		limit = v
	}

	db.Find(&users).Count(&count)

	if result := db.Order(sortby + " " + order).Offset(offset).Limit(limit).Find(&users); result.Error != nil {
		c.JSON(403, gin.H{
			"status_code": 403,
			"description": result.GetErrors(),
			"count":       count,
			"offset":      offset,
			"limit":       limit,
			"payload":     nil,
		})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{
		"status_code": 200,
		"description": "success",
		"count":       count,
		"offset":      offset,
		"limit":       limit,
		"payload":     users,
	})
}

func GetUser(c *gin.Context) {
	db := database.GetDB()
	id := c.Params.ByName("id")
	var user model.User

	if db.First(&user, id).Error != nil {
		c.JSON(403, gin.H{
			"status_code": 403,
			"description": "record not found",
			"payload":     nil,
		})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{
		"status_code": 200,
		"description": "success",
		"payload":     user,
	})
}

func CreateUser(c *gin.Context) {
	db := database.GetDB()
	var user model.User

	c.BindJSON(&user)
	db.Create(&user)
	c.JSON(200, user)
}

func DeleteUser(c *gin.Context) {
	db := database.GetDB()
	id := c.Params.ByName("id")
	var user model.User

	db.Delete(&user, id)

	c.JSON(200, gin.H{
		"status_code": 200,
		"description": "success",
		"payload":     nil,
	})
}

func UpdateUser(c *gin.Context) {
	db := database.GetDB()
	id := c.Params.ByName("id")
	var user model.User
	if db.First(&user, id).Error != nil {
		c.AbortWithStatus(404)
	} else {
		c.BindJSON(&user)
		db.Save(&user)
		c.JSON(200, user)
	}
}
