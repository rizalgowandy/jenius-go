package api

import (
	"context"
	"errors"

	"github.com/go-resty/resty/v2"
	"github.com/rizalgowandy/jenius-go/pkg/entity"
)

func NewAuthentication(cfg Config) *Authentication {
	return &Authentication{
		client: NewRestyClient(cfg),
		id:     cfg.ClientID,
		secret: cfg.ClientSecret,
	}
}

type Authentication struct {
	client *resty.Client
	id     string
	secret string
}

func (a *Authentication) OAuth2(ctx context.Context) (*entity.OAuth2Resp, error) {
	url := "/api/oauth/token"

	var content entity.OAuth2Resp

	resp, err := a.client.R().
		SetContext(ctx).
		SetHeaders(map[string]string{
			"Content-Type":  "application/x-www-form-urlencoded",
			"client_id":     a.id,
			"client_secret": a.secret,
			"grant_type":    "client_credentials",
			"scope":         "resource.WRITE resource.READ",
		}).
		SetResult(&content).
		Post(url)
	if err != nil {
		return nil, err
	}

	respHeader := resp.Header()
	respCode := respHeader.Get("X-ResponseCode")
	if respCode != "200" {
		respDesc := respHeader.Get("X-ResponseDesc")
		return nil, errors.New(respDesc)
	}

	return &content, nil
}
