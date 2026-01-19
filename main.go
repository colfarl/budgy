package main

import (
	"fmt"
	"os"
	"database/sql"

	"github.com/colfarl/budgy/internal/database"
	//"github.com/colfarl/budgy/internal/upload"
	"github.com/colfarl/budgy/internal/app"
	_"github.com/mattn/go-sqlite3"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		fmt.Println("fatal:", err)
		os.Exit(1)
	}
	defer f.Close()

	dbURL := os.Getenv("DB_URL")		
	db, err := sql.Open("sqlite3", dbURL)
	if err != nil {
		fmt.Println("fatal:", err)
		os.Exit(1)
	}
	defer db.Close()
	_ = database.New(db)

	p := tea.NewProgram(app.NewModel())
	if _, err := p.Run(); err != nil {
		fmt.Print("Unable to start budgy...\n", err)
		os.Exit(1)
	}
}
