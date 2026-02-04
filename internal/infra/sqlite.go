package infra

import (
	"context"
	"database/sql"
	"errors"

	"github.com/colfarl/budgy/internal/core"
	"github.com/colfarl/budgy/internal/database"
)

type SqliteRunner struct{
	DB		*sql.DB
	Q		*database.Queries
}

var ErrNilUsername = errors.New("SQLITE RUNNER: cannot set nil username")

func (sr SqliteRunner) Run(ctx context.Context, fx core.Effect, emit func(core.Event)) bool {
	switch v := fx.(type) {
	case core.FxLoadSession:
		return sr.sqliteLoadSession(ctx, emit)	

	case core.FxSetSessionUser:
		return sr.sqliteSetUser(ctx, v,  emit)	

	case core.FxClearSessionUser:
		return sr.sqliteClearUser(ctx, emit)	

	case core.FxLoadAllUsers:
		return sr.sqliteLoadUsers(ctx, emit)	
	}
	return false
}

func (sr SqliteRunner) sqliteClearUser(ctx context.Context, emit func(core.Event)) bool {	
	if err := sr.Q.LogoutSession(ctx); err != nil {
		emit(core.DBFailure{Err: err})
		return true 
	}

	emit(core.ActiveUserCleared{})
	return true
}

func (sr SqliteRunner) sqliteSetUser(ctx context.Context, fx core.FxSetSessionUser, emit func(core.Event)) bool {
	if fx.Username == nil {
		emit(core.DBFailure{Err: ErrNilUsername})
		return true 
	}
	err := sr.Q.LoginSession(ctx, sql.NullString{String: *fx.Username, Valid: true})
	if err != nil {
		emit(core.DBFailure{Err: err})
		return false
	}

	emit(core.ActiveUserSet{Username: *fx.Username})
	return true
}

func (sr SqliteRunner) sqliteLoadUsers(ctx context.Context, emit func(core.Event)) bool {
	users, err := sr.Q.GetAllUserNames(ctx)	
	if err != nil {
		emit(core.DBFailure{Err: err})
		return true 
	}	
	
	emit(core.UsersLoaded{Usernames: users})
	return true
}

func (sr SqliteRunner) sqliteLoadSession(ctx context.Context, emit func(core.Event)) bool {
	prev, err := sr.Q.LoadSession(ctx)	
	if err != nil || !prev.Valid {
		emit(core.SessionLoadFailed{Err: err})
		return true 
	}

	if err = sr.Q.UpdateSessionLastOpened(ctx); err != nil {
		emit(core.DBFailure{Err: err})
		return true 
	}
	
	emit(core.ActiveUserSet{Username: prev.String})
	return true
}

