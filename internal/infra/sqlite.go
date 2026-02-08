package infra

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/colfarl/budgy/internal/core"
	"github.com/colfarl/budgy/internal/database"
	"github.com/mattn/go-sqlite3"
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

	case core.FxCreateUser:
		return sr.sqliteCreateUser(ctx, v, emit)	

	case core.FxDeleteUser:
		return sr.sqliteDeleteUser(ctx, v, emit)	
	
	// ============================== Account ==============================
	case core.FxCreateAccount:
		return sr.sqliteCreateAccount(ctx, v, emit)	

	case core.FxDeleteAccount:
		return sr.sqliteDeleteAccount(ctx, v, emit)	
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

func (sr SqliteRunner) sqliteDeleteUser(ctx context.Context, fx core.FxDeleteUser, emit func(core.Event)) bool {
	log.Printf("running effect delete user")
	if fx.Username == nil {
		emit(core.DBFailure{Err: ErrNilUsername})
		return true
	}

	rows, err := sr.Q.DeleteUserByName(ctx, *fx.Username)	
	if err != nil {
		emit(core.DBFailure{Err: ErrNilUsername})
		return true
	} else if rows == 0 {
		// Unsure if this should be handled 
		log.Printf("ATTEMPTED TO DELETE NOT EXISTENT USER")
		return true
	}
	
	emit(core.UserDeleted{Username: *fx.Username})
	return true
}

func (sr SqliteRunner) sqliteCreateUser(ctx context.Context, fx core.FxCreateUser, emit func(core.Event)) bool {
	if fx.Username == nil {
		emit(core.DBFailure{Err: ErrNilUsername})
		return true
	}
	user, err := sr.Q.CreateUser(ctx, *fx.Username)	
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) {
			if sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
				// unique constraint violation do nothing
				return true 
			}
		} 
		emit(core.DBFailure{Err: ErrNilUsername})
		return true
	}		

	emit(core.UserCreated{Username: user.Name})
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

func (sr SqliteRunner) sqliteCreateAccount(ctx context.Context, fx core.FxCreateAccount, emit func(core.Event)) bool {
	user, err := sr.Q.GetUserByName(ctx, fx.Username)	
	if err != nil {
		emit(core.DBFailure{Err: err})
		return true 
	}
	params := database.CreateAccountParams{
		Name: fx.AccountName,	
		UserID: user.ID,
	}

	account, err := sr.Q.CreateAccount(ctx, params)
	if err != nil {
		emit(core.DBFailure{Err: err})
		return true
	}

	values := core.AccountCreated{
		UserID: user.ID,	
		Username: user.Name,	
		AccountName: account.Name,	
		AccountID: account.ID,	
	}
	emit(values)
	return true
}

func (sr SqliteRunner) sqliteDeleteAccount(ctx context.Context, fx core.FxDeleteAccount, emit func(core.Event)) bool {
	user, err := sr.Q.GetUserByName(ctx, fx.Username)	
	if err != nil {
		emit(core.DBFailure{Err: err})
		return true 
	}

	params := database.DeleteAccountParams{
		Name: fx.AccountName,	
		UserID: user.ID,
	}

	err = sr.Q.DeleteAccount(ctx, params)
	if err != nil {
		emit(core.DBFailure{Err: err})
		return true
	}

	values := core.AccountDeleted{
		UserID: user.ID,	
		Username: user.Name,	
		AccountName: fx.AccountName,	
	}
	emit(values)
	return true
}

