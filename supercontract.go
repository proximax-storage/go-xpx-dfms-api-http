package apihttp

import (
	"context"

	"github.com/ipfs/go-cid"

	"github.com/proximax-storage/go-xpx-dfms-drive"
)

type apiSupercontract apiHttp

func (api *apiSupercontract) Deploy(ctx context.Context, id drive.ID, file cid.Cid, functions []string) error {
	return api.apiHttp().NewRequest("supecontract/deploy").
		Arguments(id.String()).
		Arguments(file.String()).
		Arguments(functions...).
		Exec(ctx, nil)
}

func (api *apiSupercontract) apiHttp() *apiHttp {
	return (*apiHttp)(api)
}
