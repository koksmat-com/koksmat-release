---
title: koksmat-release
---

# Koksmat Release Management

The release management for the koksmat project is supported by this tools. 

## Usage

```bash
koksmat-release
```

Outputs the current options

```
Usage:
  koksmat-release [command]

Available Commands:
  bump        Bump
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  release     Release
  serve       Serve the API
```

### Bump
Used for bumping the version of the project. Bump comes with the following options:

```
Usage:
  koksmat-release bump minor|major|patch [kitchen]

```

Sample usage:

```bash
koksmat-release bump minor .
```

This will bump the minor version of the project in the current directory. 

Version information is stored in the `./koksmat/koksmat.json" file. 

```json
{
  "version": {
    "minor": 0,
    "build": 0,
    "patch": 1,
    "major": 0
  }
}
```