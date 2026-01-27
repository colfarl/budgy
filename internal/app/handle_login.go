package app

import (
	//"context"
	"errors"
	"unicode"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/colfarl/budgy/internal/ui"
)

var ErrNameLength = errors.New("INVALID NAME: 3 <= length(username) <= 16")
var ErrBadCharacters = errors.New("INVALID NAME: name must be alphanumeric")

func isAlphaNum(str string) bool {
	for _, c := range str {
		if !unicode.IsLetter(c) && !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func isValidUserName(name string) error {
	if len(name) <= 2 || len(name) > 16 {
		return ErrNameLength
	} else if !isAlphaNum(name) {
		return ErrBadCharacters
	}
	return nil
}

// func (a *App) HandleLogin(name string) (tea.Model, tea.Cmd) {
// 	if err := isValidUserName(name); err != nil {
// 		a.Error = err		
// 		return a, nil
// 	}
// 
// 	user, err := a.db.CreateUser(context.Background(), name) 
// 	if err != nil {
// 		a.Error = err
// 		return a, nil
// 	}
// 
// 	a.User = user
// 	return a, nil
// }

func (a *App) LoadLoginFromInput(allUsers []string) (*App, tea.Cmd) {
	a.Login.AllUsers = allUsers
	a.Login.state = InputFocus
	a.Login.Input = ui.NewLoginInput(a.theme)
	a.Login.Input.Focus()
	a.Login.UserList = ui.NewUsersList(a.theme, allUsers, a.width, a.height / 8)
	return a, nil	
}
