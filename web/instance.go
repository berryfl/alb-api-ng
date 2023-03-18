package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/berryfl/alb-api-ng/database"
	"github.com/berryfl/alb-api-ng/instance"
	"github.com/gin-gonic/gin"
)

func CreateInstance(c *gin.Context) {
	var inst instance.Instance
	if err := c.ShouldBindJSON(&inst); err != nil {
		log.Printf("create_instance: bind_json_failed: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": fmt.Sprintf("create_instance: bind_json_failed: %v", err),
		})
		return
	}

	db := database.GetDB()
	if err := inst.Create(db); err != nil {
		log.Printf("create_instance: create_in_db_failed: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": fmt.Sprintf("create_instance: create_in_db_failed: %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func DeleteInstance(c *gin.Context) {
	var inst instance.Instance
	if err := c.ShouldBindJSON(&inst); err != nil {
		log.Printf("delete_instance: bind_json_failed: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": fmt.Sprintf("delete_instance: bind_json_failed: %v", err),
		})
		return
	}

	db := database.GetDB()
	if err := inst.Delete(db); err != nil {
		log.Printf("delete_instance: delete_in_db_failed: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": fmt.Sprintf("delete_instance: delete_in_db_failed: %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

type GetInstanceReq struct {
	Name string `form:"name" binding:"required"`
}

func GetInstance(c *gin.Context) {
	var req GetInstanceReq
	if err := c.ShouldBindQuery(&req); err != nil {
		log.Printf("get_instance: bind_query_failed: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": fmt.Sprintf("get_instance: bind_query_failed: %v", err),
		})
		return
	}

	db := database.GetDB()
	inst, err := instance.GetInstance(db, req.Name)
	if err != nil {
		log.Printf("get_instance: get_from_db_failed: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": fmt.Sprintf("get_instance: get_from_db_failed: %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"instance": inst,
	})
}
