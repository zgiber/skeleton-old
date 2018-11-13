package service

import (
	"context"

	"github.com/zgiber/skeleton/contract"
)

// Users service
type Users struct {
}

// Register ...
func (usr *Users) Register(context.Context, *contract.RegisterUserRequest) (*contract.RegisterUserResponse, error) {
	panic("not implemented")
}

// ValidateID ...
func (usr *Users) ValidateID(context.Context, *contract.ValidateIDRequest) (*contract.ValidateIDResponse, error) {
	panic("not implemented")
}

// Login ...
func (usr *Users) Login(context.Context, *contract.LoginRequest) (*contract.LoginResponse, error) {
	panic("not implemented")
}

// Authenticate ...
func (usr *Users) Authenticate(context.Context, *contract.AuthenticationRequest) (*contract.AuthenticationResponse, error) {
	panic("not implemented")
}

// Auth service
type Auth struct {
}

// Authorize ...
func (auth *Auth) Authorize(context.Context, *contract.AuthorizationRequest) (*contract.AuthorizationResponse, error) {
	panic("not implemented")
}
