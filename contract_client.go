package apihttp

import (
	"context"
	"fmt"
	"time"

	apis "github.com/proximax-storage/go-xpx-dfms-api"
	"github.com/proximax-storage/go-xpx-dfms-drive"
)

type apiContractClient apiHttp

func (api *apiContractClient) Compose(ctx context.Context, space uint64, subPeriod time.Duration, opts ...apis.ComposeOpt) (*drive.Contract, error) {
	options, err := apis.Parse(space, subPeriod, opts...)
	if err != nil {
		return nil, err
	}

	if options.PrivateKey != nil {
		return nil, fmt.Errorf("can't use custom PrivateKey in HTTP API") // for security reasons.
	}

	out := new(contractResponse)
	return out.Contract, api.apiHttp().NewRequest("contract/compose").
		Arguments(fmt.Sprintf("%d", space)).
		Arguments(subPeriod.String()).
		Option("replicas", options.Replicas).
		Option("min-replicators", options.MinReplicators).
		Option("subscription-price", options.SubscriptionPrice).
		Option("number-subscription-periods", options.NumberSubscriptionPeriods).
		Option("percent-approvers", options.PercentApprovers).
		Exec(ctx, out)
}

func (api *apiContractClient) List(ctx context.Context) ([]drive.ID, error) {
	out := new(contractLsResponse)
	return out.Ids, api.apiHttp().NewRequest("contract/ls").Exec(ctx, out)
}

func (api *apiContractClient) Get(ctx context.Context, id drive.ID) (*drive.Contract, error) {
	out := new(contractResponse)
	return out.Contract, api.apiHttp().NewRequest("contract/get").Arguments(id.String()).Exec(ctx, out)
}

func (api *apiContractClient) Amendments(ctx context.Context, id drive.ID) (apis.ContractSubscription, error) {
	resp, err := api.apiHttp().NewRequest("contract/ammends").Arguments(id.String()).Send(ctx)
	if err != nil {
		return nil, err
	}

	return newContractSub(ctx, resp.Output), nil
}

func (api *apiContractClient) Finish(ctx context.Context, id drive.ID) error {
	return api.apiHttp().NewRequest("contract/finish").Arguments(id.String()).Exec(ctx, nil)
}

func (api *apiContractClient) Verify(ctx context.Context, id drive.ID) (apis.VerifyResult, error) {
	out := new(verifyResponse)
	return out.Result, api.apiHttp().NewRequest("contract/verify").Arguments(id.String()).Exec(ctx, out)
}

func (api *apiContractClient) apiHttp() *apiHttp {
	return (*apiHttp)(api)
}

type inviteResponse struct {
	Invite *drive.Invite
}

type contractResponse struct {
	Contract *drive.Contract
}

type contractLsResponse struct {
	Ids []drive.ID
}

type verifyResponse struct {
	Result apis.VerifyResult
}
