# Standard Go Project Layout

## Overview

This is a basic helloworld golang web program , just for personal use.
It would provide a GRPC,JSON,Swagger Web Interface.

## Quick Start

### Requirement

* Golang 1.7

### Compile and Run

#### Build
```
git clone https://github.com/yank1/golang-hellowolrd.git
cd golang-hellowolrd
make
```

#### Run

```
./bin/golang-hellowolrd
```

Access the Helloword by :
* curl http://localhost:8080/v1/ping
* Open Browser at http://localhost:8080/v1/ping
* echo grpccurl ****
* echo View API document by ***

#### Run for develop
```
go run cmd/main.go
```

### UnitTest

```
make test
```

### Release

```
make release # build image and release
```

### Run in Docker/Kubernetes

```
docker run *******
kubectl run ********
```
