package app

import (
	"github.com/colfarl/budgy/internal/database"
	tea "github.com/charmbracelet/bubbletea"
)
type Model struct {
	User			string
	Accounts		[]database.Account	
	Transactions	[]database.Account		
	Logged_in 		bool
}

func NewModel() Model {
	return Model {
		User: "",
		Accounts: nil,
		Transactions: nil,
		Logged_in: false,
	}
}

func (m Model) Init() tea.Cmd {
	return tea.SetWindowTitle("Budgy")	
}
