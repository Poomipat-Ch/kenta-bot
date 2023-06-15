package adapters

import (
	"context"
	"testing"

	"github.com/Poomipat-Ch/kenta-bot/internal/users/domain/user"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var repo user.Repository
var initUsers []*user.User

var test_uuid1 = uuid.New().String()

func TestMain(m *testing.M) {
	repo = NewUsersMemoryRepository()

	initUsers = []*user.User{
		{
			UUID:        test_uuid1,
			Username:    "username1",
			Email:       "email1",
			Password:    "password1",
			Displayname: "displayname1",
		},
		{
			UUID:        "uuid2",
			Username:    "username2",
			Email:       "email2",
			Password:    "password2",
			Displayname: "displayname2",
		},
		{
			UUID:        "uuid3",
			Username:    "username3",
			Email:       "email3",
			Password:    "password3",
			Displayname: "displayname3",
		},
	}

	for _, user := range initUsers {
		repo.Store(context.Background(), user)
	}

	m.Run()
}

func TestFindByEmail(t *testing.T) {
	// arrange
	email := "email1"

	// act
	u, err := repo.FindByEmail(context.Background(), email)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, test_uuid1, u.UUID)
}

func TestFindByEmail_NotFoundErr(t *testing.T) {
	// arrange
	email := "not found"

	// act
	_, err := repo.FindByEmail(context.Background(), email)

	// assert
	assert.EqualError(t, err, user.NotFoundError{UserIdentity: email}.Error())
}

func TestFindByUUID(t *testing.T) {
	// act
	u, err := repo.FindByUUID(context.Background(), test_uuid1)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, test_uuid1, u.UUID)
}

func TestFindByUUID_NotFoundErr(t *testing.T) {
	// arrange
	uuid := "not found"

	// act
	_, err := repo.FindByUUID(context.Background(), uuid)

	// assert
	assert.EqualError(t, err, user.NotFoundError{UserIdentity: uuid}.Error())
}

func TestStore(t *testing.T) {
	// arrange
	u, _ := user.NewUser(uuid.New().String(), "username4", "displayname4", "email4", "picture4", "password4")

	// act
	err := repo.Store(context.Background(), u)

	// assert
	assert.NoError(t, err)
}

func TestStore_AlreadyExistsErr(t *testing.T) {
	// arrange
	u, _ := user.NewUser(test_uuid1, "username1", "displayname1", "email1", "picture1", "password1")

	// act
	err := repo.Store(context.Background(), u)

	// assert
	assert.EqualError(t, err, user.AlreadyExistsError{UserIdentity: u.UUID}.Error())
}

func TestUpdate(t *testing.T) {
	// arrange
	u, _ := user.NewUser(test_uuid1, "username1", "displayname1", "email1", "picture1", "password1")

	// act
	err := repo.Update(context.Background(), u.UUID, u)

	// assert
	assert.NoError(t, err)
}

func TestUpdate_NotFoundErr(t *testing.T) {
	// arrange
	u, _ := user.NewUser("not found", "username4", "displayname4", "email4", "picture4", "password4")

	// act
	err := repo.Update(context.Background(), u.UUID, u)

	// assert
	assert.EqualError(t, err, user.NotFoundError{UserIdentity: u.UUID}.Error())
}

// bracnh test

func BenchmarkFindByEmail(b *testing.B) {
	for i := 0; i < b.N; i++ {
		repo.FindByEmail(context.Background(), "email1")
	}
}
