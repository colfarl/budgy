package app

import (
	"database/sql"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/colfarl/budgy/internal/store"
	"github.com/colfarl/budgy/internal/database"
	"github.com/colfarl/budgy/internal/ui"
)

type LoginState int

const (
	NotLoaded LoginState = iota
	CacheChecked
	InputFocus
	ListFocus
	Done
)

type LoginModel struct {	
	AllUsers     	[]string
	Input 			textinput.Model
	UserList	   	list.Model
	
	state			LoginState
	Error 		 	error
}

type AppState int

const (
	StateInitial 	AppState = iota	
	StateLoginLoaded
	StateLoggedIn
)

type App struct {
	User			database.User 				
	Transactions 	[]database.Transaction
	Accounts 		[]database.Account
	Store           *store.Store	

	state 			AppState
	Error 		 	error
	
	Login			*LoginModel
	height		 	int	
	width		 	int	
	theme			ui.Theme
}

func NewApp(db *sql.DB, q *database.Queries) *App {	
	return &App{
		User:         	database.User{},
		Accounts:     	nil,
		Transactions: 	nil,
		Store:          store.New(db, q),
		
		Login: 			&LoginModel{},
		state: 			StateInitial,	
		height: 		64,
		width: 			64,
		theme: 			*ui.DefaultTheme(),
	}
}

func (a *App) Init() tea.Cmd {
	return tea.Batch(
		tea.SetWindowTitle("Budgy"),
		textinput.Blink,
	)
}
