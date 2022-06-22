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

replace golang.org/x/crypto v0.0.0-20191011191535-87dc89f01550 => golang.org/x/crypto v0.0.0-20220525230936-793ad666bf5e // indirect

exclude golang.org/x/crypto v0.0.0-20191011191535-87dc89f01550
