module gon.cl/algorithm-exercises

go 1.20

require (
	github.com/stretchr/testify v1.9.0
	go.uber.org/zap v1.27.0
	golang.org/x/exp v0.0.0-20240404231335-c0f41cb1a7a0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// Excluded due Snyk vulnerability reports
exclude (
	golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2
	golang.org/x/crypto v0.0.0-20191011191535-87dc89f01550
	golang.org/x/net v0.0.0-20190311183353-d8887717615a
	golang.org/x/net v0.0.0-20190620200207-3b0461eec859
	golang.org/x/net v0.0.0-20210405180319-a5a99cb37ef4
	golang.org/x/text v0.3.0
	gopkg.in/yaml.v3 v3.0.0
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
)
