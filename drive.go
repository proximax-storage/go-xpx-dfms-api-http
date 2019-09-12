package httpclient

import (
	"context"
	"os"

	"github.com/ipfs/go-cid"
	files "github.com/ipfs/go-ipfs-files"
)

type DriveAPI Client

func (api *DriveAPI) Add(ctx context.Context, id cid.Cid, path string, file files.Node) (cid.Cid, error) {
	req := api.client().NewRequest("drive/add").
		Arguments(id.String()).
		Arguments(path).
		FileBody(file)

	out := &addResponse{}
	err := req.Exec(ctx, out)
	if err != nil {
		return cid.Undef, err
	}

	return cid.Decode(out.Cid)
}

func (api *DriveAPI) Get(ctx context.Context, id cid.Cid, path string) (files.Node, error) {
	info, err := api.Stat(ctx, id, path)
	if err != nil {
		return nil, err
	}

	return api.newNode(ctx, id, path, info)
}

func (api *DriveAPI) Remove(ctx context.Context, id cid.Cid, path string) error {
	return api.client().NewRequest("drive/remove").
		Arguments(id.String()).
		Arguments(path).
		Exec(ctx, nil)
}

func (api *DriveAPI) Move(ctx context.Context, id cid.Cid, src string, dst string) error {
	return api.client().NewRequest("drive/mv").
		Arguments(id.String()).
		Arguments(src).
		Arguments(dst).
		Exec(ctx, nil)
}

func (api *DriveAPI) Copy(ctx context.Context, id cid.Cid, src string, dst string) error {
	return api.client().NewRequest("drive/cp").
		Arguments(id.String()).
		Arguments(src).
		Arguments(dst).
		Exec(ctx, nil)
}

func (api *DriveAPI) MakeDir(ctx context.Context, id cid.Cid, path string) error {
	return api.client().NewRequest("drive/mkdir").
		Arguments(id.String()).
		Arguments(path).
		Exec(ctx, nil)
}

func (api *DriveAPI) Ls(ctx context.Context, id cid.Cid, path string) ([]os.FileInfo, error) {
	out := &lsResponse{}
	err := api.client().NewRequest("drive/ls").
		Arguments(id.String()).
		Arguments(path).
		Exec(ctx, out)
	return out.toFileInfo(), err
}

func (api *DriveAPI) Stat(ctx context.Context, id cid.Cid, path string) (os.FileInfo, error) {
	out := &statResponse{}
	err := api.client().NewRequest("drive/stat").
		Arguments(id.String()).
		Arguments(path).
		Exec(ctx, out)
	return out.toFileInfo(), err
}

func (api *DriveAPI) Flush(ctx context.Context, id cid.Cid, path string) error {
	return api.client().NewRequest("drive/flush").
		Arguments(id.String()).
		Arguments(path).
		Exec(ctx, nil)
}

func (api *DriveAPI) client() *Client {
	return (*Client)(api)
}

type addResponse struct {
	Cid string
}

type lsResponse struct {
	List []*Stat
}

func (ls *lsResponse) toFileInfo() []os.FileInfo {
	fis := make([]os.FileInfo, len(ls.List))
	for i, entry := range ls.List {
		fis[i] = &FileInfo{*entry}
	}

	return fis
}

type statResponse struct {
	Stat *Stat
}

func (ls *statResponse) toFileInfo() os.FileInfo {
	return &FileInfo{*ls.Stat}
}
