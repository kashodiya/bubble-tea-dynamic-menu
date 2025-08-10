# Bubble Tea TUI Command Runner

A Terminal User Interface (TUI) application built with the [Bubble Tea](https://github.com/charmbracelet/bubbletea) library that allows users to select and run commands from a configuration file.

## Features

- Read commands from a JSON configuration file
- Display a menu of available commands with descriptions
- Execute shell commands or scripts
- View command output in a separate screen
- Simple and intuitive keyboard navigation

## Configuration File Format

The application reads a `config.json` file with the following structure:

```json
{
  "commands": [
    {
      "name": "Command Name",
      "description": "Command Description",
      "command": "shell command to execute"
    },
    {
      "name": "Script Name",
      "description": "Script Description",
      "script": "path/to/script.sh"
    }
  ]
}
```

Each command entry must have:
- `name`: The name of the command shown in the menu
- `description`: A short description of what the command does
- Either `command` (a shell command to execute) or `script` (path to a script file)

## Usage

1. Create a `config.json` file in the same directory as the application
2. Run the application: `./tui-app`
3. Use arrow keys or j/k to navigate the menu
4. Press Enter to execute the selected command
5. Press q to go back to the menu or quit the application

## Building from Source

```bash
# Clone the repository
git clone <repository-url>
cd bubbletea-tui

# Install dependencies
go mod tidy

# Build the application
go build -o tui-app

# Run the application
./tui-app
```

## Requirements

- Go 1.22 or later
- Bubble Tea library
- Lipgloss library

## License

MIT
