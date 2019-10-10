package apihttp

import (
	"context"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/multiformats/go-multiaddr"
)

type apiNetwork apiHttp

func (api *apiNetwork) Connect(ctx context.Context, addrs ...multiaddr.Multiaddr) error {
	req := api.apiHttp().NewRequest("net/connect")
	for _, addr := range addrs {
		req.Arguments(addr.String())
	}

	return req.Exec(ctx, nil)
}

func (api *apiNetwork) Disconnect(ctx context.Context, addrs ...multiaddr.Multiaddr) error {
	req := api.apiHttp().NewRequest("net/disconnect")
	for _, addr := range addrs {
		req.Arguments(addr.String())
	}

	return req.Exec(ctx, nil)
}

func (api *apiNetwork) Peers(ctx context.Context) ([]*peer.AddrInfo, error) {
	out := new(peersResponse)
	return out.Peers, api.apiHttp().NewRequest("net/peers").Exec(ctx, &out)
}

func (api *apiNetwork) ID(ctx context.Context) (peer.ID, error) {
	out := new(idResponse)
	return out.ID, api.apiHttp().NewRequest("net/id").Exec(ctx, &out)
}

func (api *apiNetwork) Addrs(ctx context.Context) ([]multiaddr.Multiaddr, error) {
	out := new(addrsResponse)

	err := api.apiHttp().NewRequest("net/addrs").Exec(ctx, &out)
	if err != nil {
		return nil, err
	}

	addrs := make([]multiaddr.Multiaddr, len(out.Addrs))
	for i, addr := range out.Addrs {
		addrs[i], err = multiaddr.NewMultiaddr(addr)
		if err != nil {
			return nil, err
		}
	}

	return addrs, nil
}

func (api *apiNetwork) apiHttp() *apiHttp {
	return (*apiHttp)(api)
}

type peersResponse struct {
	Peers []*peer.AddrInfo
}

type idResponse struct {
	ID peer.ID
}

type addrsResponse struct {
	Addrs []string
}
