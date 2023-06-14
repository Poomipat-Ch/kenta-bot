package user

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	t.Parallel()
	userUUID := uuid.New().String()
	username := "poomipat"
	displayname := "Poomipat Ch"
	email := "poomipat002@gmail.com"
	picture := "https://www.google.com"
	password := "12345678"

	user, err := NewUser(userUUID, username, displayname, email, picture, password)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	assert.Equal(t, userUUID, user.UUID)
	assert.Equal(t, username, user.Username)
	assert.Equal(t, displayname, user.Displayname)
	assert.Equal(t, email, user.Email)
	assert.Equal(t, picture, user.Picture)
	assert.Equal(t, password, user.Password)
}

func TestNewUser_invalid(t *testing.T) {
	t.Parallel()
	userUUID := uuid.New().String()
	username := "poomipat"
	displayname := "Poomipat Ch"
	email := "poomipat002@gmail.com"
	picture := "https://www.google.com"
	password := "12345678"

	_, err := NewUser("", username, displayname, email, picture, password)
	assert.Error(t, err)

	_, err = NewUser(userUUID, "", displayname, email, picture, password)
	assert.Error(t, err)

	_, err = NewUser(userUUID, username, displayname, "", picture, password)
	assert.Error(t, err)

	_, err = NewUser(userUUID, username, displayname, email, picture, "")
	assert.Error(t, err)
}

func TestUser_UpdateDisplayname(t *testing.T) {
	t.Parallel()
	user := newValidUser()

	newDisplayname := "Poomipat Chaiyapong"
	err := user.UpdateDisplayname(newDisplayname)
	assert.NoError(t, err)
}

func TestUser_UpdateDisplayname_invalid(t *testing.T) {
	t.Parallel()
	user := newValidUser()

	err := user.UpdateDisplayname("")
	assert.Error(t, err)
}

func TestUser_UpdatePassword(t *testing.T) {
	t.Parallel()
	user := newValidUser()

	newPassword := "87654321"
	err := user.UpdatePassword(newPassword)
	assert.NoError(t, err)
}

func TestUser_UpdatePassword_invalid(t *testing.T) {
	t.Parallel()
	user := newValidUser()

	err := user.UpdatePassword("")
	assert.Error(t, err)
}

func TestUnmarshalUserFromDatabase(t *testing.T) {
	t.Parallel()

	userObj := newValidUser()

	user, err := UnmarshalUserFromDatabase(userObj.UUID, userObj.Username, userObj.Displayname, userObj.Email, userObj.Picture, userObj.Password)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	assert.Equal(t, userObj, user)
}

func TestUnmarshalUserFromDatabase_invalid(t *testing.T) {
	t.Parallel()

	userObj := newValidUser()
	userObj.UUID = ""

	_, err := UnmarshalUserFromDatabase(userObj.UUID, userObj.Username, userObj.Displayname, userObj.Email, userObj.Picture, userObj.Password)

	assert.Error(t, err)
}

func newValidUser() *User {
	userUUID := uuid.New().String()
	username := "poomipat"
	displayname := "Poomipat Ch"
	email := "poomipat002@gmail.com"
	picture := "https://www.google.com"
	password := "12345678"

	user, _ := NewUser(userUUID, username, displayname, email, picture, password)

	return user
}
