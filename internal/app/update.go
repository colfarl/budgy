package app

import (
	"context"
	"database/sql"
	//"fmt"

	//"github.com/colfarl/budgy/internal/ui"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/colfarl/budgy/internal/database"
)

type loginCheckCache struct{}
type loginGetUsersReq struct{}
type loginUserReq struct{name string; existing bool}

func (l *LoginModel) Update(msg tea.Msg) (*LoginModel, tea.Cmd) {
	var cmd tea.Cmd
	switch l.state {
	case NotLoaded:
		return l, func() tea.Msg {return loginCheckCache{}} 
	case CacheChecked:
		return l, func() tea.Msg {return loginGetUsersReq{}} 
	}
	
	if msg, ok := msg.(tea.KeyMsg); ok && msg.String() == "enter" {
		return l.HandleLogin()
	}
	l.Input, cmd = l.Input.Update(msg)
	return l, cmd
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
	
	case loginCheckCache:
		username, err := a.Store.Q.LoadSession(context.Background())
		if err != nil {
			a.Login.Error = err
			return a, nil
		} else if !username.Valid {
			a.Login.state = CacheChecked  
			return a, nil
		}
		user, err := a.Store.LoginAndGetUser(context.Background(), username)
		if err != nil {
			a.Login.state = CacheChecked 
			a.Login.Error = err
			return a, nil
		}
		a.User = user
		sessionParam := sql.NullString{String: a.User.Name, Valid: true}
		if err = a.Store.Q.LoginSession(context.Background(), sessionParam); err != nil {
			a.Login.Error = err
			return a, nil
		}
		a.Login.state = Done
		a.state = StateLoggedIn
		return a, nil

	case loginGetUsersReq:
		var cmd tea.Cmd
		allUsers, err := a.Store.Q.GetAllUserNames(context.Background())
		if err != nil {
			a.Error = err
			return a, nil
		}

		a, cmd = a.LoadLoginFromInput(allUsers) 
		a.Login.Input.Focus()	
		a.Login.state = InputFocus 
		return a, cmd 
	
	case loginUserReq:
		var user database.User
		var err error

		if msg.existing {
			user, err = a.Store.Q.GetUserByName(context.Background(), msg.name)
			if err != nil {
				a.Login.Error = err
				return a, nil
			}
		} else {
			user, err = a.Store.Q.CreateUser(context.Background(), msg.name)
			if err != nil {
				a.Login.Error = err
				return a, nil
			}
		}		

		a.User = user
		sessionParam := sql.NullString{String: a.User.Name, Valid: true}
		if err = a.Store.Q.LoginSession(context.Background(), sessionParam); err != nil {
			a.Login.Error = err
			return a, nil
		}
		a.Login.state = Done
		a.state = StateLoggedIn
		return a, nil
	}
	
	//No messages received
	switch a.state {
	case StateInitial:
		var cmd tea.Cmd	
		a.Login, cmd = a.Login.Update(msg)	
		return a, cmd
	default:
		return a, nil
	}
}
