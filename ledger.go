package apihttp

import (
	"context"
	"encoding/json"
	"fmt"

	pcrypto "github.com/libp2p/go-libp2p-core/crypto"
	xcrypto "github.com/proximax-storage/go-xpx-crypto"
	drive "github.com/proximax-storage/go-xpx-dfms-drive"
)

type apiLedger apiHttp

func (api *apiLedger) ListDrives(ctx context.Context) ([]drive.ID, error) {
	r, err := api.apiHttp().http.Get(fmt.Sprintf("%s/drives", api.address))
	if err != nil {
		return nil, err
	}

	ldl := &ledgerDlsPage{}
	err = json.NewDecoder(r.Body).Decode(ldl)
	if err != nil {
		return nil, err
	}

	ids := make([]drive.ID, len(ldl.Drives))
	for i, d := range ldl.Drives {

		xpkey, err := xcrypto.NewPublicKeyfromHex(d.Drive.DriveKey)
		if err != nil {
			return nil, err
		}

		pkey, err := pcrypto.UnmarshalEd25519PublicKey(xpkey.Raw)
		if err != nil {
			return nil, err
		}

		ids[i], err = drive.IDFromPubKey(pkey)
		if err != nil {
			return nil, err
		}
	}

	return ids, nil
}

func (api *apiLedger) apiHttp() *apiHttp {
	return (*apiHttp)(api)
}

type ledgerDlsPage struct {
	Drives []ledgerDrive `json:"data"`
}

type ledgerDrive struct {
	Drive struct {
		DriveKey string `json:"multisig"`
	}
}
