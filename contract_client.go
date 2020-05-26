package apihttp

import (
	"context"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/libp2p/go-libp2p-core/crypto"
	apis "github.com/proximax-storage/go-xpx-dfms-api"
	"github.com/proximax-storage/go-xpx-dfms-drive"
)

type apiContractClient apiHttp

func (api *apiContractClient) Compose(ctx context.Context, space uint64, subPeriod time.Duration, opts ...apis.ComposeOpt) (*drive.Contract, error) {
	options, err := apis.Parse(space, subPeriod, opts...)
	if err != nil {
		return nil, err
	}

	var keyString string
	if options.PrivateKey != nil {
		bytePK, err := crypto.MarshalPrivateKey(options.PrivateKey)
		if err != nil {
			return nil, err
		}

		keyString = hex.EncodeToString(bytePK)
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
		Option("private-key", keyString).
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
