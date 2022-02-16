package service

import (
	"github.com/gin-gonic/gin"
	"hinccvi/go-template/repository"
	"hinccvi/go-template/request"
	"net/http"
)

func Create(user request.User, c *gin.Context) {
	err, _ := repository.CreateUser(user)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"data": err.Error(),
			"msg":  "Create user fail",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"data": struct{}{},
			"msg":  "Success",
		})
	}
}

func Read(c *gin.Context) {
	users := repository.GetAllUser()
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": users,
		"msg":  "Success",
	})
}

func Update(user request.UpdateUser, userId request.UserId, c *gin.Context) {
	userModel := repository.GetUserById(userId.Id)
	if userModel.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"data": struct{}{},
			"msg":  "User not found",
		})
		return
	}

	err := repository.UpdateUser(userModel, user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"data": err.Error(),
			"msg":  "Update user fail",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"data": struct{}{},
			"msg":  "Success",
		})
	}
}

func Delete(userId request.UserId, c *gin.Context) {
	userModel := repository.GetUserById(userId.Id)
	if userModel.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"data": struct{}{},
			"msg":  "User not found",
		})
		return
	}

	err := repository.DeleteUser(userModel)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"data": err,
			"msg":  "Delete user fail",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"data": struct{}{},
			"msg":  "Success",
		})
	}
}
