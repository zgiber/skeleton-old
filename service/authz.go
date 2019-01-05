package service

import (
	"context"

	"github.com/zgiber/skeleton/contract"
)

// Authz is an authorisation layer
type Authz struct {
}

// Authorize the request based on the ResourceID, the Scope and the ClientId
func (a *Authz) Authorize(context.Context, *contract.AuthorizationRequest) (*contract.AuthorizationResponse, error) {
	panic("not implemented")
}
