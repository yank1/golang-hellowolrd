# Standard Go Project Layout

## Overview

This is a basic helloworld golang web program , just for personal use.
It would provide a GRPC,JSON,Swagger Web Interface.

## Quick Start

### Requirement

* Golang 1.7

### Compile and Run

```
git clone https://github.com/yank1/golang-hellowolrd.git
cd golang-hellowolrd
make
RUN COMMAND
echo Access the Helloword by :
echo curl ****
echo grpccurl ****
echo View API document by ***
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
