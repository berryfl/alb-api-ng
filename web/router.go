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
	routerAPI.POST("/delete", DeleteRouter)
	routerAPI.GET("/get", GetRouter)

	targetAPI := r.Group("/target")
	targetAPI.POST("/create", CreateTarget)
	targetAPI.POST("/delete", DeleteTarget)
	targetAPI.GET("/get", GetTarget)

	certificateAPI := r.Group("/certificate")
	certificateAPI.POST("/create", CreateCertificate)
	certificateAPI.POST("/delete", DeleteCertificate)
	certificateAPI.GET("/get", GetCertificate)

	return r
}
