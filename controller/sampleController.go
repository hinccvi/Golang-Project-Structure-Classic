package controller

import (
	"github.com/gin-gonic/gin"
	"hinccvi/go-template/log"
	"hinccvi/go-template/request"
	"hinccvi/go-template/service"
	"net/http"
)

func Create(c *gin.Context) {
	var user request.User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"data": struct{}{},
			"msg":  "Invalid param",
		})
		return
	}
	service.Create(user, c)
}

func Read(c *gin.Context) {
	service.Read(c)
}

func Update(c *gin.Context) {
	var user request.UpdateUser
	var userId request.UserId
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"data": struct{}{},
			"msg":  "Invalid param",
		})
		return
	}
	if err := c.ShouldBindQuery(&userId); err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"data": struct{}{},
			"msg":  "Invalid param",
		})
		return
	}
	service.Update(user, userId, c)
}

func Delete(c *gin.Context) {
	var userId request.UserId
	if err := c.ShouldBindQuery(&userId); err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"data": struct{}{},
			"msg":  "Invalid param",
		})
		return
	}
	service.Delete(userId, c)
}
