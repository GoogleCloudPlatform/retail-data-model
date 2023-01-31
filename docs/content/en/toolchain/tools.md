---
title: "Tools"
weight: 5
date: 2022-08-25T23:12:18-04:00
---

## Overview

This product uses [Bazel](https://bazel.build/) as the build system. Bazel is
supported on all major platforms [Linux, Windows, and OS X]. The build chain
is hermetic, meaning that only bazel is required to build the artifacts,
which are built for the following targets:

* Go Lang - Version 1.19+
* Java - Version 11+
* Python - 3.9+

## Build

The CI/CD process of this project creates multiple artifacts: documentation,
a Go language module, a set of Java Libraries, and a set of Python Libraries.

Each Library contains the domain model representation (pb files), and the
service API definitions (grpc files).

These products are then used by core implementation teams to build service
offering upon.

### Building

The following commands may be used for building and verifying artifacts. All
artifacts are published to the sym-link directory 'bazel-bin'.

```shell
# Build All
$ bazel build //...

# Test All
$ bazel test //... 
```

## Tool Chain

- [Protocol Buffers](https://developers.google.com/protocol-buffers) - A language-neutral, platform-neutral extensible mechanism for serializing data.
- [gRPC](https://grpc.io/) - A high-performance, open source, universal RPC framework.
- [GitHub](https://github.com/) - Source Code Repository.
- [Bazel](https://bazel.build) - Build and Dependency System.
- [Hugo](https://gohugo.io/) - A markdown driven documentation static website generator.
  - [Hugo Relearn Theme](https://github.com/McShelby/hugo-theme-relearn) - An update to the learn theme used for organizing documentation.
- [Google gCloud CLI](https://cloud.google.com/sdk/gcloud).
- [Go](https://go.dev/) - Go Programming Language.
  - [Testify](https://github.com/stretchr/testify) - Test assertion Framework.
- [Java 11/17](https://openjdk.org/) - The Java Programming Language.
  - [JUnit 5](https://junit.org/junit5/docs/current/user-guide/) - Testing Framework.

## Required Tools

Only two of the tools are required to be installed, the remainder of the tools are self-contained and updated via the Bazel build rules.

- [GitHub Account](https://github.com/) - In order to contribute to the codebase, you'll need a GitHub account.

- [Installing Bazel](https://bazel.build/start) - In accordance to the community recomendation, we too recommend using Bazelisk. This utility will help you keep Bazel and it's dependencies up-to-date transparently.

- [Installing Google gCloud CLI](https://cloud.google.com/sdk/docs/install) - This link has the most up-to-date information for installing the gCloud CLI.

## Integrated Development Environment (IDE)

In order to take advantage of your modern IDEs capabilities for autocomplete, verification, formatting, etc.
we highly recommend using the IntelliJ Ultimate IDE. This will allow you to get the most out of every language used in the project.

We recommend using [IntelliJ Idea](https://www.jetbrains.com/idea/) with this project. It's not required, but it will make your life easier as the default workspace has been created using the code styles and standards.

You may also use [Jet Brains: Go Land](https://www.jetbrains.com/go/) but will be limited to the go and starlark languages.

Please see the [Code Style Guide](code_style_guide.html) if you are using another editor.

## Recommended Tools

### IntelliJ Ultimate / GoLand 

Install the Bazel Plugin

In the .ijwb directory, create (or edit) the .bazelproject file.

```text
directories:
  # Add the directories you want added as source here
  # By default, we've added your entire workspace ('.')
  .
  ./docs
  ./tools
  -conf
  -.vscode
  -.idea
  -.editorconfig
  -.markdownlint.yaml
  -.shellcheckrc
  -.ijwb
  -.github

# Automatically includes all relevant targets under the 'directories' above
derive_targets_from_directories: true

targets:
  # If source code isn't resolving, add additional targets that compile it here

additional_languages:
  go
  # Uncomment any additional languages you want supported
  # android
  # dart
  # go
  # javascript
  # kotlin
  # python
  # scala
  # typescript

java_language_level: 11

build_flags:
  --java_language_version=11

test_flags:
  --java_language_version=11 --test_output=all
```

### Visual Studio Code

You may simply open the root directory. We recommend using the Bazel plugin. 
Due to limitations in this IDE and the Java and Go plugins, there are significant issues
with autocomplete and other features. This includes poor performance from the language server.

### Sublime Text

Untested.