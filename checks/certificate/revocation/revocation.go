package revocation

import (
	"net/url"

	"github.com/weyhmueller/certlint/certdata"
	"github.com/weyhmueller/certlint/checks"
	"github.com/weyhmueller/certlint/errors"
)

const checkName = "Certificate Revocation Information Check"

func init() {
	checks.RegisterCertificateCheck(checkName, nil, Check)
}

// Check performs a strict verification on the extension according to the standard(s)
func Check(d *certdata.Data) *errors.Errors {
	var e = errors.New(nil)

	if len(d.Cert.CRLDistributionPoints) == 0 && len(d.Cert.OCSPServer) == 0 {
		e.Err("Certificate contains no CRL or OCSP server")
		return e
	}

	// Check CRL information
	for _, crl := range d.Cert.CRLDistributionPoints {
		l, err := url.Parse(crl)
		if err != nil {
			e.Err("Certificate contains an invalid CRL (%s)", crl)
		} else if l.Scheme != "http" {
			e.Err("Certificate contains a CRL with an non-preferred scheme (%s)", l.Scheme)
		}
	}

	// Check OCSP information
	for _, server := range d.Cert.OCSPServer {
		s, err := url.Parse(server)
		if err != nil {
			e.Err("Certificate contains an invalid OCSP server (%s)", s)
		} else if s.Scheme != "http" {
			e.Err("Certificate contains a OCSP server with an non-preferred scheme (%s)", s.Scheme)
		}
	}

	return e
}
