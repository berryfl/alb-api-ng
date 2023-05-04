package validate

import (
	"fmt"
	"log"

	"github.com/berryfl/alb-api-ng/certificate"
	"github.com/berryfl/alb-api-ng/router"
	"github.com/berryfl/alb-api-ng/target"
	"gorm.io/gorm"
)

func ValidateRouterTargets(db *gorm.DB, r *router.Router) error {
	var targetNames []string
	for _, rule := range r.Content.Rules {
		if len(rule.TargetName) > 0 {
			targetNames = append(targetNames, rule.TargetName)
		}
	}

	targets, err := target.ListTargets(db, r.InstanceName, targetNames)
	if err != nil {
		return err
	}

	targetMap := make(map[string]bool)
	for _, t := range targets {
		targetMap[t.Name] = true
	}

	var missingTargets []string
	for _, targetName := range targetNames {
		if _, ok := targetMap[targetName]; !ok {
			missingTargets = append(missingTargets, targetName)
		}
	}
	if len(missingTargets) > 0 {
		return fmt.Errorf("no_such_targets: instance_name(%v) targets(%+v)", r.InstanceName, missingTargets)
	}

	return nil
}

func ValidateRouterCert(db *gorm.DB, r *router.Router) error {
	verifyCert := len(r.CertName) > 0
	for _, rule := range r.Content.Rules {
		if rule.UsedInHTTPS {
			verifyCert = true
		}
	}

	if !verifyCert {
		log.Printf("skip_verify_cert: instance_name(%v) domain(%v)\n", r.InstanceName, r.Domain)
		return nil
	}

	cert, err := certificate.GetCertificate(db, r.InstanceName, r.CertName)
	if err != nil {
		return fmt.Errorf("get_cert_error: instance_name(%v) cert_name(%v) %v", r.InstanceName, r.CertName, err)
	}

	if !certificate.IsDomainInCertDomains(cert.Domains, r.Domain) {
		return fmt.Errorf("cert_not_match: instance_name(%v) cert_name(%v)", r.InstanceName, r.CertName)
	}

	return nil
}

func ValidateRouter(db *gorm.DB, r *router.Router) error {
	if err := ValidateRouterTargets(db, r); err != nil {
		log.Printf("validate_router_targets_error: instance_name(%v) domain(%v) %v\n", r.InstanceName, r.Domain, err)
		return err
	}

	if err := ValidateRouterCert(db, r); err != nil {
		log.Printf("validate_router_cert_error: instance_name(%v) domain(%v) %v\n", r.InstanceName, r.Domain, err)
		return err
	}

	return nil
}
