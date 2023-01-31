---
title: "Workstation Setup"
---

## Prerequisites

Each or these setup instruction sets requires you to be familiar with your operating system
and the terminal environments on those operating systems. In addition, they may require system
access for you to set up Development tools and change filesystem permissions. This IS NOT 
required to run the software, only to load the supporting tools.

## Linux 

### Ubuntu (apt)

On Ubuntu, you'll need to have sudo access to run the apt package manager.

```shell
$ sudo apt install build-essential
$ sudo apt install python3
$ sudo apt install bazelisk
```

## Mac OS X

On Mac OS, you'll need to have sudo access to install the development tools.

```shell

# Install XCode (This install Git, the C language and the GCC compiler.
xcode-select --install

# Install Brew - A package manager for OS X
$ /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"

# Python GRCP REQUIRES python headers, Install cython, all other tools are self-contained.

# You should change the ownership of these directories to your user.
$ chown -R $(whoami) /usr/local/bin /usr/local/etc /usr/local/sbin /usr/local/share /usr/local/share/doc

# And make sure that your user has write permission.
$ chmod u+w /usr/local/bin /usr/local/etc /usr/local/sbin /usr/local/share /usr/local/share/doc

# Install Python 3.9
$ brew install cython

# Install Bazel
$ brew install bazelisk
```

## Windows