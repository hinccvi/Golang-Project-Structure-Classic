package v1Services

import (
	"hinccvi/go-template/helpers"
	"hinccvi/go-template/model"
	"hinccvi/go-template/repository"
	request "hinccvi/go-template/resources/api/v1"

	"github.com/gin-gonic/gin"
)

func Create(user request.User, c *gin.Context) {
	_, err := repository.CreateUser(user)

	if err != nil {
		if err := c.ShouldBindJSON(&user); err != nil {
			helpers.RespWithSystemError(c, err.Error(), struct{}{})
		}
	} else {
		helpers.RespWithOK(c, "Success", struct{}{})
	}
}

func Read(c *gin.Context) {
	users := repository.GetAllUser()

	helpers.RespWithOK(c, "Success", struct {
		Users []model.User `json:"users"`
	}{
		Users: users,
	})
}

func Update(user request.UpdateUser, userId request.UserId, c *gin.Context) {
	userModel := repository.GetUserById(userId.Id)

	if userModel.ID == 0 {
		helpers.RespWithBadRequest(c, "User not found", struct{}{})
		return
	}

	err := repository.UpdateUser(userModel, user)
	if err != nil {
		helpers.RespWithSystemError(c, err.Error(), struct{}{})
	} else {
		helpers.RespWithOK(c, "Success", struct{}{})
	}
}

func Delete(userId request.UserId, c *gin.Context) {
	userModel := repository.GetUserById(userId.Id)
	if userModel.ID == 0 {
		helpers.RespWithBadRequest(c, "User not found", struct{}{})
		return
	}

	err := repository.DeleteUser(userModel)
	if err != nil {
		helpers.RespWithSystemError(c, err.Error(), struct{}{})
	} else {
		helpers.RespWithOK(c, "Success", struct{}{})
	}
}
