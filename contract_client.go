package apihttp

import (
	"context"

	"github.com/ipfs/go-cid"
	"github.com/proximax-storage/go-xpx-dfms-api"
	"github.com/proximax-storage/go-xpx-dfms-drive"
)

type apiContractClient apiHttp

func (api *apiContractClient) Compose(ctx context.Context, space, duration uint64, opts ...api.ComposeOpt) (drive.Contract, error) {
	out := new(drive.BasicContract)
	return out, api.apiHttp().NewRequest("contract/compose").
		Arguments(string(space)).
		Arguments(string(duration)).
		Exec(ctx, out)
}

func (api *apiContractClient) List(ctx context.Context) ([]drive.ID, error) {
	out := new(contractLsResponse)
	return out.Cids, api.apiHttp().NewRequest("contract/ls").Exec(ctx, out)
}

func (api *apiContractClient) Get(ctx context.Context, id drive.ID) (drive.Contract, error) {
	out := new(drive.BasicContract)
	return out, api.apiHttp().NewRequest("contract/get").Exec(ctx, out)
}

func (api *apiContractClient) Amendments(ctx context.Context, id drive.ID) (drive.ContractSubscription, error) {
	resp, err := api.apiHttp().NewRequest("contract/ammends").Send(ctx)
	if err != nil {
		return nil, err
	}

	return newContractSub(ctx, resp.Output), nil
}

func (api *apiContractClient) apiHttp() *apiHttp {
	return (*apiHttp)(api)
}

type contractLsResponse struct {
	Cids []cid.Cid
}
