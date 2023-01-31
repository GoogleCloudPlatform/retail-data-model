# Retail API Go Language Build and Example

## Introduction

The Go module is setup similar to a production Go package,
in order to illustrate best practices for the Go language.

* pkg - is for internal source code.
* test - is for all test cases, and test harness data.

## Build

The 'BUILD.bazel' file in this directory is responsible for building the 
Go Language Library that is importable for other projects.

## Test Cases

The test directory contains the following examples:

* [Using Enumerations](test/enum_test.go)
* [Using Model Objects](test/model_test.go)
* [Creating a Service](test/mock_country_server.go)
* [Creating a Client](test/country_client.go)
* [Testing the Client and Server](test/service_test.go)
    * [Test GRPC Server](test/server.go)
    
## Build targets

### Build Go

```shell
$ bazel build //go/...
```

### Testing Go

```shell
$ bazel test //go/...
```