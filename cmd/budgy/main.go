package main

import (
	"os"
	"context"
	"log"
	"time"

	"github.com/alecthomas/kong"
	"github.com/colfarl/budgy/internal/adapters/cli"
	"github.com/colfarl/budgy/internal/infra"
	"github.com/colfarl/budgy/internal/store"
)


func main() {
	db, q, err := infra.LoadDb()	
	if err != nil {
		log.Fatal("FATAL:", err)
	}
	defer db.Close()

	file, err := os.OpenFile("cli.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer file.Close() 
	log.SetOutput(file)

	s := store.New(db, q)
	
	appctx := context.Background()
	ctx, cancel := context.WithCancel(appctx)
	defer cancel()
	go s.AppRun(ctx)	
		
	ctx, timeout := context.WithTimeout(appctx, 2 * time.Second)
	defer timeout()

	var input cli.CliCommands
	cmd := kong.Parse(&input)
	err = cmd.Run(&cli.Context{Store: s, Ctx: ctx})
	cmd.FatalIfErrorf(err)
}
