package apihttp

import (
	"context"
	"time"

	api "github.com/proximax-storage/go-xpx-dfms-api"
)

type apiAuth apiHttp

func (a *apiAuth) GenerateToken(ctx context.Context, id string, duration time.Duration, s ...string) (api.AccessToken, error) {
	out := &generateResponse{}
	return out.Token, a.apiHttp().NewRequest("auth/generate").
		Arguments(id).
		Arguments(duration.String()).
		Arguments(s...).
		Exec(ctx, out)
}

func (a *apiAuth) VerifyAccessTo(ctx context.Context, to string, token api.AccessToken) error {
	return a.apiHttp().NewRequest("auth/verify").
		Arguments(to).
		Arguments(string(token)).
		Exec(ctx, nil)
}

func (a *apiAuth) List(ctx context.Context) ([]string, error) {
	out := &listResponse{}
	return out.TokenIDs, a.apiHttp().NewRequest("auth/ls").
		Exec(ctx, out)
}

func (a *apiAuth) RevokeToken(ctx context.Context, tokenID string) error {
	return a.apiHttp().NewRequest("auth/revoke").
		Arguments(tokenID).
		Exec(ctx, nil)
}

func (a *apiAuth) apiHttp() *apiHttp {
	return (*apiHttp)(a)
}

type generateResponse struct {
	Token api.AccessToken
}

type listResponse struct {
	TokenIDs []string
}
