package tui

import (
	//"log"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"

	//tea "github.com/charmbracelet/bubbletea"
	"github.com/colfarl/budgy/internal/adapters/tui/ui"
)

type LoginState int

const (
	Unloaded LoginState = iota
	Input 
	UserList	
)

type LoginScreen struct{
	TextInput		textinput.Model
	UserList		list.Model	
	State 			LoginState	
	Help 			help.Model
	TIKeymap		LoginTextKeys			
	ListKeymap		LoginListKeys			
}

func (l LoginScreen) GetActiveHelp () help.KeyMap {
	switch l.State {
		case Input:
			return l.TIKeymap
		case UserList:
			return l.ListKeymap
		default:
			return nil
	}
}

func NewLoginScreen(t ui.Theme, names []string, width, height int) *LoginScreen {
	return &LoginScreen{
		TextInput:  ui.NewLoginInput(t),			
		UserList:	ui.NewUsersList(t, names, width, height / 4),
		State: 		LoginState(Input),
		Help: 		help.New(),
		TIKeymap: 	NewLoginTextKeys(),
		ListKeymap: NewLoginListKeys(),
	}
}
