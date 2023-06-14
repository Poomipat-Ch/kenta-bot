package user

import (
	commonerrors "github.com/Poomipat-Ch/kenta-bot/internal/common/errors"
	"github.com/pkg/errors"
)

type User struct {
	UUID        string
	Username    string
	Displayname string
	Email       string
	Picture     string
	Password    string
}

func NewUser(uuid, username, displayname, email, picture, password string) (*User, error) {

	if uuid == "" {
		return nil, errors.New("uuid cannot be empty")
	}

	if username == "" {
		return nil, errors.New("username cannot be empty")
	}

	if email == "" {
		return nil, errors.New("email cannot be empty")
	}

	if password == "" {
		return nil, errors.New("password cannot be empty")
	}

	return &User{
		UUID:        uuid,
		Username:    username,
		Displayname: displayname,
		Email:       email,
		Picture:     picture,
		Password:    password,
	}, nil
}

var ErrDisplaynameCannotBeEmpty = commonerrors.NewInvalidInputError("displayname cannot be empty", "displayname-cannot-be-empty")

func (u *User) UpdateDisplayname(displayname string) error {
	if displayname == "" {
		return errors.WithStack(ErrDisplaynameCannotBeEmpty)
	}

	u.Displayname = displayname
	return nil
}

var ErrPasswordCannotBeEmpty = commonerrors.NewInvalidInputError("password cannot be empty", "password-cannot-be-empty")

func (u *User) UpdatePassword(password string) error {
	if password == "" {
		return errors.WithStack(ErrPasswordCannotBeEmpty)
	}

	u.Password = password
	return nil
}

func UnmarshalUserFromDatabase(
	uuid, username, displayname, email, picture, password string,
) (*User, error) {
	user, err := NewUser(uuid, username, displayname, email, picture, password)

	if err != nil {
		return nil, err
	}

	return user, nil
}
