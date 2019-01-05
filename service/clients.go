package service

import (
	"context"

	"github.com/zgiber/skeleton/contract"
)

// Client is a simple CRUD for Client resource types
type Client struct {
}

// Create a new client
func (c *Client) Create(context.Context, *contract.CreateClientRequest) (*contract.CreateClientResponse, error) {
	panic("not implemented")
}

// Retrieve an existing client
func (c *Client) Retrieve(context.Context, *contract.RetrieveClientRequest) (*contract.RetrieveClientResponse, error) {
	panic("not implemented")
}

// Delete an existing client
func (c *Client) Delete(context.Context, *contract.DeleteClientRequest) (*contract.DeleteClientResponse, error) {
	panic("not implemented")
}
