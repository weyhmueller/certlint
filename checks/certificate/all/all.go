package all

import (
	// Import all default checks
	_ "github.com/weyhmueller/certlint/checks/certificate/aiaissuers"
	_ "github.com/weyhmueller/certlint/checks/certificate/basicconstraints"
	_ "github.com/weyhmueller/certlint/checks/certificate/extensions"
	_ "github.com/weyhmueller/certlint/checks/certificate/extkeyusage"
	_ "github.com/weyhmueller/certlint/checks/certificate/internal"
	_ "github.com/weyhmueller/certlint/checks/certificate/issuerdn"
	_ "github.com/weyhmueller/certlint/checks/certificate/keyusage"
	_ "github.com/weyhmueller/certlint/checks/certificate/publickey"
	_ "github.com/weyhmueller/certlint/checks/certificate/publicsuffix"
	_ "github.com/weyhmueller/certlint/checks/certificate/revocation"
	_ "github.com/weyhmueller/certlint/checks/certificate/serialnumber"
	_ "github.com/weyhmueller/certlint/checks/certificate/signaturealgorithm"
	_ "github.com/weyhmueller/certlint/checks/certificate/subject"
	_ "github.com/weyhmueller/certlint/checks/certificate/subjectaltname"
	_ "github.com/weyhmueller/certlint/checks/certificate/validity"
	_ "github.com/weyhmueller/certlint/checks/certificate/version"
	_ "github.com/weyhmueller/certlint/checks/certificate/wildcard"
)
