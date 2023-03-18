package instance

import (
	"log"

	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

type Instance struct {
	ID          uint
	Name        string
	EnableHTTP  bool
	EnableHTTPS bool `gorm:"column:enable_https"`
	CertName    string
	CreatedAt   int64
	UpdatedAt   int64
	UpdatedBy   string
	DeletedAt   soft_delete.DeletedAt
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
