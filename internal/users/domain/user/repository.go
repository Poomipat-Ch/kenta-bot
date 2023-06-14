package user

import (
	"context"
	"fmt"
)

type NotFoundError struct {
	UserUUID string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("User '%s' not found", e.UserUUID)
}

type Repository interface {
	FindByUUID(ctx context.Context, uuid string) (*User, error)
	FindByUsername(ctx context.Context, username string) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	Store(ctx context.Context, user *User) error
	Update(ctx context.Context, uuid string, user User)
}
