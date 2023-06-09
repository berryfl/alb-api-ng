package router

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"log"

	"github.com/berryfl/alb-api-ng/datatypes"
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

func (r *Router) Delete(db *gorm.DB) error {
	result := db.Where("instance_name = ? AND domain = ?", r.InstanceName, r.Domain).Delete(r)
	if result.Error != nil {
		log.Printf("delete_router_failed: instance_name(%v) domain(%v) %v\n", r.InstanceName, r.Domain, result.Error)
		return result.Error
	}
	log.Printf("delete_router_success: instance_name(%v) domain(%v)\n", r.InstanceName, r.Domain)
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

func GetRoutersByTarget(db *gorm.DB, instanceName string, targetName string) ([]*Router, error) {
	var routers []*Router
	result := db.Where(
		"instance_name = ?",
		instanceName,
	).Find(&routers, datatypes.JSONArrayQuery("content", "$.rules[*].target_name").HasValue(targetName))
	if result.Error != nil {
		log.Printf("get_routers_by_target_failed: instance_name(%v) target(%v) %v\n", instanceName, targetName, result.Error)
		return nil, result.Error
	}
	return routers, nil
}

func GetRoutersByCert(db *gorm.DB, instanceName string, certName string) ([]*Router, error) {
	var routers []*Router
	result := db.Where("cert_name = ?", certName).Find(&routers)
	if result.Error != nil {
		log.Printf("get_routers_by_cert_failed: instance_name(%v) cert(%v) %v\n", instanceName, certName, result.Error)
		return nil, result.Error
	}
	return routers, nil
}
