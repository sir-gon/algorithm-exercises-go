module gon.cl/algorithm-exercises

go 1.20

require (
	github.com/stretchr/testify v1.9.0
	go.uber.org/zap v1.27.0
	golang.org/x/exp v0.0.0-20240506185415-9bf2ced13842
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	gopkg.in/yaml.v3 v3.0.0 // indirect
)

// Excluded due Snyk vulnerability reports
exclude (
	golang.org/x/crypto v0.23.0
	golang.org/x/net v0.25.0
	golang.org/x/net v0.25.0
	golang.org/x/net v0.25.0
	golang.org/x/text v0.15.0
	gopkg.in/yaml.v3 v3.0.1
)
