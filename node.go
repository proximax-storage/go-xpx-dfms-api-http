package apihttp

import (
	"context"
	"net/http"

	api "github.com/proximax-storage/go-xpx-dfms-api"
)

type node struct {
	*apiHttp

	token api.AccessToken
	tp    api.NodeType
}

// newNodeAPI creates new *node from given address, http.Client and type
func newNodeAPI(address string, token api.AccessToken, client *http.Client, tp api.NodeType) *node {
	return &node{
		apiHttp: newHTTP(address, token, client),
		token:   token,
		tp:      tp,
	}
}

func (n *node) Network() api.Network {
	return (*apiNetwork)(n.apiHttp)
}

func (n *node) Type() api.NodeType {
	return n.tp
}

func (n *node) Version(ctx context.Context) (api.Version, error) {
	out := api.Version{}
	return out, n.apiHttp.NewRequest("version").Exec(ctx, &out)
}

func (n *node) Auth() api.Auth {
	return (*apiAuth)(n.apiHttp)
}
