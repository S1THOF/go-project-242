### Hexlet tests and linter status:
[![Actions Status](https://github.com/S1THOF/go-project-242/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/S1THOF/go-project-242/actions)

# Analize-path-size

A simple CLI tool to display the size of files and directories in bytes or human-readable format.

## Installation


```bash
git clone https://github.com/S1THOF/Analize-path-size.git
cd Analize-path-size
make build
```

## Usage

```bash
hexlet-path-size [FLAGS] PATH
```

### Examples

Show file size in bytes:
```bash
hexlet-path-size data.csv
# Output: 81B	data.csv
```

Show size in human-readable format:
```bash
hexlet-path-size --human data.csv
# Output: 81B	data.csv

hexlet-path-size --human large-file.bin
# Output: 24.0MB	large-file.bin
```

Calculate directory size recursively with hidden files:
```bash
hexlet-path-size -r -a --human /path/to/directory
```

## Flags

| Flag | Short | Description |
|------|-------|-------------|
| `--human` | `-H` | Display sizes in human-readable format (auto-selects appropriate unit: B, KB, MB, GB, etc.) |
| `--recursive` | `-r` | Recursively calculate directory sizes (includes all subdirectories) |
| `--all` | `-a` | Include hidden files and directories (those starting with a dot `.`) |

## Format Details

### Byte Format
- Displays exact size in bytes followed by `B`
- Example: `12345B	filename.txt`

### Human-Readable Format
- Automatically selects the largest appropriate unit where the value is ≥ 1
- Units: B → KB → MB → GB → TB → PB → EB (base-2, 1 KB = 1024 bytes)
- Values < 1024 bytes: displayed as whole numbers (e.g., `1000B`)
- Values ≥ 1024 bytes: displayed with one decimal place (e.g., `1.2KB`, `24.0MB`)

## Exit Codes

- `0`: Success
- `1`: Error (file not found, permission denied, invalid arguments, etc.)

## Requirements

- Go 1.22 or higher

## License

MIT
