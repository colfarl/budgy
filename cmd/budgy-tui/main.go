package main 

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/colfarl/budgy/internal/database"
	"github.com/colfarl/budgy/internal/store"
	"github.com/joho/godotenv"

	//"github.com/colfarl/budgy/internal/upload"
	"github.com/colfarl/budgy/internal/adapters/tui"
	tea "github.com/charmbracelet/bubbletea"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	err := godotenv.Load()
    if err != nil {
		fmt.Println("fatal:", err)
		os.Exit(1)
    }

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
	queries := database.New(db)
	
	s := store.New(db, queries)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go s.AppRun(ctx)	
		
	p := tea.NewProgram(tui.NewModel(s))
	if _, err := p.Run(); err != nil {
		fmt.Print("Unable to start budgy...\n", err)
		os.Exit(1)
	}
}
