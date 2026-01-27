package store

import (
	"context"
	"database/sql"
	"errors"

	"github.com/colfarl/budgy/internal/database"
)

type Store struct {
	db 		*sql.DB
	Q		*database.Queries
}

func New(db *sql.DB, q *database.Queries) *Store {
	return &Store {
		db: db, 
		Q: q,
	}
}

func (s *Store) LoginAndGetUser (ctx context.Context, username sql.NullString) (database.User, error) {
	if !username.Valid {
		return database.User{}, errors.New("Cannot login null user")
	}

	err := s.Q.LoginSession(ctx, username)
	if err != nil {
		return database.User{}, err
	}

	user, err := s.Q.GetUserByName(ctx, username.String)
	if err != nil {
		return database.User{}, err
	}
	return user, nil
}

// TODO: Write in transaction helpers
