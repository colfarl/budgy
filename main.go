package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/colfarl/budgy/internal/database"
	"github.com/joho/godotenv"
	//"github.com/colfarl/budgy/internal/upload"
	"github.com/colfarl/budgy/internal/app"
	_ "github.com/mattn/go-sqlite3"
	tea "github.com/charmbracelet/bubbletea"
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
	
	if err != nil {
		fmt.Println("fatal:", err)
		os.Exit(1)
	}
	
	p := tea.NewProgram(app.NewApp(db, queries))
	if _, err := p.Run(); err != nil {
		fmt.Print("Unable to start budgy...\n", err)
		os.Exit(1)
	}
}
