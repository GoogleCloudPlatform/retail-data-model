# Retail API Python Language Build and Example

## Build

The 'BUILD.bazel' file in this directory is responsible for building the
Python Language Library that is importable for other projects.

## Test Cases

The test directory contains the following examples:

* [Using Enumerations](test/enum_test.py)
* [Using Model Objects](test/model_test.py)
* [Creating a Service](test/mock_country_server.py)
* [Creating a Client](test/country_client.py)
* [Testing the Client and Server](test/grpc_test.py)

## Build targets

### Build Go

```shell
$ bazel build //python/...
```

### Testing Go

```shell
$ bazel test //python/...
```