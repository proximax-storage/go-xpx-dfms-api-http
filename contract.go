package client

import (
	"context"
	"encoding/json"
	"time"

	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p-core/peer"
)

// ContractAPI implements DFMS's ContractAPI
// Drive contracts are contracts between network peers for physical disk space.
// The API allows to create, join, invite, listen for contract updates and invitations.
// After creating a Drive contract an owner will be able to use disk space of members through DriveAPI.
// NOTE: Currently creates Drive contracts with DFMS(R) identity as owner.
type ContractAPI Client

// InviteSubscription is a subscription for new invitations
// to join Drive contract.
type InviteSubscription interface {
	// Next blocks till new invite is received
	Next() (InviteResponse, error)

	// Cancel stops subscription
	Cancel() error
}

// UpdatesSubscription is a subscription for all updates
// of the specific Drive contract
type UpdatesSubscription interface {
	// Next blocks till new update is received
	Next() (UpdatesResponse, error)

	// Cancel stops subscription
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

// Get retrieves Contract by it's content id.
func (api *ContractAPI) Get(ctx context.Context, id cid.Cid) (*ContractResponse, error) {
	out := &ContractResponse{}
	err := api.client().NewRequest("contract/get").
		Arguments(id.String()).
		Exec(ctx, out)
	return out, err
}

// Join accepts Drive contract invitation by content id.
// Note: Operation is non revertible, once join Node must follow the contract.
func (api *ContractAPI) Join(ctx context.Context, id cid.Cid) error {
	return api.client().NewRequest("contract/join").
		Arguments(id.String()).
		Exec(ctx, nil)
}

// List Drive contracts the node has relation with.
func (api *ContractAPI) List(ctx context.Context) ([]cid.Cid, error) {
	out := &listContractResponse{}
	err := api.client().NewRequest("contract/list").
		Exec(ctx, out)
	return out.Cids, err
}

// Updates creates new subscription for updates of specific Drive contract
func (api *ContractAPI) Updates(ctx context.Context, id cid.Cid) (UpdatesSubscription, error) {
	out := UpdatesResponse{}
	resp, err := api.client().NewRequest("contract/updates").
		Arguments(id.String()).
		Send(ctx)
	if err != nil {
		return out, err
	}
	out.ctx = ctx
	out.reader = resp
	return out, err
}

// Invites create new subscription for Drive contract invitation from the network
func (api *ContractAPI) Invites(ctx context.Context) (InviteSubscription, error) {
	out := InviteResponse{}
	resp, err := api.client().NewRequest("contract/invites").
		Send(ctx)
	if err != nil {
		return out, err
	}
	out.ctx = ctx
	out.reader = resp
	return out, err
}

// Compose constructs new Drive contract from specific parameters.
// It announces invitation on the network and waits till minimum(2) amount of nodes Join the invitation.
// Compose does not guarantee successful completion and may error if minimum(2) amount of nodes have not found through timeout.
func (api *ContractAPI) Compose(ctx context.Context, space uint64, duration time.Duration) (*ContractResponse, error) {
	out := &ContractResponse{}
	err := api.client().NewRequest("contract/compose").
		Arguments(string(space)).
		Arguments(string(duration)).
		Exec(ctx, out)
	return out, err
}

// StartAccepting triggers node to automatically accept incoming contracts.
func (api *ContractAPI) StartAccepting(ctx context.Context) error {
	return api.client().NewRequest("contract/start-accepting").
		Exec(ctx, nil)
}

// StopAccepting stops accepting process.
func (api *ContractAPI) StopAccepting(ctx context.Context) error {
	return api.client().NewRequest("contract/stop-accepting").
		Exec(ctx, nil)
}

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
