package adapters

// import (
// 	"context"
// 	"testing"

// 	"github.com/Poomipat-Ch/kenta-bot/internal/users/domain/user"
// 	"github.com/stretchr/testify/assert"
// )

// func TestFindByEmail(t *testing.T) {
// 	t.Parallel()

// 	// arrange
// 	repo := NewUserPgsqlRepository()

// 	u, _ := user.NewUser("uuid", "username", "displayname", "email", "picture", "password")

// 	// repo.On("FindByEmail", context.Background(), mock.AnythingOfType("string")).Return(u, nil)

// 	// act
// 	_, err := repo.FindByEmail(context.Background(), "email")

// 	// assert
// 	assert.NoError(t, err)
// }

// func TestFindByEmail_NotFoundErr(t *testing.T) {
// 	t.Parallel()

// 	// arrange
// 	repo := NewUserPgsqlRepository()

// 	// repo.On("FindByEmail", context.Background(), mock.AnythingOfType("string")).Return(nil, user.NotFoundError{UserIdentity: "not found"})

// 	// act
// 	_, err := repo.FindByEmail(context.Background(), "not found")

// 	// assert
// 	assert.EqualError(t, err, user.NotFoundError{UserIdentity: "not found"}.Error())
// }

// func TestFindByUUID(t *testing.T) {
// 	t.Parallel()

// 	// arrange
// 	repo := NewUserPgsqlRepository()
// 	u, _ := user.NewUser("uuid", "username", "displayname", "email", "picture", "password")

// 	// repo.On("FindByUUID", context.Background(), mock.AnythingOfType("string")).Return(u, nil)

// 	// act
// 	_, err := repo.FindByUUID(context.Background(), "uuid")

// 	// assert
// 	assert.NoError(t, err)
// }

// func TestFindByUUID_NotFoundErr(t *testing.T) {
// 	t.Parallel()

// 	// arrange
// 	repo := NewUserPgsqlRepository()

// 	// repo.On("FindByUUID", context.Background(), mock.AnythingOfType("string")).Return(nil, user.NotFoundError{UserIdentity: "not found"})

// 	// act
// 	_, err := repo.FindByUUID(context.Background(), "not found")

// 	// assert
// 	assert.EqualError(t, err, user.NotFoundError{UserIdentity: "not found"}.Error())
// }

// func TestFindByUsername(t *testing.T) {
// 	t.Parallel()

// 	// arrange
// 	repo := NewUserPgsqlRepository()
// 	u, _ := user.NewUser("uuid", "username", "displayname", "email", "picture", "password")

// 	// repo.On("FindByUsername", context.Background(), mock.AnythingOfType("string")).Return(u, nil)

// 	// act
// 	_, err := repo.FindByUsername(context.Background(), "username")

// 	// assert
// 	assert.NoError(t, err)
// }

// func TestFindByUsername_NotFoundErr(t *testing.T) {
// 	t.Parallel()

// 	// arrange
// 	repo := NewUserPgsqlRepository()

// 	// repo.On("FindByUsername", context.Background(), mock.AnythingOfType("string")).Return(nil, user.NotFoundError{UserIdentity: "not found"})

// 	// act
// 	_, err := repo.FindByUsername(context.Background(), "not found")

// 	// assert
// 	assert.EqualError(t, err, user.NotFoundError{UserIdentity: "not found"}.Error())
// }

// func TestStore(t *testing.T) {
// 	t.Parallel()

// 	// arrange
// 	repo := NewUserPgsqlRepository()

// 	u, _ := user.NewUser("uuid", "username", "displayname", "email", "picture", "password")

// 	// repo.On("Store", context.Background(), mock.AnythingOfType("*user.User")).Return(nil)

// 	// act
// 	err := repo.Store(context.Background(), u)

// 	// assert
// 	assert.NoError(t, err)
// }

// func TestUpdate(t *testing.T) {
// 	t.Parallel()

// 	// arrange
// 	repo := NewUserPgsqlRepository()

// 	u, _ := user.NewUser("uuid", "username", "displayname", "email", "picture", "password")

// 	// repo.On("Update", context.Background(), mock.AnythingOfType("string"), mock.AnythingOfType("*user.User")).Return(nil)

// 	// act
// 	err := repo.Update(context.Background(), "uuid", u)

// 	// assert
// 	assert.NoError(t, err)
// }

// func TestUpdate_NotFoundErr(t *testing.T) {
// 	t.Parallel()

// 	// arrange
// 	repo := NewUserPgsqlRepository()

// 	u, _ := user.NewUser("uuid", "username", "displayname", "email", "picture", "password")

// 	// repo.On("Update", context.Background(), mock.AnythingOfType("string"), mock.AnythingOfType("*user.User")).Return(user.NotFoundError{UserIdentity: "not found"})

// 	// act
// 	err := repo.Update(context.Background(), "uuid", u)

// 	// assert
// 	assert.EqualError(t, err, user.NotFoundError{UserIdentity: "not found"}.Error())
// }
