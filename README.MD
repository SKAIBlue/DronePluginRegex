# Drone Plugin Regex
This is plugin for drone CI/CD

Tests that the commit message matches the given Regex


# Usage
``` yaml
kind: pipeline
type: docker
name: create-mr-for-feature

steps:
  - name: Check Commit Message
    image: jwmtp2/drone-plugin-regex
    settings:
      regex: "feat:.*"

  - name: You want to do if matches the given regex
    image: some/plugin
    ...
```