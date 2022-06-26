[![Go CI](https://github.com/sir-gon/projecteuler-go/actions/workflows/go.yml/badge.svg)](https://github.com/sir-gon/projecteuler-go/actions/workflows/go.yml) [![codecov](https://codecov.io/gh/sir-gon/projecteuler-go/branch/main/graph/badge.svg?token=U3N3HSC3YC)](https://codecov.io/gh/sir-gon/projecteuler-go)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fsir-gon%2Fprojecteuler-go.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fsir-gon%2Fprojecteuler-go?ref=badge_shield)

![GitHub](https://img.shields.io/github/license/sir-gon/projecteuler-go) ![GitHub language count](https://img.shields.io/github/languages/count/sir-gon/projecteuler-go) ![GitHub top language](https://img.shields.io/github/languages/top/sir-gon/projecteuler-go)



# What is this?

[Project Euler](https://projecteuler.net/) provide some algorithms and mathematical problems to solve to be used as experience tests.

Use this answers to learn some tip and tricks for algorithms tests.

# Using Go native runtime

## Requirements

You must install dependencies:

```
go mod download
```

Or using make 

```
make dependencies
```

## Testing silently

Every problem is a function with unit test.
Unit test has test cases and input data to solve the problem.

Run all tests:

```
go test -v -coverprofile=coverage/c.out ./...
go tool -func=coverage/c.out
go tool -html=coverage/c.out
```

Or using make:

```
make test coverage
```

# Running with Docker 🐳

## Build a complete image with and run all tests
Running container with testing (final) target.

Designed to store all application files and dependencies as a complete runnable image.
Coverage results will be stored in host **/coverage** directory (mounted as volume).

```
# Build a complete image
docker-compose build projecteuler-go
docker-compose run --rm projecteuler-go make test coverage
```


## Build and run a development image

Running container with development target.
Designed to develop on top of this image. All source application is mounted as a volume in **/app** directory.
Dependencies should be installed to run (not present in this target) so, you must install dependencies before run (or after a dependency add/change).

```
# install dependencies using docker runtime and store them in host directory
docker-compose build projecteuler-go-dev
docker-compose run --rm projecteuler-go-dev make dependencies
```

# About development

Developed with runtime:

```
go version
go version go1.18.3 darwin/amd64
```

# Why I publish solutions?

As Project Euler says:

https://projecteuler.net/about#publish


```
I learned so much solving problem XXX, so is it okay to publish my solution elsewhere?
It appears that you have answered your own question. There is nothing quite like that "Aha!" moment when you finally beat a problem which you have been working on for some time. It is often through the best of intentions in wishing to share our insights so that others can enjoy that moment too. Sadly, that will rarely be the case for your readers. Real learning is an active process and seeing how it is done is a long way from experiencing that epiphany of discovery. Please do not deny others what you have so richly valued yourself.

However, the rule about sharing solutions outside of Project Euler does not apply to the first one-hundred problems, as long as any discussion clearly aims to instruct methods, not just provide answers, and does not directly threaten to undermine the enjoyment of solving later problems. Problems 1 to 100 provide a wealth of helpful introductory teaching material and if you are able to respect our requirements, then we give permission for those problems and their solutions to be discussed elsewhere.
```


If you have better answers or optimal solutions, fork and PR-me

Enjoy 😁 !


## License
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fsir-gon%2Fprojecteuler-go.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fsir-gon%2Fprojecteuler-go?ref=badge_large)