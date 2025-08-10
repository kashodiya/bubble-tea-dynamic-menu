package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Config represents the structure of our configuration file
type Config struct {
	Commands []Command `json:"commands"`
}

// Command represents a single command in our configuration
type Command struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ShellCmd    string `json:"command,omitempty"`
	Script      string `json:"script,omitempty"`
}

// Define some styles
var (
	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#25A065")).
			Padding(0, 1).
			Bold(true)

	itemStyle        = lipgloss.NewStyle().Padding(0, 2)
	selectedItemStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#FFFDF5")).
				Background(lipgloss.Color("#2D7D9A")).
				Padding(0, 2)

	descriptionStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#A49FA5"))
)

// Model represents the state of our application
type model struct {
	config      Config
	cursor      int
	selected    int
	output      string
	showOutput  bool
	outputTitle string
}

// Init initializes the model
func (m model) Init() tea.Cmd {
	return nil
}

// Update handles events and updates the model
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			if m.showOutput {
				m.showOutput = false
				return m, nil
			}
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.config.Commands)-1 {
				m.cursor++
			}

		case "enter", " ":
			m.selected = m.cursor
			cmd := m.config.Commands[m.selected]
			
			var output string
			var err error
			
			if cmd.ShellCmd != "" {
				output, err = executeShellCommand(cmd.ShellCmd)
				m.outputTitle = "Command Output: " + cmd.ShellCmd
			} else if cmd.Script != "" {
				output, err = executeScript(cmd.Script)
				m.outputTitle = "Script Output: " + cmd.Script
			}
			
			if err != nil {
				output = fmt.Sprintf("Error: %s", err)
			}
			
			m.output = output
			m.showOutput = true
		}
	}

	return m, nil
}

// View renders the UI
func (m model) View() string {
	if m.showOutput {
		return renderOutput(m)
	}
	return renderMenu(m)
}

// renderMenu renders the main menu
func renderMenu(m model) string {
	s := titleStyle.Render(" Command Menu ") + "\n\n"

	for i, cmd := range m.config.Commands {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
			s += selectedItemStyle.Render(fmt.Sprintf("%s %s", cursor, cmd.Name)) + "\n"
			s += descriptionStyle.Render(fmt.Sprintf("   %s", cmd.Description)) + "\n\n"
		} else {
			s += itemStyle.Render(fmt.Sprintf("%s %s", cursor, cmd.Name)) + "\n"
			s += descriptionStyle.Render(fmt.Sprintf("   %s", cmd.Description)) + "\n\n"
		}
	}

	s += "\nPress q to quit, enter to execute\n"
	return s
}

// renderOutput renders the command output screen
func renderOutput(m model) string {
	s := titleStyle.Render(m.outputTitle) + "\n\n"
	s += m.output + "\n\n"
	s += "Press q to go back to menu\n"
	return s
}

// executeShellCommand executes a shell command and returns its output
func executeShellCommand(command string) (string, error) {
	cmd := exec.Command("bash", "-c", command)
	output, err := cmd.CombinedOutput()
	return string(output), err
}

// executeScript executes a script and returns its output
func executeScript(scriptPath string) (string, error) {
	// Check if the script exists
	if _, err := os.Stat(scriptPath); os.IsNotExist(err) {
		// Try with full path
		scriptPath = fmt.Sprintf("%s/%s", getCurrentDir(), scriptPath)
		if _, err := os.Stat(scriptPath); os.IsNotExist(err) {
			return "", fmt.Errorf("script not found: %s", scriptPath)
		}
	}

	cmd := exec.Command("bash", scriptPath)
	output, err := cmd.CombinedOutput()
	return string(output), err
}

// getCurrentDir returns the current directory
func getCurrentDir() string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}
	return dir
}

// loadConfig loads the configuration from a JSON file
func loadConfig(filePath string) (Config, error) {
	var config Config

	// Read the file
	data, err := os.ReadFile(filePath)
	if err != nil {
		return config, err
	}

	// Parse the JSON
	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func main() {
	// Load configuration
	configPath := "config.json"
	config, err := loadConfig(configPath)
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		os.Exit(1)
	}

	// Initialize the model
	m := model{
		config:   config,
		cursor:   0,
		selected: 0,
	}

	// Run the program
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v\n", err)
		os.Exit(1)
	}
}
