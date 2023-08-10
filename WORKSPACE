# Copyright 2022 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

workspace(name = "google_retail_cloud_api")

###############################################################################
# Begin: Build Variables
#
# These variable represent the major flags that MAY be changed to support
# different environments.
###############################################################################

# Go Language Variables
RULES_GO_VERSION = "1.21.0"

# Java Variables
RULES_JVM_EXTERNAL_TAG = "5.3"
RULES_JVM_EXTERNAL_SHA ="d31e369b854322ca5098ea12c69d7175ded971435e55c18dd9dd5f29cc5249ac"

# The maven dependencies of the project, this DOES NOT include the JUNIT 5
# dependencies, please see the //build/junit.bzl file.
PROJECT_MAVEN_DEPENDENCIES = [
    "org.apache.commons:commons-lang3:3.13.0",
    "com.google.protobuf:protobuf-java-util:3.24.0",
    "io.grpc:grpc-core:1.57.1",
    "io.grpc:grpc-googleapis:1.57.1",
    "io.grpc:grpc-netty-shaded:1.57.1",
    "io.grpc:grpc-protobuf:1.57.1",
    "io.grpc:grpc-stub:1.57.1",
    "io.grpc:grpc-testing:1.57.1",
    "org.apache.tomcat:annotations-api:6.0.53",
    "org.apache.logging.log4j:log4j-api:2.18.0",
    "org.apache.logging.log4j:log4j-core:2.18.0",
    "com.google.protobuf:protoc:3.21.5",
    "com.google.code.gson:gson:2.9.0",
    "io.netty:netty-all:4.1.79.Final",
]

# Python Rules
RULES_PYTHON_VERSION = "3.9"

RULES_PYTHON_REVISION = "740825b7f74930c62f44af95c9a4c1bd428d2c53"

# Hugo Documentation Variables
RULES_HUGO_COMMIT = "02234789fa9f2112807c1642eacb9f9728fc179d"

RULES_HUGO_SHA256 = "4ce20c981ad50ac0c956e85ef991e59b204778bde59d81e40be05450259ae969"

RULES_HUGO_VERSION = "0.101.0"

HUGO_THEME_SHA256 = "7fdd57f7d4450325a778629021c0fff5531dc8475de6c4ec70ab07e9484d400e"

HUGO_THEME_URL = "https://github.com/thegeeklab/hugo-geekdoc/releases/download/v0.34.2/hugo-geekdoc.tar.gz"

###############################################################################
# End: Build Variables
###############################################################################

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

###############################################################################
# GRPC Tool Chain
###############################################################################

git_repository(
    name = "com_google_googleapis",
    commit = "59b73bd6f7c00dc5af895414c444b08055849bdf",
    remote = "https://github.com/googleapis/googleapis",
    shallow_since = "1665087085 -0700",
)

load("@com_google_googleapis//:repository_rules.bzl", "switched_rules_by_language")

switched_rules_by_language(
    name = "com_google_googleapis_imports",
    go = True,
    grpc = True,
    java = True,
    python = True,
)

http_archive(
    name = "rules_proto_grpc",
    sha256 = "b244cbede34638ad0e1aec0769f62b8348c7ed71f431a027e252a07d6adba3d6",
    strip_prefix = "rules_proto_grpc-4.4.0",
    urls = ["https://github.com/rules-proto-grpc/rules_proto_grpc/archive/4.4.0.tar.gz"],
)

load("@rules_proto_grpc//:repositories.bzl", "rules_proto_grpc_repos", "rules_proto_grpc_toolchains")

rules_proto_grpc_toolchains()

rules_proto_grpc_repos()

load("@rules_proto//proto:repositories.bzl", "rules_proto_dependencies", "rules_proto_toolchains")

rules_proto_dependencies()

rules_proto_toolchains()

load("@rules_proto_grpc//doc:repositories.bzl", rules_proto_grpc_doc_repos = "doc_repos")

rules_proto_grpc_doc_repos()

###############################################################################
# GO Tool Chain
###############################################################################

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "278b7ff5a826f3dc10f04feaf0b70d48b68748ccd512d7f98bf442077f043fe3",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.41.0/rules_go-v0.41.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.41.0/rules_go-v0.41.0.zip",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "29218f8e0cebe583643cbf93cae6f971be8a2484cdcfa1e45057658df8d54002",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.32.0/bazel-gazelle-v0.32.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.32.0/bazel-gazelle-v0.32.0.tar.gz",
    ],
)

load("@rules_proto_grpc//:repositories.bzl", "io_bazel_rules_go")  # buildifier: disable=same-origin-load

io_bazel_rules_go()

load("@rules_proto_grpc//go:repositories.bzl", rules_proto_grpc_go_repos = "go_repos")

rules_proto_grpc_go_repos()

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains(version = RULES_GO_VERSION)

###############################################################################
# Gazelle Dependencies
###############################################################################
load("//:build/go_deps.bzl", "go_dependencies")

# gazelle:repository_macro build/go_deps.bzl%go_dependencies
go_dependencies()

###############################################################################
# Java Tool Chain
###############################################################################

# External JDK
http_archive(
    name = "rules_jvm_external",
    strip_prefix = "rules_jvm_external-%s" % RULES_JVM_EXTERNAL_TAG,
    sha256 = RULES_JVM_EXTERNAL_SHA,
    url = "https://github.com/bazelbuild/rules_jvm_external/releases/download/%s/rules_jvm_external-%s.tar.gz" % (RULES_JVM_EXTERNAL_TAG, RULES_JVM_EXTERNAL_TAG)
)

load("@rules_jvm_external//:repositories.bzl", "rules_jvm_external_deps")

rules_jvm_external_deps()

load("@rules_jvm_external//:setup.bzl", "rules_jvm_external_setup")

rules_jvm_external_setup()

# Java GRPC
load("@rules_proto_grpc//java:repositories.bzl", rules_proto_grpc_java_repos = "java_repos")

rules_proto_grpc_java_repos()

load("@rules_jvm_external//:defs.bzl", "maven_install")
load("@io_grpc_grpc_java//:repositories.bzl", "IO_GRPC_GRPC_JAVA_ARTIFACTS", "IO_GRPC_GRPC_JAVA_OVERRIDE_TARGETS", "grpc_java_repositories")

maven_install(
    artifacts = PROJECT_MAVEN_DEPENDENCIES + IO_GRPC_GRPC_JAVA_ARTIFACTS,
    generate_compat_repositories = True,
    override_targets = IO_GRPC_GRPC_JAVA_OVERRIDE_TARGETS,
    repositories = [
        "https://repo.maven.apache.org/maven2/",
    ],
)

load("@maven//:compat.bzl", "compat_repositories")

compat_repositories()

grpc_java_repositories()

###############################################################################
# JUnit 5 Tool Chain
###############################################################################
load(
    "//:build/junit.bzl",
    "junit_jupiter_java_repositories",
    "junit_platform_java_repositories",
)

junit_jupiter_java_repositories()

junit_platform_java_repositories()

###############################################################################
# Python Tool Chain
###############################################################################
http_archive(
    name = "rules_python",
    sha256 = "3474c5815da4cb003ff22811a36a11894927eda1c2e64bf2dac63e914bfdf30f",
    strip_prefix = "rules_python-{}".format(RULES_PYTHON_REVISION),
    url = "https://github.com/bazelbuild/rules_python/archive/{}.zip".format(RULES_PYTHON_REVISION),
)

load("@rules_python//python:repositories.bzl", "python_register_toolchains")

python_register_toolchains(
    name = "python_toolchain",
    # Available versions are listed in @rules_python//python:versions.bzl.
    # We recommend using the same version your team is already standardized on.
    python_version = RULES_PYTHON_VERSION,
)

load("@python_toolchain//:defs.bzl", "interpreter")
load("@rules_python//python:pip.bzl", "pip_parse")

pip_parse(
    name = "python_deps",
    python_interpreter_target = interpreter,
    requirements_lock = "//:build/requirements.txt",
)

load("@python_deps//:requirements.bzl", "install_deps")

install_deps()

load("@rules_proto_grpc//python:repositories.bzl", rules_proto_grpc_python_repos = "python_repos")

rules_proto_grpc_python_repos()

load("@com_github_grpc_grpc//bazel:grpc_deps.bzl", "grpc_deps")

grpc_deps()

###############################################################################
# Rules Package for Tar
###############################################################################
http_archive(
    name = "rules_pkg",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_pkg/releases/download/0.8.0/rules_pkg-0.8.0.tar.gz",
        "https://github.com/bazelbuild/rules_pkg/releases/download/0.8.0/rules_pkg-0.8.0.tar.gz",
    ],
    sha256 = "eea0f59c28a9241156a47d7a8e32db9122f3d50b505fae0f33de6ce4d9b61834",
)
load("@rules_pkg//:deps.bzl", "rules_pkg_dependencies")
rules_pkg_dependencies()

###############################################################################
# Hugo Tool Chain
###############################################################################

# Update these to latest
# Use a modified version of the popular rules_hugo, this modified version
# allows modern Hugo themes to be used.
http_archive(
    name = "build_stack_rules_hugo",
    sha256 = RULES_HUGO_SHA256,
    strip_prefix = "rules_hugo-%s" % RULES_HUGO_COMMIT,
    url = "https://github.com/rrmcguinness/rules_hugo/archive/%s.zip" % RULES_HUGO_COMMIT,
)

load("@build_stack_rules_hugo//hugo:rules.bzl", "hugo_repository")

#
# Load hugo binary itself
#
# Optionally, load a specific version of Hugo, with the 'version' argument
hugo_repository(
    name = "hugo",
    extended = True,
    version = RULES_HUGO_VERSION,
)

# Create a readable archive from a GitHub Hugo Theme that DOES NOT support the theme layout.
http_archive(
    name = "theme_geekdoc",
    build_file_content = """
filegroup(
    name = "files",
    srcs = glob(["**"]),
    visibility = ["//visibility:public"],
)
    """,
    sha256 = HUGO_THEME_SHA256,
    url = HUGO_THEME_URL,
)

# DO NOT MOVE THESE LINES
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

gazelle_dependencies()
