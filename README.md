# llog

A simple life logger. Use it as your dev log or a timestamped journal.

All logs are stored locally in a SQLite file.

![screenshot](./assets/screenshot.png)

## Installation

### Download Binary

Download binary from [release](https://github.com/ethn1ee/llog/releases/tag/v0.0.2). Available for macOS, Linux, and Windows on both arm64 and amd64.

### Go Install

```sh
go install github.com/ethn1ee/llog
```

## Usage

### Adding Entry

```sh
llog add "trying out llog"

```

### Viewing Entry

- View all entries

```sh
llog get
```

- View entries from today

```sh
llog get --today
```

- View entries from yesterday

```sh
llog get --yesterday
```

- View entries since September 2025

```sh
llog get --from 2025-09-01
```

- View entries until September 2025

```sh
llog get --to 2025-09-01
```

- View entries from a specific date range

```sh
llog get --from 2025-09-01 --to 2025-09-31
```

## WIP

These are features I plan to add:

- `llog summarize`: Summarizing entries with an LLM (can be useful to summarize a day for a standup meeting)
- `llog search`: Fuzzy finding past logs
- `llog config`: Viewing config and db locations
- `llog reset`: Deleting all entries from data
