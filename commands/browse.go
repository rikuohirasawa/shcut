package commands

import (
	"fmt"
	"os"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
	lipgloss "github.com/charmbracelet/lipgloss"
	internals "github.com/rikuohirasawa/shcut/internals"

	"github.com/spf13/cobra"
)

type listItem struct {
	name, script string
}

type model struct {
	options []listItem
	cursor  int // which alias our cursor is pointing at
}

func (i listItem) Title() string {
	return i.name
}
func (i listItem) Description() string {
	return i.script
}

func (i listItem) FilterValue() string {
	return i.name
}

func (m model) Init() tea.Cmd {
	return nil
}

func initialModel(config map[string]string) *model {
	var options []listItem
	fmt.Println(config)
	if len(config) == 0 {
		fmt.Println("No config found")
		return nil
	}
	for name, script := range config {
		options = append(options, listItem{name: name, script: script})
		fmt.Println(name, script)
	}

	return &model{
		options: options,
		cursor:  0,
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		switch msg.String() {

		// exit the program.
		case "ctrl+c", "q", "esc":
			return m, tea.Quit

		// move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// move the cursor down
		case "down", "j":
			if m.cursor < len(m.options)-1 {
				m.cursor++
			}

		case "enter", " ":
			script := m.options[m.cursor].script
			cmd := exec.Command("sh", "-c", script)
			cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
			cmd.Run()
			return m, tea.Quit

		}
	}

	return m, nil
}

var (
	infoStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#5FAFD7")) // Blue
	cursorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF87D7")) // Pink
	nameStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#87D787")) // Green
	scriptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")) // Gray
)

func (m model) View() string {
	s := ("\n" + infoStyle.Render("Select a shortcut to execute") + "\n\n")

	for i, choice := range m.options {
		cursor := " "
		if m.cursor == i {
			cursor = cursorStyle.Render(">")
		}

		name := nameStyle.Render(choice.name)
		script := scriptStyle.Render("(" + choice.script + ")")

		s += fmt.Sprintf("%s %s - %s\n", cursor, name, script)
	}

	s += "\n" + infoStyle.Render("Press Ctrl+c to quit.") + "\n"
	return s
}

func Tea(config map[string]string) {
	m := initialModel(config)
	if m == nil {
		fmt.Println("No config found")
		return
	}
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func Browse(configFilePath string) *cobra.Command {
	return &cobra.Command{
		Use:     "browse",
		Aliases: []string{"browse", "list", "ls", "show"},
		Short:   "Browse the shortcuts",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			config, err := internals.LoadConfig(configFilePath)
			if err != nil {
				return err
			}

			Tea(config)
			return nil
		},
	}
}
