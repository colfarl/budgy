package app

import (
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

func (l *LoginModel) getNameFromLogin() (loginUserReq, error) {
	var resp loginUserReq 

	switch l.state {
	case InputFocus:
		resp = loginUserReq{
			name: l.Input.Value(),
			existing: false,
		}
		
	case ListFocus:
		if item, ok := l.UserList.SelectedItem().(ui.LoginUserItem); ok {
			resp = loginUserReq{
				name: string(item),
				existing: true,
			}
		} 
	} 
	
	if err := isValidUserName(resp.name); err != nil {
		return loginUserReq{}, err
	}
	return resp, nil
}

func (l *LoginModel) HandleLogin() (*LoginModel, tea.Cmd) {
	req, err := l.getNameFromLogin()	
	if err != nil {
		l.Error = err
		return l, nil
	}	

	return l, func() tea.Msg {return req}
}

func (a *App) LoadLoginFromInput(allUsers []string) (*App, tea.Cmd) {
	a.Login.AllUsers = allUsers
	a.Login.state = InputFocus
	a.Login.Input = ui.NewLoginInput(a.theme)
	a.Login.Input.Focus()
	a.Login.UserList = ui.NewUsersList(a.theme, allUsers, a.width, a.height / 8)
	return a, nil	
}
