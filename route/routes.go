package route

import (
	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"
	"hinccvi/go-template/controller"
)

func Init(env string) *gin.Engine {
	if env != "dev" {
		gin.SetMode(gin.ReleaseMode)
	}
	gin.ForceConsoleColor()
	gin.DefaultWriter = colorable.NewColorableStdout()

	e := gin.Default()
	e.POST("/create", controller.Create)
	e.GET("/read", controller.Read)
	e.PATCH("/update", controller.Update)
	e.DELETE("/delete", controller.Delete)

	return e
}
