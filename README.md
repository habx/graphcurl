# Graphcurl

[![codecov](https://codecov.io/gh/habx/graphcurl/branch/dev/graph/badge.svg?token=376464GH1H)](https://codecov.io/gh/habx/graphcurl)
[![Release](https://img.shields.io/github/v/release/habx/graphcurl)](https://github.com/habx/graphcurl/releases/latest)
[![Go version](https://img.shields.io/github/go-mod/go-version/habx/graphcurl/dev)](https://golang.org/doc/devel/release.html)
[![Docker](https://img.shields.io/docker/pulls/habx/graphcurl)](https://hub.docker.com/r/habx/graphcurl)
[![CircleCI](https://img.shields.io/circleci/build/github/habx/graphcurl/dev)](https://app.circleci.com/pipelines/github/habx/graphcurl)
[![License](https://img.shields.io/github/license/habx/graphcurl)](/LICENSE)


I created this tool to facilitate GraphQL usage in CLI.

There are many use cases in devops use : [Argo](https://github.com/argoproj/argo-workflows), crontab, jenkins..

I was inspired by another open source project. [hasura/graphqurl](https://github.com/hasura/graphqurl)

## Getting Started

### CLI usage

[Command line](docs)

### Installing


#### Docker

```shell script
docker pull habx/graphcurl
docker container run --rm habx/graphcurl --help
```

#### Binary

##### MACOS

Set VERSION

```shell script
VERSION=vx.x.x wget https://github.com/habx/graphcurl/releases/download/${VERSION}/graphcurl_darwin_amd64.gz
```

##### LINUX

```shell script
VERSION=vx.x.x wget https://github.com/habx/graphcurl/releases/download/${VERSION}/graphcurl_linux_amd64.gz
```

#### go source

```shell script
go get -t github.com/habx/graphcurl
graphcurl --help
```

## Build and tun tests

Git clone
```shell script
git clone git@github.com:habx/graphcurl.git
cd graphcurl
```

Go build
```shell script
go build
```

Go tests
```shell script
go test -v $(go list ./... | grep -v "vendor") -covermode=atomic -coverprofile=coverage.txt
```

## Built With

* [machinebox/graphql](https://github.com/machinebox/graphql) - Low-level GraphQL client for Go.
* [spf13/cobra](https://github.com/spf13/cobra) - Cobra is both a library for creating powerful modern CLI applications

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/habx/graphcurl/tags). 

## Authors

* [Clement](http://github.com/clementlecorre)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
