package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/huandu/go-sqlbuilder"
	proxy "zenrows-proxy/internal"
)

// UserRepository is a MySQL proxy.UserRepository implementation.
type UserRepository struct {
	db        *sql.DB
	dbTimeout time.Duration
}

// NewUserRepository initializes a MariaDB-based implementation of proxy.UserRepository.
func NewUserRepository(db *sql.DB, dbTimeout time.Duration) *UserRepository {
	return &UserRepository{
		db:        db,
		dbTimeout: dbTimeout,
	}
}

func (r *UserRepository) Save(ctx context.Context, user proxy.User) error {
	userSQLStruct := sqlbuilder.NewStruct(new(sqlUser))
	query, args := userSQLStruct.InsertInto(sqlUserTable, sqlUser{
		ApiKey:         user.ApiKey().String(),
		ExpirationDate: user.ExpirationDate().String(),
	}).Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	_, err := r.db.ExecContext(ctxTimeout, query, args...)
	if err != nil {
		return fmt.Errorf("error trying to persist user on database: %v", err)
	}

	return nil
}

func (r *UserRepository) FindByApiKey(ctx context.Context, apiKey proxy.ApiKey) (users []proxy.User, err error) {
	userSQLStruct := sqlbuilder.NewStruct(new(sqlUser))
	selectBuilder := userSQLStruct.SelectFrom(sqlUserTable)
	selectBuilder = selectBuilder.Where(selectBuilder.Equal("api_key", apiKey.String()))

	query, args := selectBuilder.Build()

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	users = []proxy.User{}

	for rows.Next() {
		var sqlUser sqlUser
		err := rows.Scan(userSQLStruct.Addr(&sqlUser)...)

		if err != nil {
			return nil, err
		}
		user, err := proxy.NewUser(sqlUser.ApiKey, sqlUser.ExpirationDate)
		if err != nil {
			return nil, err
		}
		users = append(users, user)

	}
	return users, nil
}

func (r *UserRepository) SearchAll(ctx context.Context) (users []proxy.User, err error) {
	userSQLStruct := sqlbuilder.NewStruct(new(sqlUser))
	selectBuilder := userSQLStruct.SelectFrom(sqlUserTable)

	query, args := selectBuilder.Build()

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	users = []proxy.User{}

	for rows.Next() {
		var sqlUser sqlUser
		err := rows.Scan(userSQLStruct.Addr(&sqlUser)...)

		if err != nil {
			return nil, err
		}
		user, err := proxy.NewUser(sqlUser.ApiKey, sqlUser.ExpirationDate)
		if err != nil {
			return nil, err
		}
		users = append(users, user)

	}
	return users, nil
}
