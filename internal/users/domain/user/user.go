package user

import "errors"

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

func (u *User) UpdateDisplayname(displayname string) error {
	if displayname == "" {
		return errors.New("displayname cannot be empty")
	}

	u.Displayname = displayname
	return nil
}

func (u *User) UpdatePassword(password string) error {
	if password == "" {
		return errors.New("password cannot be empty")
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
