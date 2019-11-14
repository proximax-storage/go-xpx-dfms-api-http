package apihttp

import (
	"archive/tar"
	"context"
	"io"
	"os"
	"path/filepath"

	files "github.com/ipfs/go-ipfs-files"
	drive "github.com/proximax-storage/go-xpx-dfms-drive"
)

func (api *apiDriveFS) newNode(ctx context.Context, ctr drive.ID, path string, info os.FileInfo) (files.Node, error) {
	if info.IsDir() {
		return api.newDir(ctx, ctr, path, info.Size())
	}

	return api.newFile(ctx, ctr, path)
}

func (api *apiDriveFS) newFile(ctx context.Context, ctr drive.ID, path string) (files.File, error) {
	driveId, err := drive.IDToString(ctr)
	if err != nil {
		return nil, err
	}

	resp, err := api.apiHttp().NewRequest("drive/get").
		Arguments(driveId).
		Arguments(path).
		Send(ctx)
	if err != nil {
		return nil, err
	}

	r := tar.NewReader(resp.Output)
	header, err := r.Next()
	if err != nil && err != io.EOF {
		return nil, err
	}

	return files.NewReaderStatFile(r, header.FileInfo()), nil
}

func (api *apiDriveFS) newDir(ctx context.Context, ctr drive.ID, path string, size int64) (files.Directory, error) {
	infos, err := api.Ls(ctx, ctr, path)
	if err != nil {
		return nil, err
	}

	return &dir{
		ctx:   ctx,
		ctr:   ctr,
		infos: infos,
		path:  path,
		size:  size,
		api:   api,
	}, nil
}

type dir struct {
	ctx context.Context

	ctr   drive.ID
	infos []os.FileInfo
	path  string
	size  int64

	api *apiDriveFS
}

func (d *dir) Entries() files.DirIterator {
	return &dirIter{
		dir: d,
		i:   -1,
	}
}

func (d *dir) Size() (int64, error) {
	return d.size, nil
}

func (d *dir) Close() error {
	return nil
}

type dirIter struct {
	*dir

	i    int
	node files.Node
	err  error
}

func (di *dirIter) Name() string {
	if di.i == -1 {
		return ""
	}

	return di.current().Name()
}

func (di *dirIter) Node() files.Node {
	return di.node
}

func (di *dirIter) Err() error {
	return di.err
}

func (di *dirIter) Next() bool {
	if di.ctx.Err() != nil {
		di.err = di.ctx.Err()
		return false
	}

	l := len(di.infos)
	if l == 0 || di.i == l-1 {
		return false
	}

	di.i++

	node, err := di.api.newNode(di.ctx, di.ctr, filepath.Join(di.path, di.Name()), di.current())
	if err != nil {
		di.err = err
		return false
	}

	di.node = node
	return true
}

func (di *dirIter) current() os.FileInfo {
	return di.infos[di.i]
}
