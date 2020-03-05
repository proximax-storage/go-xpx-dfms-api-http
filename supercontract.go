package apihttp

import (
	"context"
	"fmt"

	"github.com/ipfs/go-cid"
	"github.com/proximax-storage/go-xpx-dfms-drive"
	sc "github.com/proximax-storage/go-xpx-dfms-drive/supercontract"
)

type apiSuperContract apiHttp

type scContractLsResponse struct {
	Ids []sc.ID
}

type scContractResultsResponse struct {
	Res []string
}

type scExecutionsResponse struct {
	Ids []cid.Cid
}

type scContractResponse struct {
	Contract *sc.SuperContract
}

func (api *apiSuperContract) Deploy(ctx context.Context, id drive.ID, file string) (sc.ID, error) {
	out := new(sc.ID)
	return *out, api.apiHttp().
		NewRequest("supercontract/deploy").
		Arguments(id.String()).
		Arguments(file).
		Exec(ctx, out)
}

func (api *apiSuperContract) Execute(ctx context.Context, id sc.ID, gas uint64, function sc.Function) (cid.Cid, error) {
	var funcParams []string
	for _, param := range function.Parameters {
		funcParams = append(funcParams, fmt.Sprintf("%d", param))
	}

	var out cid.Cid
	return out, api.apiHttp().
		NewRequest("supercontract/execute").
		Arguments(id.String()).
		Arguments(fmt.Sprintf("%d", gas)).
		Arguments(function.Name).
		Arguments(funcParams...).
		Exec(ctx, out)
}

func (api *apiSuperContract) Get(ctx context.Context, id sc.ID) (*sc.SuperContract, error) {
	out := new(scContractResponse)
	return out.Contract, api.apiHttp().
		NewRequest("supercontract/get").
		Arguments(id.String()).
		Exec(ctx, out)
}

func (api *apiSuperContract) List(ctx context.Context, id drive.ID) ([]sc.ID, error) {
	out := new(scContractLsResponse)
	return out.Ids, api.apiHttp().
		NewRequest("supercontract/ls").
		Arguments(id.String()).
		Exec(ctx, out)
}

func (api *apiSuperContract) GetResults(ctx context.Context, id cid.Cid) ([]string, error) {
	out := new(scContractResultsResponse)
	return out.Res, api.apiHttp().
		NewRequest("supercontract/results").
		Arguments(id.String()).
		Exec(ctx, out)
}

func (api *apiSuperContract) GetSuperContractExecutionsHash(ctx context.Context, id sc.ID) ([]cid.Cid, error) {
	out := new(scExecutionsResponse)
	return out.Ids, api.apiHttp().
		NewRequest("supercontract/executions").
		Arguments(id.String()).
		Exec(ctx, out)
}

func (api *apiSuperContract) apiHttp() *apiHttp {
	return (*apiHttp)(api)
}
