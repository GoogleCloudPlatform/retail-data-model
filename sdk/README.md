# Software Development Kit (SDK)

## Introduction

The SDK folder is where language specific SDKs are defined.
There is intentionally no language specific code in these
directories, only definitions used to generate the language
specific artifacts.

## [Go](go/README.md)

The Go directory contains the build and module files used
by the Go language for publishing to the `github/rrmcguinness/retail-data-model@go`
branch. This allows a developer to access the generated
artifacts using `go get`.

## [Java](java/)

The Java directory contains the build files for the Java
language and is used for publishing artifacts to
Maven Central.

## [Python](python/)

The Python directory contains the build files for Python 3.9
and above and is used to publish artifacts to the
`github.com/rrmcguinness/retail-data-model@python` branch.