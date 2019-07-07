package httpclient

import (
	"context"
	"encoding/json"
	"time"

	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p-core/peer"
)

// ContractAPI basic Contract Client representation
type ContractAPI Client

// InviteSubscription interface represent basic func
// for handle subscription data
type InviteSubscription interface {
	Next() (InviteResponse, error)
	Cancel() error
}

// UpdatesSubscription interface represent basic func
// for handle subscription data
type UpdatesSubscription interface {
	Next() (UpdatesResponse, error)
	Cancel() error
}

// subscriptionResponse state for handle subscriptions
type subscriptionResponse struct {
	ctx    context.Context
	reader *Response
}

// UpdatesResponse with subscription funcs
type InviteResponse struct {
	subscriptionResponse
	Cid           cid.Cid
	Created       time.Time
	Duration      time.Duration
	RequiredSpace uint64
	Owner         peer.ID
}

// UpdatesResponse with subscription funcs
type UpdatesResponse struct {
	subscriptionResponse
	ContractResponse
}

// ContractResponse common contract data struct
type ContractResponse struct {
	Cid        cid.Cid
	Owner      peer.ID
	Members    []peer.ID
	Duration   time.Duration
	Created    time.Time
	Root       cid.Cid
	TotalSpace uint64
}

// listContractResponse list contract response data
type listContractResponse struct {
	Cids []cid.Cid
}

// Get Contract request
func (api *ContractAPI) Get(ctx context.Context, id cid.Cid) (*ContractResponse, error) {
	out := &ContractResponse{}
	err := api.client().Request("contract/get").
		Arguments(id.String()).
		Exec(ctx, out)
	return out, err
}

// Join Contract request
func (api *ContractAPI) Join(ctx context.Context, id cid.Cid) error {
	return api.client().Request("contract/join").
		Arguments(id.String()).
		Exec(ctx, nil)
}

// List Contracts request
func (api *ContractAPI) List(ctx context.Context) ([]cid.Cid, error) {
	out := &listContractResponse{}
	err := api.client().Request("contract/list").
		Exec(ctx, out)
	return out.Cids, err
}

// Updates subscription Contract request listener
func (api *ContractAPI) Updates(ctx context.Context, id cid.Cid) (UpdatesSubscription, error) {
	out := UpdatesResponse{}
	resp, err := api.client().Request("contract/updates").
		Arguments(id.String()).
		Send(ctx)
	if err != nil {
		return out, err
	}
	out.ctx = ctx
	out.reader = resp
	return out, err
}

// Invites subscription Contract request listener
func (api *ContractAPI) Invites(ctx context.Context) (InviteSubscription, error) {
	out := InviteResponse{}
	resp, err := api.client().Request("contract/invites").
		Send(ctx)
	if err != nil {
		return out, err
	}
	out.ctx = ctx
	out.reader = resp
	return out, err
}

// Compose Contract request
func (api *ContractAPI) Compose(ctx context.Context, space uint64, duration time.Duration) (*ContractResponse, error) {
	out := &ContractResponse{}
	err := api.client().Request("contract/compose").
		Arguments(string(space)).
		Arguments(string(duration)).
		Exec(ctx, out)
	return out, err
}

// StartAccepting Contract request
func (api *ContractAPI) StartAccepting(ctx context.Context) error {
	return api.client().Request("contract/start-accepting").
		Exec(ctx, nil)
}

// StopAccepting Contract request
func (api *ContractAPI) StopAccepting(ctx context.Context) error {
	return api.client().Request("contract/stop-accepting").
		Exec(ctx, nil)
}

// client - init Contract Client
func (api *ContractAPI) client() *Client {
	return (*Client)(api)
}

// Next subscription event
func (di InviteResponse) Next() (InviteResponse, error) {
	err := json.NewDecoder(di.reader.Output).Decode(di)
	return di, err
}

// Cancel subscription listening
func (di InviteResponse) Cancel() error {
	return di.reader.Cancel()
}

// Next subscription event
func (cr UpdatesResponse) Next() (UpdatesResponse, error) {
	err := json.NewDecoder(cr.reader.Output).Decode(cr)
	return cr, err
}

// Cancel subscription listening
func (cr UpdatesResponse) Cancel() error {
	return cr.reader.Cancel()
}
