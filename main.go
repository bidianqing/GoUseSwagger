package main

import (
	"github.com/bidianqing/go-use-swagger/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	app := gin.Default()
	docs.SwaggerInfo.BasePath = ""

	app.GET("/", Index)

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	app.Run()
}

// ShowAccount godoc
//
//	@Summary		Show an account
//	@Description	get string by ID
//	@Router			/ [get]
func Index(c *gin.Context) {
	c.JSON(200, gin.H{"success": true, "message": "ok", "data": nil})
}
