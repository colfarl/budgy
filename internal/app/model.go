package app

import (
	"github.com/colfarl/budgy/internal/database"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	User			string
	Accounts		[]database.Account	
	Transactions	[]database.Account		
	Logged_in 		bool

	LoginInput 		textinput.Model
}

func NewModel() Model {
	ti := textinput.New()
	ti.Placeholder = "Default"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return Model {
		User: "",
		Accounts: nil,
		Transactions: nil,
		Logged_in: false,
		LoginInput: ti,
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		tea.SetWindowTitle("Budgy"),
		textinput.Blink,
	)
}
