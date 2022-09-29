package service

import "tax_me/internal/datastruct"

type UserService interface {
	LogIn(user datastruct.User) error
	SignIn(user datastruct.User) error
}
