package apihttp

import (
	"context"
	"time"

	"github.com/proximax-storage/go-xpx-dfms-api"
	"github.com/proximax-storage/go-xpx-dfms-drive"
)

type apiContractReplicator apiHttp

func (api *apiContractReplicator) Compose(ctx context.Context, space uint64, subPeriod time.Duration, opts ...api.ComposeOpt) (*drive.Contract, error) {
	return api.apiContractClient().Compose(ctx, space, subPeriod, opts...)
}

func (api *apiContractReplicator) List(ctx context.Context) ([]drive.ID, error) {
	return api.apiContractClient().List(ctx)
}

func (api *apiContractReplicator) Get(ctx context.Context, id drive.ID) (*drive.Contract, error) {
	return api.apiContractClient().Get(ctx, id)
}

func (api *apiContractReplicator) Amendments(ctx context.Context, id drive.ID) (api.ContractSubscription, error) {
	return api.apiContractClient().Amendments(ctx, id)
}

func (api *apiContractReplicator) Accept(ctx context.Context, id drive.ID) error {
	return api.apiHttp().NewRequest("contract/accept").Arguments(id.String()).Exec(ctx, nil)
}

func (api *apiContractReplicator) Accepted(ctx context.Context) (api.ContractSubscription, error) {
	resp, err := api.apiHttp().NewRequest("contract/accepted").Send(ctx)
	if err != nil {
		return nil, err
	}

	return newContractSub(ctx, resp.Output), nil
}

func (api *apiContractReplicator) Finish(ctx context.Context, id drive.ID) error {
	return api.apiContractClient().Finish(ctx, id)
}

func (api *apiContractReplicator) Verify(ctx context.Context, id drive.ID) (api.VerifyResult, error) {
	return api.apiContractClient().Verify(ctx, id)
}

func (api *apiContractReplicator) Invites(ctx context.Context) (api.InviteSubscription, error) {
	resp, err := api.apiHttp().NewRequest("contract/invites").Send(ctx)
	if err != nil {
		return nil, err
	}

	return newInviteSub(ctx, resp.Output), nil
}

func (api *apiContractReplicator) apiContractClient() *apiContractClient {
	return (*apiContractClient)(api)
}

func (api *apiContractReplicator) apiHttp() *apiHttp {
	return (*apiHttp)(api)
}
