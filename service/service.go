package service

import (
	"context"

	"github.com/zgiber/skeleton/contract"
)

// Clients service
type Clients struct {
}

// Register ...
func (cli *Clients) Register(ctx context.Context, req *contract.RegisterClientRequest) (*contract.RegisterClientResponse, error) {
	panic("not implemented")
}

// ValidateID ...
func (cli *Clients) ValidateID(ctx context.Context, req *contract.ValidateIDRequest) (*contract.ValidateIDResponse, error) {
	panic("not implemented")
}

// Login ...
func (cli *Clients) Login(ctx context.Context, req *contract.LoginRequest) (*contract.LoginResponse, error) {
	panic("not implemented")
}

// Authenticate ...
func (cli *Clients) Authenticate(ctx context.Context, req *contract.AuthenticationRequest) (*contract.AuthenticationResponse, error) {
	panic("not implemented")
}

// Auth service
type Auth struct {
}

// Authorize ...
func (auth *Auth) Authorize(ctx context.Context, req *contract.AuthorizationRequest) (*contract.AuthorizationResponse, error) {
	panic("not implemented")
}
