package postgres

import (
	"context"
	"eSearcher/internal/models"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var LoginError = errors.New("wrong login or password")

type AuthDB struct {
	pool *pgxpool.Pool
}

func NewAuthStore(pool *pgxpool.Pool) *AuthDB {
	return &AuthDB{pool: pool}
}

func (db *AuthDB) CreateUser(login, password string, role int) (int, error) {
	ctx := context.Background()
	conn, err := db.pool.Acquire(ctx)
	if err != nil {
		return -1, err
	}
	defer conn.Release()
	var id int
	fmt.Println(login, password, role)
	if err = conn.QueryRow(ctx,
		`INSERT INTO users (login, password, role_id) VALUES($1, $2, $3) RETURNING id`,
		login, password, role).Scan(&id); err != nil {
		return -1, err
	}
	return id, nil
}

func (db *AuthDB) GetUser(login, password string) (*models.User, error) {
	ctx := context.Background()
	conn, err := db.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	var id, role int
	if err = conn.QueryRow(ctx,
		`SELECT id, role_id FROM users WHERE login = $1 AND password = $2`,
		login, password).Scan(&id, &role); err != nil {
		if err == pgx.ErrNoRows {
			return nil, LoginError
		}
		return nil, err
	}
	var user models.User
	user.ID = id
	user.Role = role
	return &user, nil
}
