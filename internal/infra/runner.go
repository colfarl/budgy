package infra

import (
	"fmt"
	"log"
	"context"
	"database/sql"

	"github.com/colfarl/budgy/internal/core"
	"github.com/colfarl/budgy/internal/database"
)

type Runner interface {
	Run(ctx context.Context, fx core.Effect, emit func(core.Event)) bool
}

type EffectRunner []Runner 

func (er EffectRunner) Run(ctx context.Context, fx core.Effect, emit func(core.Event)) {
	for _, runner := range er {
		if runner.Run(ctx, fx,  emit) {
			log.Printf("EFFECT: %T, PROCESSED", fx)
			return 
		}
	}
	log.Printf("EFFECT: %T, IGNORED", fx)
	emit(core.EffectUnhandled{Kind: fmt.Sprintf("%T", fx)})
}

func NewEffectRunner(db *sql.DB, querier *database.Queries) EffectRunner {
	return []Runner{
		SqliteRunner{DB: db, Q: querier},
	}
}
