package extensions

import (
	"github.com/weyhmueller/certlint/certdata"
	"github.com/weyhmueller/certlint/checks"
	"github.com/weyhmueller/certlint/errors"
)

const checkName = "Extensions Check"

func init() {
	checks.RegisterCertificateCheck(checkName, nil, Check)
}

// Check performs a strict verification on the extension according to the standard(s)
func Check(d *certdata.Data) *errors.Errors {
	var e = errors.New(nil)
	for _, ext := range d.Cert.Extensions {
		// Check for any imported extensions and run all matching
		e.Append(checks.Extensions.Check(ext, d))
	}
	return e
}
