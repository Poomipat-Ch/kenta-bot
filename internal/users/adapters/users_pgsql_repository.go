package adapters

import (
	"context"
	"database/sql"
	"time"

	"github.com/Poomipat-Ch/kenta-bot/internal/users/domain/user"
)

type userPgsqlRepository struct {
	pgsql *sql.DB
}

// FindByEmail implements user.Repository.
func (u *userPgsqlRepository) FindByEmail(ctx context.Context, email string) (*user.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	row := u.pgsql.QueryRowContext(ctx, "SELECT * FROM users WHERE email = $1", email)

	var uuid, username, displayname, picture, password string

	err := row.Scan(&uuid, &username, &displayname, &email, &picture, &password)

	if err != nil {
		return nil, user.NotFoundError{UserIdentity: uuid}
	}

	return user.UnmarshalUserFromDatabase(uuid, username, displayname, email, picture, password)
}

// FindByUUID implements user.Repository.
func (u *userPgsqlRepository) FindByUUID(ctx context.Context, uuid string) (*user.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	row := u.pgsql.QueryRowContext(ctx, "SELECT * FROM users WHERE uuid = $1", uuid)

	var username, displayname, email, picture, password string

	err := row.Scan(&uuid, &username, &displayname, &email, &picture, &password)

	if err != nil {
		return nil, user.NotFoundError{UserIdentity: uuid}
	}

	return user.UnmarshalUserFromDatabase(uuid, username, displayname, email, picture, password)
}

// FindByUsername implements user.Repository.
func (u *userPgsqlRepository) FindByUsername(ctx context.Context, username string) (*user.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	row := u.pgsql.QueryRowContext(ctx, "SELECT * FROM users WHERE username = $1", username)

	var uuid, displayname, email, picture, password string

	err := row.Scan(&uuid, &username, &displayname, &email, &picture, &password)

	if err != nil {
		return nil, user.NotFoundError{UserIdentity: uuid}
	}

	return user.UnmarshalUserFromDatabase(uuid, username, displayname, email, picture, password)
}

// Store implements user.Repository.
func (u *userPgsqlRepository) Store(ctx context.Context, user *user.User) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	tx, err := u.pgsql.Begin()

	if err != nil {
		return err
	}

	_, err = u.pgsql.ExecContext(ctx, "INSERT INTO users (uuid, username, displayname, email, picture, password) VALUES ($1, $2, $3, $4, $5, $6)", user.UUID, user.Username, user.Displayname, user.Email, user.Picture, user.Password)

	if err := tx.Commit(); err != nil {
		return err
	}

	if err != nil {
		return err
	}

	return nil
}

// Update implements user.Repository.
func (u *userPgsqlRepository) Update(ctx context.Context, uuid string, user *user.User) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	tx, err := u.pgsql.Begin()

	if err != nil {
		return err
	}

	_, err = u.pgsql.ExecContext(ctx, "UPDATE users SET username = $1, displayname = $2, email = $3, picture = $4, password = $5 WHERE uuid = $6", user.Username, user.Displayname, user.Email, user.Picture, user.Password, uuid)

	if err := tx.Commit(); err != nil {
		return err
	}

	if err != nil {
		return err
	}

	return nil
}

func NewUserPgsqlRepository(pgsql *sql.DB) user.Repository {
	return &userPgsqlRepository{
		pgsql: pgsql,
	}
}
