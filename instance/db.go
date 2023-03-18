package instance

import (
	"log"

	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

type Instance struct {
	ID          uint                  `gorm:"column:id" json:"-"`
	Name        string                `gorm:"column:name" json:"name"`
	EnableHTTP  bool                  `gorm:"column:enable_http" json:"enable_http"`
	EnableHTTPS bool                  `gorm:"column:enable_https" json:"enable_https`
	CertName    string                `gorm:"column:cert_name" json:"cert_name"`
	CreatedAt   int64                 `gorm:"column:created_at" json:"created_at`
	UpdatedAt   int64                 `gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy   string                `gorm:"column:updated_by" json:"updated_by"`
	DeletedAt   soft_delete.DeletedAt `json:"-"`
}

func (inst *Instance) TableName() string {
	return "instance_tab"
}

func (inst *Instance) Create(db *gorm.DB) error {
	result := db.Create(inst)
	if result.Error != nil {
		log.Printf("create_instance_failed: %v\n", result.Error)
		return result.Error
	}
	log.Printf("create_instance_success: id(%v) name(%v)\n", inst.ID, inst.Name)
	return nil
}

func (inst *Instance) Delete(db *gorm.DB) error {
	result := db.Where("name = ?", inst.Name).Delete(inst)
	if result.Error != nil {
		log.Printf("delete_instance_failed: %v\n", result.Error)
	}
	log.Printf("delete_instance_success: affected_rows(%v) name(%v)\n", result.RowsAffected, inst.Name)
	return nil
}
