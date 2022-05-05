package routers

import (
	v1Controller "hinccvi/go-template/controllers/api/v1"
	"hinccvi/go-template/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"
)

func Init(env string) *gin.Engine {
	if env != "dev" {
		gin.SetMode(gin.ReleaseMode)
	}
	gin.ForceConsoleColor()
	gin.DefaultWriter = colorable.NewColorableStdout()

	e := gin.Default()
	v1 := e.Group("/v1")
	v1.Use(middlewares.UserMiddlewares())
	{
		v1.POST("/create", v1Controller.Create)
		v1.GET("/read", v1Controller.Read)
		v1.PATCH("/update", v1Controller.Update)
		v1.DELETE("/delete", v1Controller.Delete)
	}

	return e
}
