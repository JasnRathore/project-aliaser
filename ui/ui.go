package ui
import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	count int
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "a":
			m.count++
		}
	}
	return m, nil
}

func (m model) View() string {
	return fmt.Sprintf("Press 'a' to increment, 'q' to quit. Count: %d", m.count)
}

// RunBubbleTeaApp starts the Bubble Tea program and returns any errors
func RunBubbleTeaApp() error {
	p := tea.NewProgram(model{})

	// Run the Bubble Tea application
	return p.Start()
}
