#!/bin/bash

export GOPACKAGESDRIVER_RULES_GO_REPOSITORY_NAME=
exec bazel run --tool_tag=gopackagesdriver -- @io_bazel_rules_go//go/tools/gopackagesdriver "${@}"
