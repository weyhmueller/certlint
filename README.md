# certlint

[![Build Status](https://travis-ci.org/weyhmueller/certlint.svg?branch=master)](https://travis-ci.org/weyhmueller/certlint)
[![Go Report Card](https://goreportcard.com/badge/github.com/weyhmueller/certlint)](https://goreportcard.com/report/github.com/weyhmueller/certlint)
[![Coverage Status](http://codecov.io/github/weyhmueller/certlint/coverage.svg?branch=master)](http://codecov.io/github.com/weyhmueller/certlint?branch=master)
[![GoDoc](https://godoc.org/github.com/weyhmueller/certlint?status.svg)](https://godoc.org/github.com/weyhmueller/certlint)

X.509 certificate linter written in Go. Originally developed by Globalsign.

#### General
This package is a work in progress.

Please keep in mind that:
- This is an early release and may contain bugs or false reports
- Not all checks have been fully implemented or verified against the standard
- CLI flag, APIs and CSV export are subject to change

Code contributions and tests are highly welcome!

#### Installation

To install from source, just run:
```bash
go get -u github.com/weyhmueller/certlint
go install github.com/weyhmueller/certlint
```

#### CLI: Usage
The 'certlint' command line utility included with this package can be used to test a single certificate or a large pem container to bulk test millions of certificates. The command is used to test the linter on a large number of certificates but could use fresh up to reduce code complexity.

```
Usage of ./certlint:
  -bulk string
        Bulk certificates file
  -cert string
        Certificate file
  -expired
        Test expired certificates
  -help
        Show this help
  -include
        Include certificates in report
  -issuer string
        Certificate file
  -pprof
        Generate pprof profile
  -report string
        Report filename (default "report.csv")
  -revoked
        Check if certificates are revoked
```

##### CLI: One certificate
```bash
$ certlinter -cert certificate.pem
```

##### CLI: A series of PEM encoded certificates
```bash
$ certlinter -bulk largestore.pem
```

##### CLI: Testing expired certificates
```bash
$ certlinter -expired -bulk largestore.pem
```

##### API: Usage
Import one or all of these packages:

```go
import "github.com/weyhmueller/certlint/asn1"
import "github.com/weyhmueller/certlint/certdata"
import "github.com/weyhmueller/certlint/checks"
```

You can import all available checks:
```go
_ "github.com/weyhmueller/certlint/checks/extensions/all"
_ "github.com/weyhmueller/certlint/checks/certificate/all"
```

Or you can just import a restricted set:
```go
// Check for certificate (ext) KeyUsage extension
_ "github.com/weyhmueller/certlint/checks/extensions/extkeyusage"
_ "github.com/weyhmueller/certlint/checks/extensions/keyusage"

// Also check the parsed certificate (ext) keyusage content
_ "github.com/weyhmueller/certlint/checks/certificate/extkeyusage"
_ "github.com/weyhmueller/certlint/checks/certificate/keyusage"
```

##### API: Check ASN.1 value formatting
```go
al := new(asn1.Linter)
e := al.CheckStruct(der)
if e != nil {
  for _, err := range e.List() {
    fmt.Println(err)
  }
}
```

##### API: Check certificate details
```go
d, err := certdata.Load(der)
if err == nil {
  e := checks.Certificate.Check(d)
  if e != nil {
    for _, err := range e.List() {
      fmt.Println(err)
    }
  }
}
```
