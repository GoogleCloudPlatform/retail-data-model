---
title: "Using Visual Studio Code"

resources:
- name: banner
  src: vscode-build.png
  title: "Visual Studio Code IDE"
  params:
  credits: "[Ryan McGuinness](https://www.github.com/rrmcguinness)"
---

<div style="text-align: center">
{{< img name="banner" size="small" lazy=false >}}
</div>

## Prerequisites

To effectively develop in multiple language in VS Code, you'll need to add some system extensions
and several plugins. This requires you to understand your terminal environment.

## Language Server Preparation

VS Code uses LSP similar to Sublime Text. LSP is a system level abstraction allowing multiple
languages, syntax highlighting, ect. to function within the IDE. LSP DOES NOT stand alone and
requires a few programs to be on your system path. Please install the following:

1. [Buildifer](https://github.com/bazelbuild/buildtools/blob/master/buildifier/README.md) 
2. [Buildozer](https://github.com/bazelbuild/buildtools/blob/master/buildozer/README.md) 
3. [Unused Deps](https://github.com/bazelbuild/buildtools/blob/master/unused_deps/README.md)

> NOTE: The fastest way to install these tools is to have Go installed and on your path, and
> simply execute the following commands:

```shell
go install github.com/bazelbuild/buildtools/buildifier@latest
go install github.com/bazelbuild/buildtools/buildozer@latest
go install github.com/bazelbuild/buildtools/unused_deps@latest
go install -v golang.org/x/tools/gopls@latest

# Install the go debugger
go install github.com/go-delve/delve/cmd/dlv@latest
# Static code analysis
go install honnef.co/go/tools/cmd/staticcheck@latest

// Symlink your bazelisk version of bazel to /usr/bin
sudo ln -s ${HOME}/go/bin/bazelisk /usr/bin/bazel
```

Ensure your `${HOME}/go/bin` is on your PATH.

## VS Code Plugins

* Bazel
* IntelliCode
* Language Support for Java
* Debugger for Java
* Test Runner for Java
* Go
* Go Test Explorer
* Python

Once installed, re-open your project as a workspace, using the workspace file found in the root
directory.

## Additional Personal Settings

You'll still need to configure your local Java JDK, Go, and Python settings by adding sections
to your configuration.

## Building and testing

If you used the workspace, you'll have preconfigured tasks, simply open the command pallet
`ctrl+shift+p` and Type: 'Tasks: Run Build Task' or 'Tasks: Run Test Task' to test all.

Task List:

* Build
* Test
* Build Documents
* Server Document WebSite

