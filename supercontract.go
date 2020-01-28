package apihttp

import (
	"context"
	"fmt"

	"github.com/proximax-storage/go-xpx-dfms-drive"
	sc "github.com/proximax-storage/go-xpx-dfms-drive/supercontract"
)

type apiSuperContract apiHttp

type scContractLsResponse struct {
	Ids []sc.ID
}

type scContractResponse struct {
	Contract *sc.SuperContract
}

func (api *apiSuperContract) Deploy(ctx context.Context, id drive.ID, file string) error {
	return api.apiHttp().NewRequest("supercontract/deploy").
		Arguments(id.String()).
		Arguments(file).
		Exec(ctx, nil)
}

func (api *apiSuperContract) Execute(ctx context.Context, id sc.ID, gas uint64, function string, functionParams []int64) error {
	var funcParams []string
	for _, param := range functionParams {
		funcParams = append(funcParams, fmt.Sprintf("%d", param))
	}

	return api.apiHttp().NewRequest("supercontract/execute").
		Arguments(id.String()).
		Arguments(fmt.Sprintf("%d", gas)).
		Arguments(function).
		Arguments(funcParams...).
		Exec(ctx, nil)
}

func (api *apiSuperContract) GetSuperContract(ctx context.Context, id sc.ID) (*sc.SuperContract, error) {
	out := new(scContractResponse)
	return out.Contract, api.apiHttp().NewRequest("supercontract/get").
		Arguments(id.String()).
		Exec(ctx, out)
}

func (api *apiSuperContract) List(ctx context.Context, id drive.ID) ([]sc.ID, error) {
	out := new(scContractLsResponse)
	return out.Ids, api.apiHttp().NewRequest("supercontract/ls").
		Arguments(id.String()).
		Exec(ctx, out)
}

func (api *apiSuperContract) apiHttp() *apiHttp {
	return (*apiHttp)(api)
}
