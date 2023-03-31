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

	routerAPI := r.Group("/router")
	routerAPI.POST("/create", CreateRouter)
	routerAPI.GET("/get", GetRouter)

	targetAPI := r.Group("/target")
	targetAPI.POST("/create", CreateTarget)
	targetAPI.GET("/get", GetTarget)

	return r
}
