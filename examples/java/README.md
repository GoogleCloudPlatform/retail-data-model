# Retail API Java Language Build and Example

## Introduction

The Java example is laid out in a traditional Maven
directory structure. The two primary classes illustrating
production and consumption are found in `src/main/java`
directory, while all tests are located in `src/test/java`.

```shell
bazel build //examples/java/...
bazel test //examples/java/...
```

## Build

The 'BUILD.bazel' file in this directory is responsible for building the
Java Language Library that is importable for other projects.

## Test Cases

The test directory contains the following examples:

* [Using Enumerations](src/test/com/google/retail/EnumTest.java)
* [Using Model Objects](src/test/com/google/retail/ModelTest.java)
* [Creating a Service](src/test/com/google/retail/CountryService.java)
* [Creating a Client](src/test/com/google/retail/CountriesClient.java)
* [Testing the Client and Server](src/test/com/google/retail/ServiceTest.java)