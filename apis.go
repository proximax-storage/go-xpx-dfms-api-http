package apihttp

import (
	"net/http"

	api "github.com/proximax-storage/go-xpx-dfms-api"
)

type clientAPI struct {
	*node
}

// NewNewClientAPI creates new api.Client from given address and default http.Client
func NewClientAPI(address string) api.Client {
	return NewCustomClientAPI(address, http.DefaultClient)
}

// NewCustomClientAPI creates new api.Client from given address and custom http.Client
func NewCustomClientAPI(address string, client *http.Client) api.Client {
	return &clientAPI{newNodeAPI(address, client, api.ClientType)}
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
	return NewCustomReplicatorAPI(address, http.DefaultClient)
}

// NewCustomReplicatorAPI creates new api.Replicator from given address and custom http.Client
func NewCustomReplicatorAPI(address string, client *http.Client) api.Replicator {
	return &replicatorAPI{newNodeAPI(address, client, api.ReplicatorType)}
}

func (r *replicatorAPI) Contract() api.ContractReplicator {
	return (*apiContractReplicator)(r.node.apiHttp)
}
