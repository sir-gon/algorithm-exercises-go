module gon.cl/algorithms

go 1.22.0

toolchain go1.23.6

require (
	github.com/stretchr/testify v1.10.0
	go.uber.org/zap v1.27.0
	golang.org/x/exp v0.0.0-20250128182459-e0ece0dbea4c
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// Excluded due Snyk vulnerability reports
exclude (
	golang.org/x/crypto v0.32.0
	golang.org/x/net v0.34.0
	golang.org/x/text v0.22.0
)
