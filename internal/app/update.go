package app

import (
	tea "github.com/charmbracelet/bubbletea"
)
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {	    
	case tea.KeyMsg:
		switch msg.String() {
		default:	
			return m, tea.Quit
		}
	default:
		return m, nil
	}
}
