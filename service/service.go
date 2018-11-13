package service

import (
	"context"

	"github.com/zgiber/skeleton/contract"
)

// Clients service
type Clients struct {
}

// Register ...
func (usr *Clients) Register(context.Context, *contract.RegisterClientRequest) (*contract.RegisterClientResponse, error) {
	panic("not implemented")
}

// ValidateID ...
func (usr *Clients) ValidateID(context.Context, *contract.ValidateIDRequest) (*contract.ValidateIDResponse, error) {
	panic("not implemented")
}

// Login ...
func (usr *Clients) Login(context.Context, *contract.LoginRequest) (*contract.LoginResponse, error) {
	panic("not implemented")
}

// Authenticate ...
func (usr *Clients) Authenticate(context.Context, *contract.AuthenticationRequest) (*contract.AuthenticationResponse, error) {
	panic("not implemented")
}

// Auth service
type Auth struct {
}

// Authorize ...
func (auth *Auth) Authorize(context.Context, *contract.AuthorizationRequest) (*contract.AuthorizationResponse, error) {
	panic("not implemented")
}
