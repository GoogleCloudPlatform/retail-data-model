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

load("@build_stack_rules_hugo//hugo:rules.bzl", "hugo_serve", "hugo_site", "hugo_theme")
load("//:build/copy_file_groups.bzl", "copy_file_groups")

hugo_theme(
    name = "hugo_theme_geekdoc",
    srcs = [
        "@theme_geekdoc//:files",
    ],
    theme_name = "hugo-geekdoc",
)

genrule(
    name = "copy_gen_files",
    srcs = ["//api:api_documents_tar"],
    outs = ["content/en/model"],
    cmd = "mkdir -p $(RULEDIR)/content/en/model && tar -C $(RULEDIR)/content/en/model -zxf $(SRCS)",
)

hugo_site(
    name = "docs",
    assets = glob(["assets/**"]),
    config_dir = glob(["config/**"]),
    content = glob(["content/**"]) + [":copy_gen_files"],
    data = glob(["data/**"]),
    images = glob(["images/**"]),
    layouts = glob(["layouts/**"]),
    quiet = False,
    static = glob(["static/**"]),
    theme = ":hugo_theme_geekdoc",
    verbose = True,
)

# Run local development server
hugo_serve(
    name = "serve",
    dep = [
        ":docs",
    ],
)

# Tar it up
# pkg_tar(
#     name = "documentation_tar",
#     srcs = [":docs"],
# )
