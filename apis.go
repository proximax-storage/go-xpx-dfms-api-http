package apihttp

import (
	"net/http"

	api "github.com/proximax-storage/go-xpx-dfms-api"
)

type clientAPI struct {
	*node
}

// NewClientAPI creates new api.Client from given address and default http.Client
func NewClientAPI(address string) api.Client {
	return NewCustomClientAPI(address, nil, http.DefaultClient)
}

// NewClientWithToken creates new api.Client from given address, token and default http.Client
func NewClientAPIWithToken(address string, token api.AccessToken) api.Client {
	return NewCustomClientAPI(address, token, http.DefaultClient)
}

// NewCustomClientAPI creates new api.Client from given address and custom http.Client
func NewCustomClientAPI(address string, token api.AccessToken, client *http.Client) api.Client {
	return &clientAPI{newNodeAPI(address, token, client, api.ClientType)}
}

func (c *clientAPI) Contract() api.ContractClient {
	return (*apiContractClient)(c.node.apiHttp)
}

func (c *clientAPI) FS() api.DriveFS {
	return (*apiDriveFS)(c.node.apiHttp)
}

func (c *clientAPI) SuperContract() api.SuperContract {
	return (*apiSuperContract)(c.node.apiHttp)
}

type replicatorAPI struct {
	*node
}

// NewReplicatorAPI creates new api.Replicator from given address and default http.Client
func NewReplicatorAPI(address string) api.Replicator {
	return NewCustomReplicatorAPI(address, nil, http.DefaultClient)
}

// NewReplicatorAPI creates new api.Replicator from given address and default http.Client
func NewReplicatorAPIWithToken(address string, token api.AccessToken) api.Replicator {
	return NewCustomReplicatorAPI(address, token, http.DefaultClient)
}

// NewCustomReplicatorAPI creates new api.Replicator from given address and custom http.Client
func NewCustomReplicatorAPI(address string, token api.AccessToken, client *http.Client) api.Replicator {
	return &replicatorAPI{newNodeAPI(address, token, client, api.ReplicatorType)}
}

func (r *replicatorAPI) Contract() api.ContractReplicator {
	return (*apiContractReplicator)(r.node.apiHttp)
}
