# StatNLP Ladrawex (LaTeX Draw) Service 
[![Build Status](https://travis-ci.org/sutd-statnlp/service-ladrawex.svg?branch=master)](https://travis-ci.org/sutd-statnlp/service-ladrawex)
[![codecov](https://codecov.io/gh/sutd-statnlp/service-ladrawex/branch/master/graph/badge.svg)](https://codecov.io/gh/sutd-statnlp/service-ladrawex)
[![Go Report Card](https://goreportcard.com/badge/github.com/sutd-statnlp/service-ladrawex)](https://goreportcard.com/report/github.com/sutd-statnlp/service-ladrawex)
[![GoDoc](https://godoc.org/github.com/sutd-statnlp/service-ladrawex?status.svg)](https://godoc.org/github.com/sutd-statnlp/service-ladrawex)
[![](https://images.microbadger.com/badges/image/statnlp/service-ladrawex.svg)](https://microbadger.com/images/statnlp/service-ladrawex)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/sutd-statnlp/service-ladrawex/blob/master/LICENSE)

## Installation

Get the service-ladrawex repository

```
go get github.com/sutd-statnlp/service-ladrawex

cd $GOPATH/src/github.com/sutd-statnlp/service-ladrawex
```

And install dependencies

```
go get -u github.com/golang/dep/cmd/dep

dep ensure
```

## Running the tests

Run all tests

```
go test ./...
```

Or run all tests with coverage

```
bash script/coverage.sh
```

## Build and Run

Run main.go
``` bash
go run main.go
# serve at localhost:9000
```

Build and run native binary

``` bash
bash script/Build.sh

./service-ladrawex.out
```
Build native binary for multiple platforms (darwin, windows and linux)

```
bash script/BuildMulti.sh
```

## Environment variables

```bash
    # enable production mode, default is false
    env GBP_PROMODE=true

    # set host address, default is :9000
    env GBP_HOST=":8000"
```
## Docker support 

Build docker image

```
bash script/Dockerbuild.sh
```

Run docker container

```
docker run -d --name service-ladrawex -p 9000:9000 statnlp/service-ladrawex
```
## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/sutd-statnlp/service-ladrawex/tags). 

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details

