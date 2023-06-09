package target

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

func (t *Target) TableName() string {
	return "target_tab"
}

func (t *Target) Create(db *gorm.DB) error {
	result := db.Create(t)
	if result.Error != nil {
		log.Printf("create_target_failed: instance_name(%v) name(%v) %v\n", t.InstanceName, t.Name, result.Error)
		return result.Error
	}
	log.Printf("create_target_success: id(%v) instance_name(%v) name(%v)\n", t.ID, t.InstanceName, t.Name)
	return nil
}

func (t *Target) Delete(db *gorm.DB) error {
	result := db.Where("instance_name = ? AND name = ?", t.InstanceName, t.Name).Delete(t)
	if result.Error != nil {
		log.Printf("delete_target_failed: instance_name(%v) name(%v) %v\n", t.InstanceName, t.Name, result.Error)
		return result.Error
	}
	log.Printf("delete_target_success: affected_rows(%v) instance_name(%v) name(%v)\n", result.RowsAffected, t.InstanceName, t.Name)
	return nil
}

func GetTarget(db *gorm.DB, instance_name string, name string) (*Target, error) {
	var t Target
	result := db.Where("instance_name = ? AND name = ?", instance_name, name).First(&t)
	if result.Error != nil {
		log.Printf("get_target_failed: instance_name(%v) name(%v) %v\n", instance_name, name, result.Error)
		return nil, result.Error
	}
	return &t, nil
}

func ListTargets(db *gorm.DB, instance_name string, names []string) ([]*Target, error) {
	var targets []*Target
	result := db.Where("instance_name = ? AND name IN ?", instance_name, names).Find(&targets)
	if result.Error != nil {
		log.Printf("list_targets_failed: instance_name(%v) names(%v) %v\n", instance_name, len(names), result.Error)
		return nil, result.Error
	}
	return targets, nil
}
