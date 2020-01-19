package apihttp

import (
	"context"

	"github.com/proximax-storage/go-xpx-dfms-drive"
)

type apiSuperContract apiHttp

func (api *apiSuperContract) Deploy(ctx context.Context, id drive.ID, file string) error {
	return api.apiHttp().NewRequest("supercontract/deploy").
		Arguments(id.String()).
		Arguments(file).
		Exec(ctx, nil)
}

func (api *apiSuperContract) apiHttp() *apiHttp {
	return (*apiHttp)(api)
}
