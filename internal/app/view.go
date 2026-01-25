package app

import (
	"fmt"
)

var quitStr string = "(ctrl-c or esc to quit)"

func (m Model) View() string {
	viewStr := ""
	switch {
	case !m.Logged_in:
		viewStr += fmt.Sprintf(
			"Enter username \n\n%s\n",
			m.LoginInput.View(),
		) + "\n"
	default:
		viewStr += "Welcome to Budgy"
	}
	return viewStr + "\n" + quitStr
}
