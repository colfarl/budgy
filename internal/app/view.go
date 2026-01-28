package app

import (
	//"fmt"
	"github.com/charmbracelet/lipgloss"
)

var quitStr string = "(ctrl-c or esc to quit)"
var redStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("9"))

func (a *App) View() string {
	viewStr := ""

	switch a.state{
	case StateInitial:
		switch a.Login.state {
		case InputFocus, ListFocus:
			viewStr += a.Login.Input.View() + "\n\n" + a.Login.UserList.View()
		default:
			viewStr += "..." // loading or done
		}
	case StateLoginLoaded:
	default:
		viewStr += "Welcome to Budgy"
	}	
	return viewStr + "\n" + quitStr
}
