package domain

import (
	"fmt"
	"regexp"

	core_errors "github.com/ilez0660/Golangclicker/Server/internal/core/errors"
)

type User struct {
	ID      int
	Version int

	FullName string
	Email    *string
}

func NewUser(
	id int,
	version int,
	fullName string,
	email *string,
) User {
	return User{
		ID:       id,
		Version:  version,
		FullName: fullName,
		Email:    email,
	}
}

func NewUserUninitialized(
	fullName string,
	email *string,
) User {
	return NewUser(
		UninitializedID,
		UninitializedVersion,
		fullName,
		email,
	)
}

func (u *User) Validate() error{
	fullNameLength := len([]rune(u.FullName))
	if fullNameLength < 3 || fullNameLength > 100{
		return fmt.Errorf(
			"invalid `FullName` len: %d: %w",
			fullNameLength,
			core_errors.ErrInvalidArgument,
		)
	}
	if u.Email != nil{
		emailLen := len([]rune(*u.Email))
		if emailLen < 10 || emailLen > 50{
			return fmt.Errorf(
				"invalid `Email` len: %d: %w",
				emailLen,
				core_errors.ErrInvalidArgument,
			)
		}
		re := regexp.MustCompile(`^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$`)

		if !re.MatchString(*u.Email){
			return fmt.Errorf(
				"invalid `Email` format: %w",
				core_errors.ErrInvalidArgument,
			)
		}
	}
	return nil
}
