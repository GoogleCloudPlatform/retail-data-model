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

workspace(name = "google_retail_data_model")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

###############################################################################
# Skylib Tool Chain
###############################################################################
http_archive(
    name = "bazel_skylib",
    sha256 = "66ffd9315665bfaafc96b52278f57c7e2dd09f5ede279ea6d39b2be471e7e3aa",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-skylib/releases/download/1.4.2/bazel-skylib-1.4.2.tar.gz",
        "https://github.com/bazelbuild/bazel-skylib/releases/download/1.4.2/bazel-skylib-1.4.2.tar.gz",
    ],
)

load("@bazel_skylib//:workspace.bzl", "bazel_skylib_workspace")

bazel_skylib_workspace()

###############################################################################
# Proto Gen Schema
###############################################################################
http_archive(
    name = "protoc_gen_bq_schema",
    sha256 = "98f700a6856686be10fb844ed8792cacb1db94854b4931e795e3657d46fd8c99",
    strip_prefix = "protoc-gen-bq-schema-fork-1.0.0",
    urls = [
        "https://github.com/rrmcguinness/protoc-gen-bq-schema-fork/archive/refs/tags/v1.0.0.tar.gz",
    ],
)

###############################################################################
# Google APIs
###############################################################################
http_archive(
    name = "com_google_googleapis",
    sha256 = "9d1a930e767c93c825398b8f8692eca3fe353b9aaadedfbcf1fca2282c85df88",
    strip_prefix = "googleapis-64926d52febbf298cb82a8f472ade4a3969ba922",
    urls = [
        "https://github.com/googleapis/googleapis/archive/64926d52febbf298cb82a8f472ade4a3969ba922.zip",
    ],
)

load("@com_google_googleapis//:repository_rules.bzl", "switched_rules_by_language")

switched_rules_by_language(
    name = "com_google_googleapis_imports",
    cc = True,
    go = True,
    grpc = True,
    java = True,
    python = True,
)

###############################################################################
# GRPC Tool Chain
###############################################################################

http_archive(
    name = "rules_proto_grpc",
    sha256 = "9ba7299c5eb6ec45b6b9a0ceb9916d0ab96789ac8218269322f0124c0c0d24e2",
    strip_prefix = "rules_proto_grpc-4.5.0",
    urls = ["https://github.com/rules-proto-grpc/rules_proto_grpc/releases/download/4.5.0/rules_proto_grpc-4.5.0.tar.gz"],
)

load("@rules_proto_grpc//:repositories.bzl", "rules_proto_grpc_repos", "rules_proto_grpc_toolchains")

rules_proto_grpc_toolchains()

rules_proto_grpc_repos()

load("@rules_proto//proto:repositories.bzl", "rules_proto_dependencies", "rules_proto_toolchains")

rules_proto_dependencies()

rules_proto_toolchains()

load("@rules_proto_grpc//doc:repositories.bzl", rules_proto_grpc_doc_repos = "doc_repos")

rules_proto_grpc_doc_repos()

load("@com_github_grpc_grpc//bazel:grpc_deps.bzl", "grpc_deps")

grpc_deps()

###############################################################################
# GO Tool Chain
###############################################################################

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "51dc53293afe317d2696d4d6433a4c33feedb7748a9e352072e2ec3c0dafd2c6",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.40.1/rules_go-v0.40.1.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.40.1/rules_go-v0.40.1.zip",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains(version = "1.20.7")

http_archive(
    name = "bazel_gazelle",
    sha256 = "727f3e4edd96ea20c29e8c2ca9e8d2af724d8c7778e7923a854b2c80952bc405",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.30.0/bazel-gazelle-v0.30.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.30.0/bazel-gazelle-v0.30.0.tar.gz",
    ],
)

# Add Go Lang Rules
load("@rules_proto_grpc//:repositories.bzl", "io_bazel_rules_go")  # buildifier: disable=same-origin-load

io_bazel_rules_go()

load("@rules_proto_grpc//go:repositories.bzl", rules_proto_grpc_go_repos = "go_repos")

rules_proto_grpc_go_repos()

load("//:build/go_deps.bzl", "go_dependencies")

# gazelle:repository_macro build/go_deps.bzl%go_dependencies
go_dependencies()

# load("@com_github_grpc_grpc//bazel:grpc_extra_deps.bzl", "grpc_extra_deps")
# grpc_extra_deps()

###############################################################################
# Java Tool Chain
###############################################################################

RULES_JVM_EXTERNAL_TAG = "4.5"

RULES_JVM_EXTERNAL_SHA = "b17d7388feb9bfa7f2fa09031b32707df529f26c91ab9e5d909eb1676badd9a6"

http_archive(
    name = "rules_jvm_external",
    sha256 = RULES_JVM_EXTERNAL_SHA,
    strip_prefix = "rules_jvm_external-%s" % RULES_JVM_EXTERNAL_TAG,
    url = "https://github.com/bazelbuild/rules_jvm_external/archive/%s.zip" % RULES_JVM_EXTERNAL_TAG,
)

load("@rules_jvm_external//:repositories.bzl", "rules_jvm_external_deps")

rules_jvm_external_deps()

load("@rules_jvm_external//:setup.bzl", "rules_jvm_external_setup")

rules_jvm_external_setup()

# GRPC
load("@rules_proto_grpc//java:repositories.bzl", rules_proto_grpc_java_repos = "java_repos")

rules_proto_grpc_java_repos()

load("@io_grpc_grpc_java//:repositories.bzl", "IO_GRPC_GRPC_JAVA_ARTIFACTS", "IO_GRPC_GRPC_JAVA_OVERRIDE_TARGETS", "grpc_java_repositories")
load("@rules_jvm_external//:defs.bzl", "maven_install")

maven_install(
    artifacts = [
        "org.apache.commons:commons-lang3:3.13.0",
        "com.google.protobuf:protobuf-java-util:3.24.2",
        "io.grpc:grpc-core:1.57.2",
        "io.grpc:grpc-googleapis:1.57.2",
        "io.grpc:grpc-netty-shaded:1.57.2",
        "io.grpc:grpc-protobuf:1.57.2",
        "io.grpc:grpc-stub:1.57.2",
        "io.grpc:grpc-testing:1.57.2",
        "org.apache.tomcat:annotations-api:6.0.53",
        "org.apache.logging.log4j:log4j-api:2.18.0",
        "org.apache.logging.log4j:log4j-core:2.18.0",
        "com.google.protobuf:protoc:3.24.2",
        "io.netty:netty-all:4.1.79.Final",
    ] + IO_GRPC_GRPC_JAVA_ARTIFACTS,
    generate_compat_repositories = True,
    override_targets = IO_GRPC_GRPC_JAVA_OVERRIDE_TARGETS,
    repositories = [
        "https://maven.google.com",
        "https://repo1.maven.org/maven2",
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
    sha256 = "5868e73107a8e85d8f323806e60cad7283f34b32163ea6ff1020cf27abef6036",
    strip_prefix = "rules_python-0.25.0",
    url = "https://github.com/bazelbuild/rules_python/releases/download/0.25.0/rules_python-0.25.0.tar.gz",
)

load("@rules_python//python:repositories.bzl", "python_register_toolchains")

python_register_toolchains(
    name = "python311",
    python_version = "3.11.4",
    tool_versions = {
        "3.11.4": {
            "url": "20230726/cpython-{python_version}+20230726-{platform}-{build}.tar.gz",
            "sha256": {
                "aarch64-apple-darwin": "cb6d2948384a857321f2aa40fa67744cd9676a330f08b6dad7070bda0b6120a4",
                "aarch64-unknown-linux-gnu": "2e84fc53f4e90e11963281c5c871f593abcb24fc796a50337fa516be99af02fb",
                "ppc64le-unknown-linux-gnu": "df7b92ed9cec96b3bb658fb586be947722ecd8e420fb23cee13d2e90abcfcf25",
                "s390x-unknown-linux-gnu": "e477f0749161f9aa7887964f089d9460a539f6b4a8fdab5166f898210e1a87a4",
                "x86_64-apple-darwin": "47e1557d93a42585972772e82661047ca5f608293158acb2778dccf120eabb00",
                "x86_64-pc-windows-msvc": "878614c03ea38538ae2f758e36c85d2c0eb1eaaca86cd400ff8c76693ee0b3e1",
                "x86_64-unknown-linux-gnu": "e26247302bc8e9083a43ce9e8dd94905b40d464745b1603041f7bc9a93c65d05",
            },
            "strip_prefix": "python",
        },
    },
)

load("@python311//:defs.bzl", "interpreter")
load("@rules_proto_grpc//python:repositories.bzl", rules_proto_grpc_python_repos = "python_repos")

rules_proto_grpc_python_repos()

load("@rules_python//python:pip.bzl", "pip_parse")

pip_parse(
    name = "python_deps",
    python_interpreter_target = interpreter,
    requirements_lock = "//:build/requirements.txt",
)

# pip_parse(
#     name = "rules_proto_grpc_py3_deps",
#     python_interpreter = "python3",
#     requirements_lock = "@rules_proto_grpc//python:requirements.txt",
# )

# load("@rules_proto_grpc_py3_deps//:requirements.bzl", "install_deps")
# install_deps()

load("@python_deps//:requirements.bzl", local_deps = "install_deps")

local_deps()

###############################################################################
# Rules Package for Tar
###############################################################################
http_archive(
    name = "rules_pkg",
    sha256 = "eea0f59c28a9241156a47d7a8e32db9122f3d50b505fae0f33de6ce4d9b61834",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_pkg/releases/download/0.8.0/rules_pkg-0.8.0.tar.gz",
        "https://github.com/bazelbuild/rules_pkg/releases/download/0.8.0/rules_pkg-0.8.0.tar.gz",
    ],
)

load("@rules_pkg//:deps.bzl", "rules_pkg_dependencies")

rules_pkg_dependencies()

###############################################################################
# Hugo Tool Chain
###############################################################################

RULES_HUGO_COMMIT = "02234789fa9f2112807c1642eacb9f9728fc179d"

RULES_HUGO_SHA256 = "4ce20c981ad50ac0c956e85ef991e59b204778bde59d81e40be05450259ae969"

RULES_HUGO_VERSION = "0.101.0"

HUGO_THEME_SHA256 = "7fdd57f7d4450325a778629021c0fff5531dc8475de6c4ec70ab07e9484d400e"

HUGO_THEME_URL = "https://github.com/thegeeklab/hugo-geekdoc/releases/download/v0.34.2/hugo-geekdoc.tar.gz"

http_archive(
    name = "build_stack_rules_hugo",
    sha256 = RULES_HUGO_SHA256,
    strip_prefix = "rules_hugo-%s" % RULES_HUGO_COMMIT,
    url = "https://github.com/GoogleCloudPlatform/rules_hugo/archive/%s.zip" % RULES_HUGO_COMMIT,
)

load("@build_stack_rules_hugo//hugo:rules.bzl", "hugo_repository")

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

###############################################################################
# Gazelle Tool Chain (Do not move)
###############################################################################

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()
