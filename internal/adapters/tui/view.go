package tui

import (
	//	tea "github.com/charmbracelet/bubbletea"
	//"log"

	"github.com/colfarl/budgy/internal/adapters/tui/ui"
)

func (m Model) View() string {
	switch m.state {
	case Login:
		return layoutLogin(*m.ActiveTheme, m.Login, m.store.State.Error)

	case Home:
		return "Logged in"
	}

	return "uh oh"
}

func layoutLogin(t ui.Theme, ls *LoginScreen, err error) string {
	var errorMsg string
	if err != nil {
		errorMsg = t.Text.Error.Render((err.Error() + "\n"))
	}
	
	header := t.Text.Strong.Render("Budgy - Budgeting in the Terminal")
	footer := ls.Help.View(ls.GetActiveHelp())
	return header + "\n" + ls.TextInput.View() + "\n\n" + ls.UserList.View() + "\n" + errorMsg + footer
}
