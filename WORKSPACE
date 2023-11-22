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

workspace(name = "google_retail_gen_ai")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

###############################################################################
# Sky lib Tool Chain
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
    sha256 = "d55addc2b7226397e777aaa4edf6b5d8488ed332f17732ed0490e060dae2ef9e",
    strip_prefix = "protoc-gen-bq-schema-fork-1.1.0",
    urls = [
        "https://github.com/rrmcguinness/protoc-gen-bq-schema-fork/archive/refs/tags/v1.1.0.tar.gz",
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
    nodejs = True,
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

load("@rules_proto_grpc//python:repositories.bzl", rules_proto_grpc_python_repos = "python_repos")

rules_proto_grpc_python_repos()

load("@rules_proto_grpc//js:repositories.bzl", rules_proto_grpc_js_repos = "js_repos")

rules_proto_grpc_js_repos()

load("@build_bazel_rules_nodejs//:repositories.bzl", "build_bazel_rules_nodejs_dependencies")

build_bazel_rules_nodejs_dependencies()

load("@com_github_grpc_grpc//bazel:grpc_deps.bzl", "grpc_deps")

grpc_deps()

#load("@com_github_grpc_grpc//bazel:grpc_extra_deps.bzl", "grpc_extra_deps")
#
#grpc_extra_deps()

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

###############################################################################
# Gazelle Tool Chain (Do not move)
###############################################################################

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

###############################################################################
# Java Tool Chain
###############################################################################

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

RULES_JVM_EXTERNAL_TAG = "5.3"

RULES_JVM_EXTERNAL_SHA = "d31e369b854322ca5098ea12c69d7175ded971435e55c18dd9dd5f29cc5249ac"

http_archive(
    name = "rules_jvm_external",
    sha256 = RULES_JVM_EXTERNAL_SHA,
    strip_prefix = "rules_jvm_external-%s" % RULES_JVM_EXTERNAL_TAG,
    url = "https://github.com/bazelbuild/rules_jvm_external/releases/download/%s/rules_jvm_external-%s.tar.gz" % (RULES_JVM_EXTERNAL_TAG, RULES_JVM_EXTERNAL_TAG),
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
load("//:build/java_deps.bzl", "JAVA_DEPS", "TEST_DEPS")

maven_install(
    artifacts = JAVA_DEPS + IO_GRPC_GRPC_JAVA_ARTIFACTS + TEST_DEPS,
    fetch_sources = True,
    generate_compat_repositories = True,
    override_targets = IO_GRPC_GRPC_JAVA_OVERRIDE_TARGETS,
    repositories = [
        "https://maven.google.com",
        "https://repo1.maven.org/maven2",
        "https://repository.jboss.org/nexus/content/groups/public/",
    ],
)

load("@maven//:compat.bzl", "compat_repositories")

compat_repositories()

grpc_java_repositories()

###############################################################################
# JUnit 5 Tool Chain
###############################################################################
http_archive(
    name = "contrib_rules_jvm",
    sha256 = "4d62589dc6a55e74bbe33930b826d593367fc777449a410604b2ad7c6c625ef7",
    strip_prefix = "rules_jvm-0.19.0",
    url = "https://github.com/bazel-contrib/rules_jvm/releases/download/v0.19.0/rules_jvm-v0.19.0.tar.gz",
)

load("@contrib_rules_jvm//:repositories.bzl", "contrib_rules_jvm_deps")

contrib_rules_jvm_deps()

# Now ensure that the downloaded deps are properly configured
load("@contrib_rules_jvm//:setup.bzl", "contrib_rules_jvm_setup")

contrib_rules_jvm_setup()

###############################################################################
# Python Tool Chain
###############################################################################

http_archive(
    name = "rules_python",
    sha256 = "9acc0944c94adb23fba1c9988b48768b1bacc6583b52a2586895c5b7491e2e31",
    strip_prefix = "rules_python-0.27.0",
    url = "https://github.com/bazelbuild/rules_python/releases/download/0.27.0/rules_python-0.27.0.tar.gz",
)

load("@rules_python//python:repositories.bzl", "py_repositories")

py_repositories()

load("@rules_python//python:repositories.bzl", "python_register_toolchains")

python_register_toolchains(
    name = "python3",
    python_version = "3.10",
)

load("@python3//:defs.bzl", "interpreter")
load("@rules_python//python:pip.bzl", "pip_parse")

pip_parse(
    name = "rules_proto_grpc_py3_deps",
    python_interpreter_target = interpreter,
    requirements_lock = "//:build/requirements.txt",
)

load("@rules_proto_grpc_py3_deps//:requirements.bzl", local_deps = "install_deps")

local_deps()

###############################################################################
# Rules Package for Tar
###############################################################################
http_archive(
    name = "rules_pkg",
    sha256 = "8f9ee2dc10c1ae514ee599a8b42ed99fa262b757058f65ad3c384289ff70c4b8",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_pkg/releases/download/0.9.1/rules_pkg-0.9.1.tar.gz",
        "https://github.com/bazelbuild/rules_pkg/releases/download/0.9.1/rules_pkg-0.9.1.tar.gz",
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
    url = "https://github.com/rrmcguinness/rules_hugo/archive/%s.zip" % RULES_HUGO_COMMIT,
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
# NodeJS Toolchain
###############################################################################
load("@build_bazel_rules_nodejs//:index.bzl", "yarn_install")

yarn_install(
    name = "npm",
    package_json = "@rules_proto_grpc//js:requirements/package.json",  # This should be changed to your local package.json which should contain the dependencies required
    yarn_lock = "@rules_proto_grpc//js:requirements/yarn.lock",
)

http_archive(
    name = "rules_nodejs",
    sha256 = "162f4adfd719ba42b8a6f16030a20f434dc110c65dc608660ef7b3411c9873f9",
    strip_prefix = "rules_nodejs-6.0.2",
    url = "https://github.com/bazelbuild/rules_nodejs/releases/download/v6.0.2/rules_nodejs-v6.0.2.tar.gz",
)

###############################################################################
# JavaScript Toolchain
###############################################################################

http_archive(
    name = "aspect_rules_js",
    sha256 = "a949d56fed8fa0a8dd82a0a660acc949253a05b2b0c52a07e4034e27f11218f6",
    strip_prefix = "rules_js-1.33.1",
    url = "https://github.com/aspect-build/rules_js/releases/download/v1.33.1/rules_js-v1.33.1.tar.gz",
)

load("@aspect_rules_js//js:repositories.bzl", "rules_js_dependencies")

rules_js_dependencies()

###############################################################################
# Node TypeScript Toolchain
###############################################################################

http_archive(
    name = "aspect_rules_ts",
    sha256 = "8aabb2055629a7becae2e77ae828950d3581d7fc3602fe0276e6e039b65092cb",
    strip_prefix = "rules_ts-2.0.0",
    url = "https://github.com/aspect-build/rules_ts/releases/download/v2.0.0/rules_ts-v2.0.0.tar.gz",
)

load("@aspect_rules_ts//ts:repositories.bzl", "rules_ts_dependencies")

rules_ts_dependencies(
    ts_version_from = "//sdk/js:package.json",
)

load("@aspect_bazel_lib//lib:repositories.bzl", "register_copy_directory_toolchains", "register_copy_to_directory_toolchains")

register_copy_directory_toolchains()

register_copy_to_directory_toolchains()

###############################################################################
# Jest Toolchain
###############################################################################

http_archive(
    name = "aspect_rules_jest",
    sha256 = "098186ffc450f2a604843d8ba14217088a0e259ea6a03294af5360a7f1bcd3e8",
    strip_prefix = "rules_jest-0.19.5",
    url = "https://github.com/aspect-build/rules_jest/releases/download/v0.19.5/rules_jest-v0.19.5.tar.gz",
)

load("@aspect_rules_jest//jest:dependencies.bzl", "rules_jest_dependencies")

rules_jest_dependencies()

###############################################################################
# JS Dependencies Toolchain
###############################################################################

load(
    "@rules_nodejs//nodejs:repositories.bzl",
    "DEFAULT_NODE_VERSION",
    "nodejs_register_toolchains",
)

nodejs_register_toolchains(
    name = "nodejs",
    node_version = DEFAULT_NODE_VERSION,
)

load("@aspect_rules_js//npm:repositories.bzl", "npm_translate_lock")

npm_translate_lock(
    name = "retail_data_model_npm",
    data = [
        "//sdk/js:package.json",
    ],
    npm_package_lock = "//sdk/js:package-lock.json",
    npmrc = "//:sdk/js/.npmrc",
    pnpm_lock = "//sdk/js:pnpm-lock.yaml",
    verify_node_modules_ignored = "//:sdk/js/.bazelignore",
)

load("@retail_data_model_npm//:repositories.bzl", "npm_repositories")

npm_repositories()

###############################################################################
# Multirun - Allows running multiple targets in parallel
###############################################################################

git_repository(
    name = "com_github_atlassian_bazel_tools",
    commit = "8b69172a66e62060a628e13111ca8d9072c4978e",
    remote = "https://github.com/atlassian/bazel-tools.git",
)

load("@com_github_atlassian_bazel_tools//multirun:deps.bzl", "multirun_dependencies")

multirun_dependencies()
