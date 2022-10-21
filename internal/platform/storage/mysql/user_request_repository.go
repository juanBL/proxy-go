package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/huandu/go-sqlbuilder"
	proxy "zenrows-proxy/internal"
)

type UserRequestRepository struct {
	db        *sql.DB
	dbTimeout time.Duration
}

func NewUserRequestRepository(db *sql.DB, dbTimeout time.Duration) *UserRequestRepository {
	return &UserRequestRepository{
		db:        db,
		dbTimeout: dbTimeout,
	}
}

func (r *UserRequestRepository) Save(ctx context.Context, user proxy.UserRequest) error {
	userRequestSQLStruct := sqlbuilder.NewStruct(new(sqlUserRequest))
	query, args := userRequestSQLStruct.InsertInto(sqlUserRequestTable, sqlUserRequest{
		ApiKey:  user.ApiKey().String(),
		Url:     user.Url().String(),
		Headers: user.Headers().String(),
	}).Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	_, err := r.db.ExecContext(ctxTimeout, query, args...)
	if err != nil {
		return fmt.Errorf("error trying to persist user on database: %v", err)
	}

	return nil
}
