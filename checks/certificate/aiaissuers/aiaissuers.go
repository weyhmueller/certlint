package aiaissuers

import (
	"net/url"

	"github.com/weyhmueller/certlint/certdata"
	"github.com/weyhmueller/certlint/checks"
	"github.com/weyhmueller/certlint/errors"
)

const checkName = "Authority Info Access Issuers Check"

func init() {
	checks.RegisterCertificateCheck(checkName, nil, Check)
}

// Check performs a strict verification on the extension according to the standard(s)
func Check(d *certdata.Data) *errors.Errors {
	var e = errors.New(nil)
	if len(d.Cert.IssuingCertificateURL) == 0 {
		e.Err("Certificate contains no Authority Info Access Issuers")
		return e
	}

	for _, icu := range d.Cert.IssuingCertificateURL {
		l, err := url.Parse(icu)
		if err != nil {
			e.Err("Certificate contains an invalid Authority Info Access Issuer URL (%s)", icu)
		}
		if l.Scheme != "http" {
			e.Warning("Certificate contains a Authority Info Access Issuer with an non-preferred scheme (%s)", l.Scheme)
		}
	}

	return e
}
