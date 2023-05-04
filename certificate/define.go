package certificate

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

type Domains []string

type Certificate struct {
	ID           uint                  `gorm:"column:id" json:"-"`
	InstanceName string                `gorm:"column:instance_name" json:"instance_name" binding:"required"`
	Name         string                `gorm:"column:name" json:"name" binding:"required"`
	Domains      Domains               `gorm:"column:domains" json:"domains"`
	Issuer       string                `gorm:"column:issuer" json:"issuer"`
	NotBefore    time.Time             `gorm:"column:not_before" json:"not_before"`
	NotAfter     time.Time             `gorm:"column:not_after" json:"not_after"`
	Chain        string                `gorm:"column:chain" json:"chain"`
	Key          string                `gorm:"column:key" json:"-"`
	UpdatedBy    string                `gorm:"column:updated_by" json:"updated_by"`
	CreatedAt    int64                 `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    int64                 `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt    soft_delete.DeletedAt `json:"-"`
}
