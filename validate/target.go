package validate

import (
	"fmt"
	"log"

	"github.com/berryfl/alb-api-ng/router"
	"github.com/berryfl/alb-api-ng/target"
	"gorm.io/gorm"
)

func ValidateTargetNoReference(db *gorm.DB, t *target.Target) error {
	routers, err := router.GetRoutersByTarget(db, t.InstanceName, t.Name)
	if err != nil {
		log.Printf("validate_target_no_reference_error: instance_name(%v) name(%v) %v\n", t.InstanceName, t.Name, err)
		return err
	}

	if len(routers) > 0 {
		log.Printf("validate_target_no_reference_failed: instance_name(%v) name(%v) routers(%v)\n", t.InstanceName, t.Name, len(routers))
		return fmt.Errorf("validate_target_no_reference_failed: instance_name(%v) name(%v) routers(%v)", t.InstanceName, t.Name, len(routers))
	}

	return nil
}
