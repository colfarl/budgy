package app

import (
	"context"
	"fmt"

	//"github.com/colfarl/budgy/internal/ui"
	tea "github.com/charmbracelet/bubbletea"
)

/*
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (m *App) SetUpLogin() error {
	// move into get all accounts wrapper that checks if model has loaded this already
	registered_names, err := m.db.GetAllUserNames(context.Background())
	if err != nil {
		return err
	}

	ti := ui.NewLoginInput(*ui.DefaultTheme())
	lst := ui.NewUsersList(*ui.DefaultTheme(), registered_names, m.width, 8)
	ti.Focus()
	m.Login.Input = ti
	m.Login.UserList = lst
	return nil
}
*/

type loginGetUsersReq struct{}

func (l *LoginModel) Update(msg tea.Msg) (*LoginModel, tea.Cmd) {
	switch l.state {
	case NotLoaded:
		return l, func() tea.Msg {return loginGetUsersReq{}} 
	}
	return l, nil
}

// App Level update
func (a *App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	//Top Level Messages
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		a.width, a.height = msg.Width, msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return a, tea.Quit
		}

	case loginGetUsersReq:
		var cmd tea.Cmd
		allUsers, err := a.db.GetAllUserNames(context.Background())
		if err != nil {
			a.Error = err
			return a, nil
		}

		a, cmd = a.LoadLoginState(allUsers) 
		a.state = StateLoginLoaded
		return a, cmd 
	}
	
	//No messages received
	switch a.state {
	case StateInitial:
		fmt.Println("STATE INITIAL")
		var cmd tea.Cmd
		a.Login, cmd = a.Login.Update(msg)
		return a, cmd
	default:
		return a, nil
	}
}
