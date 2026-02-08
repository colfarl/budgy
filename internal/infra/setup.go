package infra

import (
	"os"
	"database/sql"
	"github.com/joho/godotenv"
	"github.com/colfarl/budgy/internal/database"

	tea "github.com/charmbracelet/bubbletea"
) 


func LoadDb() (*sql.DB, *database.Queries, error) {
	err := godotenv.Load()
    if err != nil {
		return nil, nil, err
    }

	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("sqlite3", dbURL)
	if err != nil {
		return nil, nil, err	
	}
	queries := database.New(db)
	return db, queries, nil
}

func SetLogger() (*os.File, error) {
	f, err := tea.LogToFile("tui.log", "tui")
	if err != nil {
		return nil, err	
	}
	return f, nil
}
