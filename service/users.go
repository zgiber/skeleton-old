package service

import (
	"context"

	"github.com/zgiber/skeleton/contract"
)

// User provides methods for managing users (human clients)
type User struct {
}

// Register a user
func (u *User) Register(context.Context, *contract.RegisterUserRequest) (*contract.RegisterUserResponse, error) {
	panic("not implemented")
}

// Login with user's credientials
func (u *User) Login(context.Context, *contract.LoginRequest) (*contract.LoginResponse, error) {
	panic("not implemented")
}
