package adapters

import (
	"context"
	"sync"

	userModel "github.com/Poomipat-Ch/kenta-bot/internal/users/domain/user"
)

type usersMemoryRepository struct {
	users map[string]userModel.User
	sync.RWMutex
}

// FindByEmail implements user.Repository.
func (u *usersMemoryRepository) FindByEmail(ctx context.Context, email string) (*userModel.User, error) {
	u.RLock()
	defer u.RUnlock()
	for _, user := range u.users {
		if user.Email == email {
			return &user, nil
		}
	}

	return nil, userModel.NotFoundError{UserIdentity: email}
}

// FindByUUID implements user.Repository.
func (u *usersMemoryRepository) FindByUUID(ctx context.Context, uuid string) (*userModel.User, error) {
	u.RLock()
	defer u.RUnlock()
	if user, ok := u.users[uuid]; ok {
		return &user, nil
	}

	return nil, userModel.NotFoundError{UserIdentity: uuid}
}

// FindByUsername implements user.Repository.
func (u *usersMemoryRepository) FindByUsername(ctx context.Context, username string) (*userModel.User, error) {
	u.RLock()
	defer u.RUnlock()
	for _, user := range u.users {
		if user.Username == username {
			return &user, nil
		}
	}

	return nil, userModel.NotFoundError{UserIdentity: username}
}

// Store implements user.Repository.
func (u *usersMemoryRepository) Store(ctx context.Context, user *userModel.User) error {

	if u.users == nil {
		u.Lock()
		u.users = make(map[string]userModel.User)
		u.Unlock()
	}

	u.RLock()
	if _, ok := u.users[user.UUID]; ok {
		u.RUnlock()
		return userModel.AlreadyExistsError{UserIdentity: user.UUID}
	}
	u.RUnlock()

	u.Lock()
	u.users[user.UUID] = *user
	u.Unlock()
	return nil
}

// Update implements user.Repository.
func (u *usersMemoryRepository) Update(ctx context.Context, uuid string, user *userModel.User) error {

	u.RLock()
	if _, ok := u.users[uuid]; !ok {
		u.RUnlock()
		return userModel.NotFoundError{UserIdentity: uuid}
	}
	u.RUnlock()

	u.Lock()
	u.users[uuid] = *user
	u.Unlock()
	return nil
}

// Delete implement user.Repository
func (u *usersMemoryRepository) Delete(ctx context.Context, uuid string) error {

	u.Lock()
	defer u.Unlock()
	if _, ok := u.users[uuid]; !ok {
		return userModel.NotFoundError{UserIdentity: uuid}
	}

	delete(u.users, uuid)

	return nil
}

func NewUsersMemoryRepository() userModel.Repository {
	return &usersMemoryRepository{
		users: make(map[string]userModel.User),
	}
}
