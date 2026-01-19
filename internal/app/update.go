package app

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {	    
	
	case tea.KeyMsg:
		if !m.Logged_in && msg.String() == "enter" {
			val := m.LoginInput.Value()			
			if val == "" {
				return m, nil
			}
			m.Logged_in = true
			return m, nil
		} 

		if s := msg.String(); s == "ctrl+c" || s == "esc" {
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		return m, nil	
	}

	m.LoginInput, cmd = m.LoginInput.Update(msg)
	return m, cmd
}
