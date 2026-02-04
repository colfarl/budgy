package tui

import (
	//"log"
	"errors"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/colfarl/budgy/internal/adapters/tui/ui"
	"github.com/colfarl/budgy/internal/core"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
        case "ctrl+c", "esc":
            return m, tea.Quit

        case "ctrl+l":
			m.store.Commands <- core.ClearActiveUser{}
            return m, nil
		}

		if m.state == Login && m.Login.State == UserList {
			switch {
			case key.Matches(msg, m.Login.ListKeymap.Down):
				m.Login.UserList.CursorDown()
				return m, nil				

			case key.Matches(msg, m.Login.ListKeymap.Up):
				m.Login.UserList.CursorUp()
				return m, nil				

			case key.Matches(msg, m.Login.ListKeymap.Switch):
				m.Login.TextInput.Focus()
				m.Login.State = Input 
				return m, nil

			case key.Matches(msg, m.Login.ListKeymap.Submit):
				item, ok  := m.Login.UserList.SelectedItem().(ui.LoginUserItem)
				if !ok {
					m.store.State.Error = errors.New("Invalid List Item: should not be possible")
					return m, nil
				}
				m.store.Commands <- core.SetActiveUser{Username: string(item)}

			case key.Matches(msg, m.Login.ListKeymap.Help):
				m.Login.Help.ShowAll = !m.Login.Help.ShowAll
				return m, nil
			}
		} else if m.state == Login && m.Login.State == Input {
			switch {
			case key.Matches(msg, m.Login.TIKeymap.Switch):
				m.Login.TextInput.Blur()
				m.Login.State = UserList

			case key.Matches(msg, m.Login.TIKeymap.Help):
				m.Login.Help.ShowAll = !m.Login.Help.ShowAll
				return m, nil

			case key.Matches(msg, m.Login.ListKeymap.Submit):
				username := m.Login.TextInput.Value()	
				m.store.Commands <- core.SetActiveUser{Username: username}
				m.Login.TextInput.Focus()

			}

			var cmd tea.Cmd
			m.Login.TextInput, cmd = m.Login.TextInput.Update(msg)
			return m, cmd 
		}


	case tea.WindowSizeMsg:
		m.w = msg.Width
		m.h = msg.Height
		return m, nil

	case core.ActiveUserSet:
		m.state = Home
		return m, eventCatch(m.store.Events)

	case core.ActiveUserCleared:
		m.state = Login
		return m, eventCatch(m.store.Events)
	
	case core.SessionLoadFailed:
		if msg.Err != nil {
			m.store.State.Error = msg.Err
			return m, nil
		}
		m.store.Commands <- core.LoadAllUsers{}
		return m, eventCatch(m.store.Events)		
	
	case core.UsersLoaded:
		m.Login = NewLoginScreen(*ui.DefaultTheme(), msg.Usernames, m.w, m.h)
		m.state = Login
		m.Login.TextInput.Focus()
		return m, tea.Batch(eventCatch(m.store.Events), textinput.Blink)
	
	}
		
	return m, nil
}


