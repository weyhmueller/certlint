package authoritykeyid

import (
	"crypto/x509/pkix"
	"encoding/asn1"

	"github.com/weyhmueller/certlint/certdata"
	"github.com/weyhmueller/certlint/checks"
	"github.com/weyhmueller/certlint/errors"
)

const checkName = "AuthorityKeyId Extension Check"

var extensionOid = asn1.ObjectIdentifier{2, 5, 29, 35}

func init() {
	checks.RegisterExtensionCheck(checkName, extensionOid, nil, Check)
}

// Check performs a strict verification on the extension according to the standard(s)
func Check(ex pkix.Extension, d *certdata.Data) *errors.Errors {
	var e = errors.New(nil)

	if ex.Critical {
		e.Err("AuthorityKeyId extension set critical")
	}

	return e
}
