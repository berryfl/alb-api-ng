package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/berryfl/alb-api-ng/database"
	"github.com/berryfl/alb-api-ng/target"
	"github.com/berryfl/alb-api-ng/validate"
	"github.com/gin-gonic/gin"
)

func CreateTarget(c *gin.Context) {
	var t target.Target
	if err := c.ShouldBindJSON(&t); err != nil {
		log.Printf("create_target: bind_json_failed: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": fmt.Sprintf("create_target: bind_json_failed: %v", err),
		})
		return
	}

	db := database.GetDB()
	if err := t.Create(db); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": fmt.Sprintf("create_target: create_in_db_failed: %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func DeleteTarget(c *gin.Context) {
	var t target.Target
	if err := c.ShouldBindJSON(&t); err != nil {
		log.Printf("delete_target: bind_json_failed: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": fmt.Sprintf("delete_target: bind_json_failed: %v", err),
		})
		return
	}

	db := database.GetDB()
	tx := db.Begin()

	if err := validate.ValidateTargetNoReference(tx, &t); err != nil {
		tx.Rollback()
		c.JSON(http.StatusPreconditionFailed, gin.H{
			"success": false,
			"message": fmt.Sprintf("delete_target: validation_failed: %v", err),
		})
		return
	}

	if err := t.Delete(tx); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": fmt.Sprintf("delete_target: delete_in_db_failed: %v", err),
		})
		return
	}

	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": fmt.Sprintf("delete_target: db_commit_error: %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

type GetTargetReq struct {
	InstanceName string `form:"instance-name" binding:"required"`
	Name         string `form:"name" binding:"required"`
}

func GetTarget(c *gin.Context) {
	var req GetTargetReq
	if err := c.ShouldBindQuery(&req); err != nil {
		log.Printf("get_target: bind_query_failed: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": fmt.Sprintf("get_target: bind_query_failed: %v", err),
		})
		return
	}

	db := database.GetDB()
	t, err := target.GetTarget(db, req.InstanceName, req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": fmt.Sprintf("get_target: get_from_db_failed: %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"target":  t,
	})
}
