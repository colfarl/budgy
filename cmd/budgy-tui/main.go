package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/colfarl/budgy/internal/store"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/colfarl/budgy/internal/adapters/tui"
	"github.com/colfarl/budgy/internal/infra"
	_ "github.com/mattn/go-sqlite3"
)



func main() {	
	f, err := infra.SetLogger()
	if err != nil {
		log.Fatal("FATAL:", err)
	}
	defer f.Close()

	db, queries, err := infra.LoadDb()
	if err != nil {
		log.Fatal("FATAL:", err)
	}
	defer db.Close()
	
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
