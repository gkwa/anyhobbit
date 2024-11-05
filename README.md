# AnyHobbit

CLI tool to generate Renovate bot configurations with preset strategies.

## Installation

```bash
go install github.com/gkwa/anyhobbit@latest
```

## Overview

AnyHobbit simplifies Renovate configuration by providing preset strategies for common dependency management patterns. Each preset is designed for different automation levels and dependency handling approaches.

View available commands:
```bash
anyhobbit --help
```

## Cheatsheet

Generate config with specific preset:
```bash
anyhobbit owl    # Aggressive updates including replacements
anyhobbit rat    # Auto-merge standard updates, no tests
anyhobbit dog    # Auto-merge and recreate PRs for standard updates
anyhobbit monkey # Auto-merge all updates including indirect deps
anyhobbit hamster # Auto-merge standard updates with recommended base
```

Specify output file:
```bash
anyhobbit owl -o custom.json
anyhobbit owl --outfile custom.json
```

Default output is `.renovaterc.json` in current directory.

## Preset Strategies

- **owl**: Most aggressive. Auto-merges all updates including replacements, recreates PRs when needed
- **rat**: Auto-merges standard updates without running tests
- **dog**: Auto-merges standard updates with PR recreation
- **monkey**: Auto-merges all updates including indirect dependencies
- **hamster**: Conservative auto-merging of standard updates using recommended base config

## Usage Example

```bash
# Show available commands and flags
anyhobbit --help

# Generate owl preset config
anyhobbit owl

# Check generated config
cat .renovaterc.json
```

Generated config will contain appropriate presets for Renovate bot to handle dependency updates according to the chosen strategy.