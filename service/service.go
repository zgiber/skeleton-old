package service

import (
	"context"

	"github.com/zgiber/skeleton/contract"
)

// ClientService is a simple CRUD for Client resource types
type ClientService struct {
}

// Create a new client
func (c *ClientService) Create(context.Context, *contract.CreateClientRequest) (*contract.CreateClientResponse, error) {
	panic("not implemented")
}

// Retrieve an existing client
func (c *ClientService) Retrieve(context.Context, *contract.RetrieveClientRequest) (*contract.RetrieveClientResponse, error) {
	panic("not implemented")
}

// Delete an existing client
func (c *ClientService) Delete(context.Context, *contract.DeleteClientRequest) (*contract.DeleteClientResponse, error) {
	panic("not implemented")
}

// AuthzService is an authorisation layer
type AuthzService struct {
}

// Authorize the request based on the ResourceID, the Scope and the ClientId
func (a *AuthzService) Authorize(context.Context, *contract.AuthorizationRequest) (*contract.AuthorizationResponse, error) {
	panic("not implemented")
}

// AuthService is an authentication service
type AuthService struct {
}

// Authenticate the request based on the provided token
func (a *AuthService) Authenticate(context.Context, *contract.AuthenticationRequest) (*contract.AuthenticationResponse, error) {
	panic("not implemented")
}

// UserService provides methods for managing users (human clients)
type UserService struct {
}

// Register a user
func (u *UserService) Register(context.Context, *contract.RegisterUserRequest) (*contract.RegisterUserResponse, error) {
	panic("not implemented")
}

// Login with user's credientials
func (u *UserService) Login(context.Context, *contract.LoginRequest) (*contract.LoginResponse, error) {
	panic("not implemented")
}
