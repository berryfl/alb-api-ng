package target

import (
	"gorm.io/plugin/soft_delete"
)

type Content struct {
	ZookeeperContent ZookeeperContent `json:"zookeeper_content"`
	DomainContent    DomainContent    `json:"domain_content"`
	IPPortContent    []IPPortContent  `json:"ip_port_content"`
}

type ZookeeperContent struct {
	SDU string `json:"sdu"`
}

type DomainContent struct {
	Domain string `json:"domain"`
	Port   uint16 `json:"port"`
}

type IPPortContent struct {
	IP   string `json:"ip"`
	Port uint16 `json:"port"`
}

type Target struct {
	ID           uint                  `gorm:"column:id" json:"-"`
	InstanceName string                `gorm:"column:instance_name" json:"instance_name" binding:"required"`
	Name         string                `gorm:"column:name" json:"name" binding:"required"`
	TargetType   string                `gorm:"column:target_type" json:"target_type"`
	Content      Content               `gorm:"column:content" json:"content"`
	CreatedAt    int64                 `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    int64                 `gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy    string                `gorm:"column:updated_by" json:"updated_by"`
	DeletedAt    soft_delete.DeletedAt `json:"-"`
}
