package router

import (
	"PowerTranslate/apis"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/api/translate", apis.TransApi)

	r.Use(cors.Default())

	return r
}
