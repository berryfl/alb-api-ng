package router

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"log"

	"gorm.io/gorm"
)

func (c Content) Value() (driver.Value, error) {
	return json.Marshal(c)
}

func (c *Content) Scan(value any) error {
	valueBytes, ok := value.([]byte)
	if !ok {
		log.Println("convert_content_value_to_bytes_failed")
		return errors.New("convert_content_value_to_bytes_failed")
	}
	return json.Unmarshal(valueBytes, c)
}

func (r *Router) TableName() string {
	return "router_tab"
}

func (r *Router) Create(db *gorm.DB) error {
	result := db.Create(r)
	if result.Error != nil {
		log.Printf("create_router_failed: instance_name(%v) domain(%v) %v\n", r.InstanceName, r.Domain, result.Error)
		return result.Error
	}
	log.Printf("create_router_success: id(%v) instance_name(%v) domain(%v)\n", r.ID, r.InstanceName, r.Domain)
	return nil
}

func GetRouter(db *gorm.DB, instance_name string, domain string) (*Router, error) {
	var r Router
	result := db.Where("instance_name = ? AND domain = ?", instance_name, domain).First(&r)
	if result.Error != nil {
		log.Printf("get_router_failed: instance_name(%v) domain(%v) %v\n", instance_name, domain, result.Error)
		return nil, result.Error
	}
	return &r, nil
}
