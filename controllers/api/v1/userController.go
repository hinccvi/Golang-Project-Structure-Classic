package v1Controllers

import (
	"hinccvi/go-template/helpers"
	"hinccvi/go-template/log"
	request "hinccvi/go-template/resources/api/v1"
	service "hinccvi/go-template/service/api/v1"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Create(c *gin.Context) {
	var user request.User

	if err := c.ShouldBindJSON(&user); err != nil {
		if err := c.ShouldBindJSON(&user); err != nil {
			log.Error("Fail to bind request body to struct", zap.Error(err))
			helpers.RespWithBadRequest(c, err.Error(), struct{}{})
			return
		}
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
		log.Error("Fail to bind request body to struct", zap.Error(err))
		helpers.RespWithBadRequest(c, err.Error(), struct{}{})
		return
	}

	if err := c.ShouldBindQuery(&userId); err != nil {
		log.Error("Fail to bind query string to struct", zap.Error(err))
		helpers.RespWithBadRequest(c, err.Error(), struct{}{})
		return
	}

	service.Update(user, userId, c)
}

func Delete(c *gin.Context) {
	var userId request.UserId
	if err := c.ShouldBindQuery(&userId); err != nil {
		log.Error("Fail to bind query string to struct", zap.Error(err))
		helpers.RespWithBadRequest(c, err.Error(), struct{}{})
		return
	}

	service.Delete(userId, c)
}
