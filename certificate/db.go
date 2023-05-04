package certificate

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"golang.org/x/exp/slices"
	"gorm.io/gorm"
)

func (domains Domains) Value() (driver.Value, error) {
	return json.Marshal(domains)
}

func (domains *Domains) Scan(value any) error {
	valueBytes, ok := value.([]byte)
	if !ok {
		log.Println("convert_content_value_to_bytes_failed")
		return errors.New("convert_content_value_to_bytes_failed")
	}
	return json.Unmarshal(valueBytes, domains)
}

func (c *Certificate) TableName() string {
	return "certificate_tab"
}

func (c *Certificate) Create(db *gorm.DB) error {
	result := db.Create(c)
	if result.Error != nil {
		log.Printf("create_certificate_failed: instance_name(%v) name(%v) %v\n", c.InstanceName, c.Name, result.Error)
		return result.Error
	}
	log.Printf("create_certificate_success: instance_name(%v) name(%v)\n", c.InstanceName, c.Name)
	return nil
}

func (c *Certificate) Extract() error {
	certificates, err := ParseCertificateChain([]byte(c.Chain))
	if err != nil {
		log.Printf("extract_chain_failed: instance_name(%v) name(%v) %v\n", c.InstanceName, c.Name, err)
		return err
	}

	if len(certificates) == 0 {
		log.Printf("extract_chain_empty: instance_name(%v) name(%v)\n", c.InstanceName, c.Name)
		return fmt.Errorf("extract_chain_empty: instance_name(%v) name(%v)", c.InstanceName, c.Name)
	}

	leaf := certificates[0]

	var issuer string
	if len(leaf.Issuer.Organization) > 0 {
		issuer = fmt.Sprintf("%v %v", leaf.Issuer.Organization[0], leaf.Issuer.CommonName)
	} else {
		issuer = leaf.Issuer.CommonName
	}

	domains := slices.Clone(leaf.DNSNames)
	if !slices.Contains(domains, leaf.Subject.CommonName) {
		domains = append(domains, leaf.Subject.CommonName)
	}

	c.Issuer = issuer
	c.Domains = domains
	c.NotBefore = leaf.NotBefore
	c.NotAfter = leaf.NotAfter

	return nil
}

func (c *Certificate) Delete(db *gorm.DB) error {
	result := db.Where("instance_name = ? AND name = ?", c.InstanceName, c.Name).Delete(c)
	if result.Error != nil {
		log.Printf("delete_certificate_failed: instance_name(%v) name(%v) %v\n", c.InstanceName, c.Name, result.Error)
		return result.Error
	}
	log.Printf("delete_certificate_success: affected_rows(%v) instance_name(%v) name(%v)\n", result.RowsAffected, c.InstanceName, c.Name)
	return nil
}

func GetCertificate(db *gorm.DB, instance_name string, name string) (*Certificate, error) {
	var c Certificate
	result := db.Where("instance_name = ? AND name = ?", instance_name, name).First(&c)
	if result.Error != nil {
		log.Printf("get_certificate_failed: instance_name(%v) name(%v) %v\n", instance_name, name, result.Error)
		return nil, result.Error
	}
	return &c, nil
}
