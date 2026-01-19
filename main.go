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

func initDatabase() (*sql.DB, error) {
	dbURL := os.Getenv("DB_URL")		
	db, err := sql.Open("sqlite3", dbURL)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	db, err := initDatabase()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer db.Close()
	_ = database.New(db)
	
	p := tea.NewProgram(app.NewModel())
	if _, err := p.Run(); err != nil {
		fmt.Print("Unable to start budgy...\n", err)
		os.Exit(1)
	}
	/*
	_, err = upload.ReadCsvFile("nov-dec.csv")
	if err != nil {
		fmt.Println(err)
		return
	}	
	*/
}
