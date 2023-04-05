package validate

import (
	"fmt"
	"log"

	"github.com/berryfl/alb-api-ng/router"
	"github.com/berryfl/alb-api-ng/target"
	"gorm.io/gorm"
)

func ValidateRouter(db *gorm.DB, r *router.Router) error {
	var targetNames []string
	for _, rule := range r.Content.Rules {
		if len(rule.TargetName) > 0 {
			targetNames = append(targetNames, rule.TargetName)
		}
	}

	targets, err := target.ListTargets(db, r.InstanceName, targetNames)
	if err != nil {
		log.Printf("validate_router_error: instance_name(%v) domain(%v) %v", r.InstanceName, r.Domain, err)
		return err
	}

	targetMap := make(map[string]bool)
	for _, t := range targets {
		targetMap[t.Name] = true
	}

	missingTarget := false
	for _, targetName := range targetNames {
		if _, ok := targetMap[targetName]; !ok {
			log.Printf("target_non_existent: instance_name(%v) domain(%v) target(%v)\n", r.InstanceName, r.Domain, targetName)
			missingTarget = true
		}
	}
	if missingTarget {
		return fmt.Errorf("target_non_existent: instance_name(%v) domain(%v)", r.InstanceName, r.Domain)
	}

	return nil
}
