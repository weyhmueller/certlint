package all

import (
	// Import all default extensions
	_ "github.com/weyhmueller/certlint/checks/extensions/authorityinfoaccess"
	_ "github.com/weyhmueller/certlint/checks/extensions/authoritykeyid"
	_ "github.com/weyhmueller/certlint/checks/extensions/basicconstraints"
	_ "github.com/weyhmueller/certlint/checks/extensions/crldistributionpoints"
	_ "github.com/weyhmueller/certlint/checks/extensions/ct"
	_ "github.com/weyhmueller/certlint/checks/extensions/extkeyusage"
	_ "github.com/weyhmueller/certlint/checks/extensions/keyusage"
	_ "github.com/weyhmueller/certlint/checks/extensions/nameconstraints"
	_ "github.com/weyhmueller/certlint/checks/extensions/policyidentifiers"
	_ "github.com/weyhmueller/certlint/checks/extensions/subjectaltname"
	_ "github.com/weyhmueller/certlint/checks/extensions/subjectkeyid"
)
