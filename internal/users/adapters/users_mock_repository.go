package adapters

import (
	"context"

	"github.com/Poomipat-Ch/kenta-bot/internal/users/domain/user"
	"github.com/stretchr/testify/mock"
)

type UserMockRepository struct {
	mock.Mock
}

// FindByEmail implements user.Repository.
func (u *UserMockRepository) FindByEmail(ctx context.Context, email string) (*user.User, error) {
	args := u.Called(ctx, email)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*user.User), args.Error(1)
}

// FindByUUID implements user.Repository.
func (u *UserMockRepository) FindByUUID(ctx context.Context, uuid string) (*user.User, error) {
	args := u.Called(ctx, uuid)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*user.User), args.Error(1)
}

// FindByUsername implements user.Repository.
func (u *UserMockRepository) FindByUsername(ctx context.Context, username string) (*user.User, error) {
	args := u.Called(ctx, username)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*user.User), args.Error(1)
}

// Store implements user.Repository.
func (u *UserMockRepository) Store(ctx context.Context, user *user.User) error {
	args := u.Called(ctx, user)

	return args.Error(0)
}

// Update implements user.Repository.
func (u *UserMockRepository) Update(ctx context.Context, uuid string, user *user.User) error {
	args := u.Called(ctx, uuid, user)
	return args.Error(0)
}

func NewUserMockRepository() *UserMockRepository {
	return &UserMockRepository{}
}
