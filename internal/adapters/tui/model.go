package tui

import (
	"github.com/colfarl/budgy/internal/core"
	"github.com/colfarl/budgy/internal/store"
	"github.com/colfarl/budgy/internal/adapters/tui/ui"
	tea "github.com/charmbracelet/bubbletea"
)


type ScreenState int

const (
	None ScreenState = iota
	Login 
	Home
)

type Model struct {
	store			*store.Store
	state 			ScreenState
	
	ActiveTheme 	*ui.Theme
	AllUsers		*[]string
	Login			*LoginScreen	

	w				int
	h				int
}

func NewModel(s *store.Store) Model {
	return Model{
		store: s,
		state: None,
		ActiveTheme: ui.DefaultTheme(),
		Login: &LoginScreen{State: LoginState(Unloaded)},
		w: 30, // TODO: Find better defaults
		h: 30,
	}
}

func eventCatch(ch <-chan core.Event) tea.Cmd {
	return func () tea.Msg {
		e := <-ch
		return e
	}
}

func (m Model) Init() tea.Cmd {
	m.store.Commands <- core.LoadSession{}
	return eventCatch(m.store.Events)
}
