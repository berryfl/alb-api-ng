package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/berryfl/alb-api-ng/database"
	"github.com/berryfl/alb-api-ng/router"
	"github.com/gin-gonic/gin"
)

func CreateRouter(c *gin.Context) {
	var r router.Router
	if err := c.ShouldBindJSON(&r); err != nil {
		log.Printf("create_router: bind_json_failed: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": fmt.Sprintf("create_router: bind_json_failed: %v", err),
		})
		return
	}

	db := database.GetDB()
	if err := r.Validate(db); err != nil {
		c.JSON(http.StatusPreconditionFailed, gin.H{
			"success": false,
			"message": fmt.Sprintf("create_router: validation_failed: %v", err),
		})
		return
	}

	if err := r.Create(db); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": fmt.Sprintf("create_router: create_in_db_failed: %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

type GetRouterReq struct {
	InstanceName string `form:"instance-name" binding:"required"`
	Domain       string `form:"domain" binding:"required"`
}

func GetRouter(c *gin.Context) {
	var req GetRouterReq
	if err := c.ShouldBindQuery(&req); err != nil {
		log.Printf("get_router: bind_query_failed: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": fmt.Sprintf("get_router: bind_query_failed: %v", err),
		})
		return
	}

	db := database.GetDB()
	r, err := router.GetRouter(db, req.InstanceName, req.Domain)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": fmt.Sprintf("get_router: get_from_db_failed: %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"router":  r,
	})
}
