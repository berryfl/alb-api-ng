package certificate

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
	"strings"

	"golang.org/x/exp/slices"
)

func ParseCertificateChain(chain []byte) ([]*x509.Certificate, error) {
	var result []*x509.Certificate
	for len(chain) > 0 {
		block, rest := pem.Decode(chain)
		if block == nil || block.Type != "CERTIFICATE" {
			return nil, errors.New("pem_decode_chain_failed")
		}

		c, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			return nil, err
		}

		result = append(result, c)
		chain = rest
	}
	return result, nil
}

func ConvertToWildcardDomain(domain string) string {
	parts := strings.Split(domain, ".")
	if len(parts) > 1 {
		parts[0] = "*"
	}
	return strings.Join(parts, ".")
}

func IsDomainInCertDomains(certDomains []string, domain string) bool {
	if ok := slices.Contains(certDomains, domain); ok {
		return true
	}

	wildcard := ConvertToWildcardDomain(domain)
	if ok := slices.Contains(certDomains, wildcard); ok {
		return true
	}

	return false
}
