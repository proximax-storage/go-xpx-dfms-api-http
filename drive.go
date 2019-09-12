package httpclient

import (
	"context"

	"github.com/ipfs/go-cid"
	files "github.com/ipfs/go-ipfs-files"
)

type DriveAPI Client

func (api *DriveAPI) Add(ctx context.Context, id cid.Cid, path string, file files.Node) (cid.Cid, error) {
	req := api.client().Request("drive/add").
		Arguments(id.String()).
		Arguments(path)

	d := files.NewMapDirectory(map[string]files.Node{"": file}) // unwrapped on the other side inside commands
	req.Body(files.NewMultiFileReader(d, false))

	resp, err := req.Send(ctx)
	if err != nil {
		return cid.Undef, err
	}
	if resp.Error != nil {
		return cid.Undef, resp.Error
	}
	defer resp.Output.Close()

	var out addResponse
	err = resp.decode(&out)
	if err != nil {
		return cid.Undef, err
	}

	id, err = cid.Decode(out.Cid)
	if err != nil {
		return cid.Undef, err
	}

	return id, nil
}

func (api *DriveAPI) client() *Client {
	return (*Client)(api)
}

type addResponse struct {
	Cid string
}
