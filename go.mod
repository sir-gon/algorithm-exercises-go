module gon.cl/algorithms

go 1.22.0

require (
	github.com/stretchr/testify v1.10.0
	go.uber.org/zap v1.27.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// Excluded due Snyk vulnerability reports
exclude (
	golang.org/x/crypto v0.37.0
	golang.org/x/net v0.40.0
	golang.org/x/text v0.24.0
)
