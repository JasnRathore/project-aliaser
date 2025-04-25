package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	
	"fmt"
	"strings"
	api "fa/api"
)

type model struct {
	cursor        int
	choices       []string
	mode          string
	searchText    string
	searchResults []api.Alias
	searchCursor  int
	status        string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.mode == "message" {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			if msg.Type == tea.KeyEnter {
				m.mode = ""
			}
		}
		return m, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.mode == "search" {
			switch msg.Type {
			case tea.KeyEsc:
				m.mode = ""
				m.searchText = ""
				m.searchResults = nil
				m.searchCursor = 0

			case tea.KeyEnter:
				if len(m.searchResults) > 0 {
					selected := m.searchResults[m.searchCursor]
					api.WriteToMidFile("cd", selected.Name)	
					return m, tea.Quit
				}
				m.searchCursor = 0

			case tea.KeyBackspace:
				if len(m.searchText) > 0 {
					m.searchText = m.searchText[:len(m.searchText)-1]
					results, err := api.FuzzySearchAlias(m.searchText)
					if err == nil {
						m.searchResults = results
						m.searchCursor = 0
					}
				}

			case tea.KeyUp, tea.KeyCtrlP:
				if m.searchCursor > 0 {
					m.searchCursor--
				}

			case tea.KeyDown, tea.KeyCtrlN:
				if m.searchCursor < len(m.searchResults)-1 {
					m.searchCursor++
				}

			default:
				m.searchText += msg.String()
				results, err := api.FuzzySearchAlias(m.searchText)
				if err == nil {
					m.searchResults = results
					m.searchCursor = 0
				}
			}
			return m, nil
		}

		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter":
			selected := m.choices[m.cursor]
			if selected == "Search" {
				m.mode = "search"
				m.searchText = ""
				m.searchResults = nil
				m.searchCursor = 0
			} else if selected == "Quit" {
				return m, tea.Quit
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	var b strings.Builder

	// Instructions at the top for navigation
	b.WriteString("  ↑/↓ or j/k to move • Enter to select • q to quit\n\n")

	switch m.mode {
	case "message":
		// Simple message display
		return fmt.Sprintf("\n  %s\n\n  Press Enter to continue", m.status)

	case "search":
		// Search header
		b.WriteString("Search Aliases\n")

		// Search input
		b.WriteString(fmt.Sprintf("  > %s\n\n", m.searchText))

		// No results or empty search input state
		if strings.TrimSpace(m.searchText) == "" {
			b.WriteString("  Start typing to search...\n")
		} else if len(m.searchResults) == 0 {
			b.WriteString("  No matches found.\n")
		} else {
			// Display search results
			for i, result := range m.searchResults {
				prefix := "  "
				if i == m.searchCursor {
					prefix = "→ "
				}
				b.WriteString(fmt.Sprintf("%s%s  (%s)\n", prefix, result.Name, result.Location))
			}
		}

		// Footer with instructions
		b.WriteString("\n  ↑/↓ to navigate • Enter to select • Esc to cancel")
		return b.String()

	default:
		// Menu options with cursor highlighting
		for i, choice := range m.choices {
			prefix := "  "
			if i == m.cursor {
				prefix = "→ "
			}
			b.WriteString(fmt.Sprintf("%s%s\n", prefix, choice))
		}

		// Instructions already at the top, so no need for another footer here
		return b.String()
	}
}

func RunBubbleTeaApp() error {
	m := model{
		choices: []string{"Search", "Quit"},
	}
	p := tea.NewProgram(m)
	return p.Start()
}
