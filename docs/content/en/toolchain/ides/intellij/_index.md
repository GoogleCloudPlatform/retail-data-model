---
title: "Using IntelliJ Idea Ultimate"

resources:
- name: banner
  src: intellij-ide.png
  title: "IntelliJ IDE"
  params:
  credits: "[Ryan McGuinness](https://www.github.com/rrmcguinness)"
---

<div style="text-align: center">
{{< img name="banner" size="small" lazy=false >}}
</div>

## Prerequisites

* You have a licensed copy of [IntelliJ Idea Ultimate](https://www.jetbrains.com/idea/) running on your workstation.
* You've cloned the repository to a local directory.

| Pros                                                                                                      | Cons                          |
|-----------------------------------------------------------------------------------------------------------|-------------------------------|
| Support for all languages in a single IDE with Auto Complete, syntax highlighting, and build integration. | Requires a commercial license |

## Plugins

Install and activate the following plugins, this can be done from the startup screen:

* Bazel
* Go Language Support
* Python Language Support

> Restart our IDE after complete.

## Open the project

On the startup screen, you'll now see an "Open Bazel Project", click that button and open the
directory where you cloned the repository.

Choose "Import New Project"

When the template for the project is presented, copy the following values:

```yaml
directories:
  # Add the directories you want added as source here
  # By default, we've added your entire workspace ('.')
  .
  -bazel-*
  -.github*
  -.ijwb

# Automatically includes all relevant targets under the 'directories' above
derive_targets_from_directories: true

targets:
  # If source code isn't resolving, add additional targets that compile it here

test_sources:
  */test/go/*
  */test/java/*
  */test/python/*

java_language_level: 11

additional_languages:
  go
  python
  # Uncomment any additional languages you want supported
  # android
  # dart
  # go
  # javascript
  # kotlin
  # python
  # scala
  # typescript
```

Once open, you're ready to start contributing.




