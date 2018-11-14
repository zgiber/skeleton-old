package service

import (
	"context"
	"errors"

	"github.com/zgiber/skeleton/contract"
)

// Clients service
type Clients struct {
}

// Register ...
func (cli *Clients) Register(ctx context.Context, req *contract.RegisterClientRequest) (*contract.RegisterClientResponse, error) {
	return nil, errors.New("not implemented")
}

// ValidateID ...
func (cli *Clients) ValidateID(ctx context.Context, req *contract.ValidateIDRequest) (*contract.ValidateIDResponse, error) {
	return nil, errors.New("not implemented")
}

// Login ...
func (cli *Clients) Login(ctx context.Context, req *contract.LoginRequest) (*contract.LoginResponse, error) {
	return nil, errors.New("not implemented")
}

// Authenticate ...
func (cli *Clients) Authenticate(ctx context.Context, req *contract.AuthenticationRequest) (*contract.AuthenticationResponse, error) {
	return nil, errors.New("not implemented")
}

// Auth service
type Auth struct {
}

// Authorize ...
func (auth *Auth) Authorize(ctx context.Context, req *contract.AuthorizationRequest) (*contract.AuthorizationResponse, error) {
	return nil, errors.New("not implemented")
}
