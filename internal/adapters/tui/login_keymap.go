package tui

import (
	"github.com/charmbracelet/bubbles/key"
	//tea "github.com/charmbracelet/bubbletea"
)

type LoginTextKeys struct {
	Quit 		key.Binding
	Switch 		key.Binding
	Submit		key.Binding
	Help		key.Binding
}

func NewLoginTextKeys () LoginTextKeys {
	return LoginTextKeys {
		Quit: key.NewBinding(
			key.WithKeys("ctrl+c", "esc"),
			key.WithHelp("ctrl+c", "quit"),
			key.WithHelp("esc", "quit"),
		),
		Switch: key.NewBinding(
			key.WithKeys("tab"),
			key.WithHelp("tab", "switch"),
		),
		Submit: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "submit"),
		),
		Help: key.NewBinding(
			key.WithKeys("?"),
			key.WithHelp("?", "toggle help"),
		),
	}
}

func (k LoginTextKeys) ShortHelp() []key.Binding {
	return []key.Binding{k.Quit, k.Switch, k.Submit}
}

func (k LoginTextKeys) FullHelp() [][]key.Binding {
	return [][]key.Binding{{k.Quit, k.Switch, k.Submit}}
}

type LoginListKeys struct {
	Quit 		key.Binding
	Switch 		key.Binding
	Submit		key.Binding
	Up			key.Binding
	Down		key.Binding
	Help		key.Binding
}

func NewLoginListKeys () LoginListKeys {
	return LoginListKeys {
		Quit: key.NewBinding(
			key.WithKeys("ctrl+c", "esc"),
			key.WithHelp("ctrl+c / esc", "quit"),
		),
		Switch: key.NewBinding(
			key.WithKeys("tab"),
			key.WithHelp("tab", "switch"),
		),
		Submit: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "submit"),
		),
		Up: key.NewBinding(
			key.WithKeys("k", "up"),
			key.WithHelp("↓/k", "up"),
		),
		Down: key.NewBinding(
			key.WithKeys("j", "down"),
			key.WithHelp("↓/j", "down"),
		),
		Help: key.NewBinding(
			key.WithKeys("?"),
			key.WithHelp("?", "toggle help"),
		),
	}
}

func (k LoginListKeys) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k LoginListKeys) FullHelp() [][]key.Binding {
	return [][]key.Binding{{k.Help, k.Quit},{k.Switch, k.Submit, k.Up, k.Down}}
}
