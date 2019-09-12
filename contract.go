package httpclient

import (
	"context"
	"encoding/json"
	"io"
	"time"

	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p-core/peer"
)

type ContractAPI Client

type InviteSubscription interface {
	Next() (DriveInvite, error)
	//	Cancel()
}

type UpdateSubscription interface {
	Next() (ContractResponse, error)
	//	Cancel()
}

type DriveInvite struct {
	ctx    context.Context
	reader io.ReadCloser

	Cid           cid.Cid
	Created       time.Time
	Duration      time.Duration
	RequiredSpace uint64
	Owner         peer.ID
}

type ContractResponse struct {
	ctx    context.Context
	reader io.ReadCloser

	Cid        cid.Cid
	Owner      peer.ID
	Members    []peer.ID
	Duration   time.Duration
	Created    time.Time
	Root       cid.Cid
	TotalSpace uint64
}

type listContractResponse struct {
	Cids []cid.Cid
}

func (api *ContractAPI) Get(ctx context.Context, id cid.Cid) (*ContractResponse, error) {
	out := &ContractResponse{}
	err := api.client().Request("contract/get").
		Arguments(id.String()).
		Exec(ctx, out)
	return out, err
}

func (api *ContractAPI) Join(ctx context.Context, id cid.Cid) error {
	return api.client().Request("contract/join").
		Arguments(id.String()).
		Exec(ctx, nil)
}

func (api *ContractAPI) List(ctx context.Context) ([]cid.Cid, error) {
	out := &listContractResponse{}
	err := api.client().Request("contract/list").
		Exec(ctx, out)
	return out.Cids, err
}

func (api *ContractAPI) Updates(ctx context.Context, id cid.Cid) (UpdateSubscription, error) {
	out := ContractResponse{}
	resp, err := api.client().Request("contract/updates").
		Arguments(id.String()).
		Send(ctx)
	if err != nil {
		return out, err
	}
	out.reader = resp.Output
	return out, err
}

func (api *ContractAPI) Compose(ctx context.Context, space uint64, duration time.Duration) (*ContractResponse, error) {
	out := &ContractResponse{}
	err := api.client().Request("contract/compose").
		Arguments(string(space)).
		Arguments(string(duration)).
		Exec(ctx, out)
	return out, err
}

func (api *ContractAPI) Invites(ctx context.Context) (InviteSubscription, error) {
	out := DriveInvite{}
	resp, err := api.client().Request("contract/invites").
		Send(ctx)
	if err != nil {
		return out, err
	}
	out.reader = resp.Output
	return out, err
}

func (api *ContractAPI) StartAccepting(ctx context.Context) error {
	return api.client().Request("contract/start-accepting").
		Exec(ctx, nil)
}

func (api *ContractAPI) StopAccepting(ctx context.Context) error {
	return api.client().Request("contract/stop-accepting").
		Exec(ctx, nil)
}

func (api *ContractAPI) client() *Client {
	return (*Client)(api)
}

func (di DriveInvite) Next() (DriveInvite, error) {
	err := json.NewDecoder(di.reader).Decode(di)
	return di, err
}

/*
func (di DriveInvite) Cancel() {

}
*/
func (cr ContractResponse) Next() (ContractResponse, error) {
	err := json.NewDecoder(cr.reader).Decode(cr)
	return cr, err
}

/*
func (cr ContractResponse) Cancel() {

}
*/
