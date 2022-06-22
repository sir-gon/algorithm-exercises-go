module gon.cl/projecteuler.net

go 1.18

require (
	github.com/stretchr/testify v1.7.4
	go.uber.org/zap v1.21.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// Excluded due Snyk vulnerability reports
exclude (
	golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2
	golang.org/x/crypto v0.0.0-20191011191535-87dc89f01550
	golang.org/x/net v0.0.0-20190620200207-3b0461eec859
	golang.org/x/net v0.0.0-20210405180319-a5a99cb37ef4
	golang.org/x/text v0.3.0
	gopkg.in/yaml.v3 v3.0.0
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
)
