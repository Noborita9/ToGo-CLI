# ToGoCLI

A command-line task management application built in Go, implementing the repository pattern for flexible database backends.

## Features

- Add tasks with priority levels (high, normal, low)
- List tasks with filtering by status (pending, completed)
- Mark tasks as complete
- Delete tasks
- SQLite database backend with automatic table creation

## Build

```bash
go build -o task ./cmd/main.go
```

## Usage

```bash
task add [--priority=high|normal|low] <description>
task list [--status=pending|completed]
task complete <id>
task delete <id>
```

## Examples

```bash
task add --priority=high "Call dentist"
task add "Buy groceries"
task list --status=pending
task complete 1
task delete 2
```

## Architecture

The application follows clean architecture principles with:
- **Domain layer**: Task entity definitions
- **Repository layer**: Data access abstraction with SQLite implementation
- **Service layer**: Business logic
- **CLI layer**: Command-line interface handling

## Future Enhancements

- Add PostgreSQL support
- Add custom file-based database support
- Make command arguments order-independent

