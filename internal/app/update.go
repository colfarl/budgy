package app

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case tea.KeyMsg:
		if !m.Logged_in {
			if msg.String() == "enter" {
				val := m.LoginInput.Value()
				if val == "" {
					return m, nil
				}
				// Handle login
				m.Logged_in = true
				return m, nil
			}
		} // no logged in state yet

		if s := msg.String(); s == "ctrl+c" || s == "esc" {
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		return m, nil
	}

	m.LoginInput, cmd = m.LoginInput.Update(msg)
	return m, cmd
}
