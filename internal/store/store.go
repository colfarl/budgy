package store

import (
	"context"
	"database/sql"
	"log"


	"github.com/colfarl/budgy/internal/core"
	"github.com/colfarl/budgy/internal/infra"
	"github.com/colfarl/budgy/internal/database"
)

type Store struct {
	State 			core.State

	Commands 		chan core.Command
	Events 			chan core.Event

	Runner			infra.EffectRunner
}

func New(db *sql.DB, queries *database.Queries) *Store {
	return &Store{
		State: 		core.NewState(),
		Commands: 	make(chan core.Command),
		Events: 	make(chan core.Event),
		Runner: 	infra.NewEffectRunner(db, queries),
	}
}

func (s *Store) emit(e core.Event) {
	log.Printf("EVENT: %T", e)
	s.State = core.Transition(s.State, e)	
	s.Events <- e
}

func (s *Store) AppRun(ctx context.Context) {
	for {
		select {
		case <- ctx.Done():
			return

		case cmd := <- s.Commands:
			log.Printf("COMMAND: %T", cmd)
			evs, fxs := core.Evaluate(s.State, cmd)
			for _, ev := range evs {
				s.emit(ev)
			}

			for _, fx := range fxs {
				s.Runner.Run(ctx, fx, s.emit) // potentially goroutine...
			}
		}
	}
}
