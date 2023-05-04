package validate

import (
	"fmt"
	"log"

	"github.com/berryfl/alb-api-ng/certificate"
	"github.com/berryfl/alb-api-ng/router"
	"gorm.io/gorm"
)

func ValidateCertNoReference(db *gorm.DB, cert *certificate.Certificate) error {
	routers, err := router.GetRoutersByCert(db, cert.InstanceName, cert.Name)
	if err != nil {
		log.Printf("validate_cert_no_reference_error: instance_name(%v) name(%v) %v\n", cert.InstanceName, cert.Name, err)
		return err
	}

	if len(routers) > 0 {
		log.Printf("validate_target_no_reference_failed: instance_name(%v) name(%v) routers(%v)\n", cert.InstanceName, cert.Name, len(routers))
		return fmt.Errorf("validate_target_no_reference_failed: instance_name(%v) name(%v) routers(%v)", cert.InstanceName, cert.Name, len(routers))
	}

	return nil
}
