package jenius

import (
	"github.com/rizalgowandy/jenius-go/pkg/api"
)

// NewClient creates a client to interact with Jenius API.
func NewClient(cfg api.Config) (*Client, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return &Client{
		auth: api.NewAuthentication(cfg),
	}, nil
}

// Client is the main client to interact with Jenius API.
type Client struct {
	auth *api.Authentication
}
