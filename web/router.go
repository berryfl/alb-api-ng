package web

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	instanceAPI := r.Group("/instance")
	instanceAPI.POST("/create", CreateInstance)
	instanceAPI.POST("/delete", DeleteInstance)
	instanceAPI.POST("/update", UpdateInstance)
	instanceAPI.GET("/get", GetInstance)
	return r
}
