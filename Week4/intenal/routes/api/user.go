package api

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"week4/intenal/service"
)

type User struct{}

func NewUser() User {
	return User{}
}

func (u User) GetNameById(c *gin.Context) {

	id := c.Query("id")
	if len(id) == 0 {
		c.JSON(200, gin.H{
			"err": "请传递id",
		})
		return
	}
	ids, _ := strconv.Atoi(id)
	var name, err = service.GetNameById(ids)
	if err != nil {
		c.JSON(200, gin.H{
			"err": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"name": name,
	})
	return
}
