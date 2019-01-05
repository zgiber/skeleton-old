package service

import (
	"context"

	"github.com/zgiber/skeleton/contract"
)

// Authn is an authentication service
type Authn struct {
}

// Authenticate the request based on the provided token
func (a *Authn) Authenticate(context.Context, *contract.AuthenticationRequest) (*contract.AuthenticationResponse, error) {
	panic("not implemented")
}
