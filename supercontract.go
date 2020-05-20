package apihttp

import (
	"context"
	"fmt"

	"github.com/ipfs/go-cid"
	"github.com/proximax-storage/go-xpx-dfms-drive"
	sc "github.com/proximax-storage/go-xpx-dfms-drive/supercontract"
)

type apiSuperContract apiHttp

type scDeployResponse struct {
	Result sc.ID
}

type scExecuteResponse struct {
	ScId   string
	TxHash cid.Cid
}

type scContractLsResponse struct {
	Ids []sc.ID
}

type scContractResultsResponse struct {
	Results []string
}

type scExecutionsResponse struct {
	Ids []cid.Cid
}

type scContractResponse struct {
	SuperContract *sc.SuperContract
}

func (api *apiSuperContract) Deploy(ctx context.Context, id drive.ID, file string) (sc.ID, error) {
	out := new(scDeployResponse)
	return out.Result, api.apiHttp().
		NewRequest("sc/deploy").
		Arguments(id.String()).
		Arguments(file).
		Exec(ctx, out)
}

func (api *apiSuperContract) Execute(ctx context.Context, id sc.ID, gas uint64, function sc.Function) (cid.Cid, error) {
	var funcParams []string
	for _, param := range function.Parameters {
		funcParams = append(funcParams, fmt.Sprintf("%d", param))
	}

	out := new(scExecuteResponse)
	return out.TxHash, api.apiHttp().
		NewRequest("sc/exec").
		Arguments(id.String()).
		Arguments(fmt.Sprintf("%d", gas)).
		Arguments(function.Name).
		Arguments(funcParams...).
		Exec(ctx, out)
}

func (api *apiSuperContract) Get(ctx context.Context, id sc.ID) (*sc.SuperContract, error) {
	out := new(scContractResponse)
	return out.SuperContract, api.apiHttp().
		NewRequest("sc/get").
		Arguments(id.String()).
		Exec(ctx, out)
}

func (api *apiSuperContract) List(ctx context.Context, id drive.ID) ([]sc.ID, error) {
	out := new(scContractLsResponse)
	return out.Ids, api.apiHttp().
		NewRequest("sc/ls").
		Arguments(id.String()).
		Exec(ctx, out)
}

func (api *apiSuperContract) GetResults(ctx context.Context, id cid.Cid) ([]string, error) {
	out := new(scContractResultsResponse)
	return out.Results, api.apiHttp().
		NewRequest("sc/results").
		Arguments(id.String()).
		Exec(ctx, out)
}

func (api *apiSuperContract) GetSuperContractExecutionsHashes(ctx context.Context) ([]cid.Cid, error) {
	out := new(scExecutionsResponse)
	return out.Ids, api.apiHttp().
		NewRequest("sc/executions").
		Exec(ctx, out)
}

func (api *apiSuperContract) Deactivate(ctx context.Context, id sc.ID) error {
	return api.apiHttp().
		NewRequest("sc/deactivate").
		Arguments(id.String()).
		Exec(ctx, nil)
}

func (api *apiSuperContract) apiHttp() *apiHttp {
	return (*apiHttp)(api)
}
