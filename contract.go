package httpclient

import (
	"context"
	"time"

	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p-core/peer"
)

type ContractAPI Client

type DriveInvite struct {
	Cid           cid.Cid
	Created       time.Time
	Duration      time.Duration
	RequiredSpace uint64
	Owner         peer.ID
}

type ContractResponse struct {
	Cid        cid.Cid
	Owner      peer.ID
	Members    []peer.ID
	Duration   time.Duration
	Created    time.Time
	Root       cid.Cid
	TotalSpace uint64
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

func (api *ContractAPI) Updates(ctx context.Context, id cid.Cid) (*ContractResponse, error) {
	out := &ContractResponse{}
	err := api.client().Request("contract/updates").
		Arguments(id.String()).
		Exec(ctx, out)
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

func (api *ContractAPI) Invites(ctx context.Context) (*DriveInvite, error) {
	out := &DriveInvite{}
	err := api.client().Request("contract/invites").
		Exec(ctx, out)
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

type getContractResponse struct {
	Cid string
}

type listContractResponse struct {
	Cids []cid.Cid
}
