Basic Task CLI Application To learn some go basics applying the repository pattern for changing easy databases

// Build
go build -o task ./cmd/main.go


// TODOs:
-- Add Postgresql Support
-- Add Custom File Support (FileDB)
-- Make Arguments with '--' be orderless


Usage:
  task add [--priority=high|normal|low] <description>
  task list [--status=all|pending|completed]
  task complete <id>
  task delete <id>

Examples:
  task add --priority=high "Call dentist"
  task add "Buy groceries"
  task list --status=pending

