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

List all available presets:
```bash
anyhobbit zoo
```

Generate config with specific preset:
```bash
anyhobbit renovate owl    # Aggressive updates including replacements
anyhobbit renovate rat    # Auto-merge standard updates, no tests
anyhobbit renovate monkey # Auto-merge all updates including indirect deps
anyhobbit renovate penguin # Auto-merge with PR notifications
```

Specify output file:
```bash
anyhobbit renovate owl -o custom.json
anyhobbit renovate owl --outfile custom.json
```

Default output is `.renovaterc.json` in current directory.

## Preset Strategies

- **owl**: Most aggressive. Auto-merges all updates including replacements, recreates PRs when needed
- **rat**: Auto-merges standard updates without running tests
- **monkey**: Auto-merges all updates including indirect dependencies
- **penguin**: Auto-merges with PR notifications for tracking merged updates
- **rabbit**: Auto-merges all dependency types with PR recreation
- **tiger**: Auto-merges with branch merging to reduce PR noise
- **panda**: Uses pin range strategy with auto-merging
- **koala**: Auto-merges with npm/pnpm dedupe options

## Usage Example

```bash
# Show available commands and flags
anyhobbit --help

# List all preset strategies
anyhobbit zoo

# Generate owl preset config
anyhobbit renovate owl

# Check generated config
cat .renovaterc.json
```

Generated config will contain appropriate presets for Renovate bot to handle dependency updates according to the chosen strategy.