package user

import "errors"

type User struct {
	uuid        string
	username    string
	displayname string
	email       string
	picture     string
	password    string
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
		uuid:        uuid,
		username:    username,
		displayname: displayname,
		email:       email,
		picture:     picture,
		password:    password,
	}, nil
}

func (u *User) UpdateDisplayname(displayname string) error {
	if displayname == "" {
		return errors.New("displayname cannot be empty")
	}

	u.displayname = displayname
	return nil
}

func (u *User) UpdatePassword(password string) error {
	if password == "" {
		return errors.New("password cannot be empty")
	}

	u.password = password
	return nil
}
