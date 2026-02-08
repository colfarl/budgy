package main

import (
	"log"
	"context"

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

	s := store.New(db, q)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go s.AppRun(ctx)	
		
	cmd := kong.Parse(&cli.CliCommands)
	err = cmd.Run(&cli.Context{Debug: cli.CliCommands.Debug})
	cmd.FatalIfErrorf(err)
}
