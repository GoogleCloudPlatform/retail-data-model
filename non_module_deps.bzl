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

load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

def _non_module_dependencies_impl(module_ctx):
    """Implementation of non-module dependencies extension."""

    # Proto Gen BQ Schema
    http_archive(
        name = "protoc_gen_bq_schema",
        sha256 = "d55addc2b7226397e777aaa4edf6b5d8488ed332f17732ed0490e060dae2ef9e",
        strip_prefix = "protoc-gen-bq-schema-fork-1.1.0",
        urls = [
            "https://github.com/rrmcguinness/protoc-gen-bq-schema-fork/archive/refs/tags/v1.1.0.tar.gz",
        ],
    )

    # Hugo Tool Chain
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

    # Hugo repository setup
    # Note: This needs to be loaded after build_stack_rules_hugo is available
    # The actual hugo_repository call will be handled in a separate extension or BUILD file

    # Multirun - Allows running multiple targets in parallel
    git_repository(
        name = "com_github_atlassian_bazel_tools",
        commit = "8b69172a66e62060a628e13111ca8d9072c4978e",
        remote = "https://github.com/atlassian/bazel-tools.git",
    )

non_module_dependencies = module_extension(
    implementation = _non_module_dependencies_impl,
)
