package httpclient

import (
	"context"
	"github.com/ipfs/go-cid"
)

type ContractAPI Client

func (api *ContractAPI) Get(ctx context.Context, id cid.Cid) (cid.Cid, error) {
	out := &getContractResponse{}
	err := api.client().Request("contract/get").
		Arguments(id.String()).
		Exec(ctx, out)
	if err != nil {
		return cid.Undef, err
	}
	return cid.Decode(out.Cid)
}

func (api *ContractAPI) Join(ctx context.Context, id cid.Cid) error {
	return api.client().Request("contract/join").
		Arguments(id.String()).
		Exec(ctx, nil)
}

func (api *ContractAPI) List(ctx context.Context, id cid.Cid) ([]cid.Cid, error) {
	out := &listContractResponse{}
	err := api.client().Request("contract/list").
		Arguments(id.String()).
		Exec(ctx, out)
	return out.Cids, err
}

func (api *ContractAPI) Updates(ctx context.Context, id cid.Cid) error {
	return api.client().Request("contract/updates").
		Arguments(id.String()).
		Exec(ctx, nil)
}

func (api *ContractAPI) client() *Client {
	return (*Client)(api)
}

type getContractResponse struct {
	Cid string
}

type listContractResponse struct {
	Cids []cid.Cid
}
