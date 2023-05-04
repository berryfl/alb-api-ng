package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/berryfl/alb-api-ng/certificate"
	"github.com/berryfl/alb-api-ng/database"
	"github.com/berryfl/alb-api-ng/validate"
	"github.com/gin-gonic/gin"
)

func CreateCertificate(c *gin.Context) {
	var cert certificate.Certificate
	if err := c.ShouldBindJSON(&cert); err != nil {
		log.Printf("create_certificate: bind_json_failed: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": fmt.Sprintf("create_certificate: bind_json_failed: %v", err),
		})
		return
	}

	if err := cert.Extract(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": fmt.Sprintf("create_certificate: extract_chain_failed: %v", err),
		})
		return
	}

	db := database.GetDB()
	if err := cert.Create(db); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": fmt.Sprintf("create_certificate: create_in_db_failed: %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func DeleteCertificate(c *gin.Context) {
	var cert certificate.Certificate
	if err := c.ShouldBindJSON(&cert); err != nil {
		log.Printf("delete_certificate: bind_json_failed: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": fmt.Sprintf("delete_certificate: bind_json_failed: %v", err),
		})
		return
	}

	db := database.GetDB()
	tx := db.Begin()

	if err := validate.ValidateCertNoReference(tx, &cert); err != nil {
		tx.Rollback()
		c.JSON(http.StatusPreconditionFailed, gin.H{
			"success": false,
			"message": fmt.Sprintf("delete_certificate: validation_failed: %v", err),
		})
		return
	}

	if err := cert.Delete(tx); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": fmt.Sprintf("delete_certificate: delete_in_db_failed: %v", err),
		})
		return
	}

	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": fmt.Sprintf("delete_certificate: db_commit_error: %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

type GetCertificateReq struct {
	InstanceName string `form:"instance-name" binding:"required"`
	Name         string `form:"name" binding:"required"`
}

func GetCertificate(c *gin.Context) {
	var req GetCertificateReq
	if err := c.ShouldBindQuery(&req); err != nil {
		log.Printf("get_target: bind_query_failed: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": fmt.Sprintf("get_target: bind_query_failed: %v", err),
		})
		return
	}

	db := database.GetDB()
	t, err := certificate.GetCertificate(db, req.InstanceName, req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": fmt.Sprintf("get_certificate: get_from_db_failed: %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":     true,
		"certificate": t,
	})
}
