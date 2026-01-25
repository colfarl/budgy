package app

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/colfarl/budgy/internal/database"
)

type User struct {
	name string
	id   int64
}

type Model struct {
	User         User
	Accounts     []database.Account
	Transactions []database.Account
	Logged_in    bool
	db           *database.Queries

	LoginInput textinput.Model
}

func NewModel(q *database.Queries) Model {
	ti := textinput.New()
	ti.Placeholder = "default"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return Model{
		User:         User{},
		Accounts:     nil,
		Transactions: nil,
		Logged_in:    false,
		db:           q,
		LoginInput:   ti,
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		tea.SetWindowTitle("Budgy"),
		textinput.Blink,
	)
}
