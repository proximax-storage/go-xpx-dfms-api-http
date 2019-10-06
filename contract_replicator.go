package apihttp

import (
	"context"

	"github.com/ipfs/go-cid"
	"github.com/proximax-storage/go-xpx-dfms-api"
	"github.com/proximax-storage/go-xpx-dfms-drive"
)

type apiContractReplicator apiHttp

func (api *apiContractReplicator) Compose(ctx context.Context, space, duration uint64, opts ...api.ComposeOpt) (drive.Contract, error) {
	return api.apiContractClient().Compose(ctx, space, duration, opts...)
}

func (api *apiContractReplicator) List(ctx context.Context) ([]cid.Cid, error) {
	return api.apiContractClient().List(ctx)
}

func (api *apiContractReplicator) Get(ctx context.Context, id drive.ID) (drive.Contract, error) {
	return api.apiContractClient().Get(ctx, id)
}

func (api *apiContractReplicator) Amendments(ctx context.Context, id drive.ID) (drive.ContractSubscription, error) {
	return api.apiContractClient().Amendments(ctx, id)
}

func (api *apiContractReplicator) Accept(ctx context.Context, id drive.ID) error {
	return api.apiHttp().NewRequest("contract/accept").Arguments(id.String()).Exec(ctx, nil)
}

func (api *apiContractReplicator) Accepted(context.Context) (drive.ContractSubscription, error) {
	panic("not implemented")
}

func (api *apiContractReplicator) Invites(ctx context.Context) (drive.InviteSubscription, error) {
	resp, err := api.apiHttp().NewRequest("contract/invites").Send(ctx)
	if err != nil {
		return nil, err
	}

	return newInviteSub(ctx, resp.Output), nil
}

func (api *apiContractReplicator) StartAccepting(ctx context.Context, _ api.AcceptStrategy) error {
	return api.apiHttp().NewRequest("contract/accept").Exec(ctx, nil)
}

func (api *apiContractReplicator) StopAccepting(ctx context.Context) error {
	return api.apiHttp().NewRequest("contract/accept").Exec(ctx, nil)
}

func (api *apiContractReplicator) apiContractClient() *apiContractClient {
	return (*apiContractClient)(api)
}

func (api *apiContractReplicator) apiHttp() *apiHttp {
	return (*apiHttp)(api)
}
