package service

import "errors"

var (
	ErrorAddUser               = errors.New("error when adding a user")
	ErrorReceivingUser         = errors.New("error receiving the user")
	ErrorCheckUser             = errors.New("error checking the user")
	ErrorNotFoundUser          = errors.New("the user has been not found")
	ErrorDeleteUser            = errors.New("error when deleting a user")
	ErrorRegistrationUser      = errors.New("error registration a user")
	ErrorSignedToken           = errors.New("error when signed token")
	ErrorUsernameAlreadyExists = errors.New("such a user already exists")
	ErrorCreatedToken          = errors.New("token creation error")
	ErrorDuplicateUsername     = errors.New("a user with such username already exists")
	ErrorDuplicateEmail        = errors.New("a user with such email already exists")
	ErrorEmptyEmail            = errors.New("the email field cannot be empty")
	ErrorEmptyUsername         = errors.New("the username field cannot be empty")
)
