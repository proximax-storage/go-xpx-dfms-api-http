package apihttp

import (
	"context"
	"net/http"

	api "github.com/proximax-storage/go-xpx-dfms-api"
)

type node struct {
	*apiHttp

	tp api.NodeType
}

// newNodeAPI creates new *node from given address, http.Client and type
func newNodeAPI(address string, client *http.Client, tp api.NodeType) *node {
	return &node{
		apiHttp: newHTTP(address, client),
		tp:      tp,
	}
}

func (n *node) Network() api.Network {
	return (*apiNetwork)(n.apiHttp)
}

func (n *node) Type() api.NodeType {
	return n.tp
}

func (n *node) Version(ctx context.Context) (*api.Version, error) {
	out := new(api.Version)
	return out, n.apiHttp.NewRequest("version").Exec(ctx, &out)
}
