package apihttp

import (
	"context"

	drive "github.com/proximax-storage/go-xpx-dfms-drive"
)

type apiLedger apiHttp

func (api *apiLedger) ListDrives(ctx context.Context) ([]drive.ID, error) {
	out := new(ledgerDrsResponse)
	return out.Ids, api.apiHttp().NewRequest("ledger/drives").Exec(ctx, out)
}

func (api *apiLedger) apiHttp() *apiHttp {
	return (*apiHttp)(api)
}

type ledgerDrsResponse struct {
	Ids []drive.ID
}
