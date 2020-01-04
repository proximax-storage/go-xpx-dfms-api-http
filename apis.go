package apihttp

import (
	"net/http"

	api "github.com/proximax-storage/go-xpx-dfms-api"
)

type clientAPI struct {
	*apiHttp
}

// New creates new apiHttp.Client from given address and default http.Client
func NewClientAPI(address string) api.Client {
	return &clientAPI{newHTTP(address, http.DefaultClient)}
}

// NewCustomClientAPI creates new apiHttp.Client from given address and custom http.Client
func NewCustomClientAPI(address string, client *http.Client) api.Client {
	return &clientAPI{newHTTP(address, client)}
}

func (c *clientAPI) Contract() api.ContractClient {
	return (*apiContractClient)(c.apiHttp)
}

func (c *clientAPI) FS() api.DriveFS {
	return (*apiDriveFS)(c.apiHttp)
}

func (c *clientAPI) Network() api.Network {
	return (*apiNetwork)(c.apiHttp)
}

func (c *clientAPI) Supercontract() api.Supercontract {
	return (*apiSupercontract)(c.apiHttp)
}

type replicatorAPI struct {
	*apiHttp
}

// NewReplicatorAPI creates new api.Client from given address and default http.Client
func NewReplicatorAPI(address string) api.Replicator {
	return &replicatorAPI{newHTTP(address, http.DefaultClient)}
}

// NewCustomReplicatorAPI creates new api.Client from given address and custom http.Client
func NewCustomReplicatorAPI(address string, client *http.Client) api.Replicator {
	return &replicatorAPI{newHTTP(address, client)}
}

func (r *replicatorAPI) Contract() api.ContractReplicator {
	return (*apiContractReplicator)(r.apiHttp)
}

func (r *replicatorAPI) Network() api.Network {
	return (*apiNetwork)(r.apiHttp)
}

func (c *replicatorAPI) Supercontract() api.Supercontract {
	return (*apiSupercontract)(c.apiHttp)
}
