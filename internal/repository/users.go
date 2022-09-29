package repository

type UserInterface interface {
	UserExists()
	PasswordMatches()
	ValidateEmail()
	ValidatePassword()
	CreateUser()
}
