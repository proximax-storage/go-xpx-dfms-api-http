package apihttp

import (
	"context"
	"os"

	"github.com/ipfs/go-cid"

	files "github.com/ipfs/go-ipfs-files"
	iapi "github.com/proximax-storage/go-xpx-dfms-api"
	drive "github.com/proximax-storage/go-xpx-dfms-drive"
)

// apiDriveFS is an implementation of DFMS's drive apiHttp.
// Drive is FS like abstraction, which allows Drive contract
// owners to interact with disks of contract members owner occupies.
// All FS manipulations are executed locally on the DFMS node to
// prepare the Drive and all it's data for further transportation and replication.
// That means clientAPI's Drive would make a full copy of the files.
// Files can be cleared from DFMS node disk after uploading to members.
// To upload local Drive state to members use Flush.
type apiDriveFS apiHttp

// Add adds the file to a specific Drive and to a given path.
func (api *apiDriveFS) Add(ctx context.Context, id drive.ID, path string, file files.Node, opts ...iapi.DriveOption) (cid.Cid, error) {
	opt := iapi.ParseDriveOptions(opts...)

	req := api.apiHttp().NewRequest("drive/add").
		Arguments(id.String()).
		Arguments(path).
		Option("flush", opt.Flush).
		FileBody(file)

	out := &addResponse{}
	err := req.Exec(ctx, out)
	if err != nil {
		return cid.Undef, err
	}
	return cid.Decode(out.Id)
}

// Get retrieves file from a specific contract at a given path.
func (api *apiDriveFS) Get(ctx context.Context, id drive.ID, path string, opts ...iapi.DriveOption) (files.Node, error) {
	info, err := api.Stat(ctx, id, path)
	if err != nil {
		return nil, err
	}

	return api.newNode(ctx, id, path, info)
}

// FIXME Currently, directories are not supported and the name is empty.
// File gets the file from the Drive by given cid.
func (api *apiDriveFS) File(ctx context.Context, id drive.ID, cid cid.Cid, opts ...iapi.DriveOption) (files.Node, error) {
	resp, err := api.apiHttp().NewRequest("drive/file").
		Arguments(id.String()).
		Arguments(cid.String()).
		Send(ctx)
	if err != nil {
		return nil, err
	}

	return fileFromResp(resp)
}

// Remove removes reference of the file at a given path from a specific Drive.
// If the file is cached and no more references left it clears the cache.
func (api *apiDriveFS) Remove(ctx context.Context, id drive.ID, path string, opts ...iapi.DriveOption) error {
	opt := iapi.ParseDriveOptions(opts...)
	return api.apiHttp().NewRequest("drive/rm").
		Arguments(id.String()).
		Arguments(path).
		Option("flush", opt.Flush).
		Exec(ctx, nil)
}

// Move moves file in a specific Drive from one path to another.
func (api *apiDriveFS) Move(ctx context.Context, id drive.ID, src string, dst string, opts ...iapi.DriveOption) error {
	opt := iapi.ParseDriveOptions(opts...)
	return api.apiHttp().NewRequest("drive/mv").
		Arguments(id.String()).
		Arguments(src).
		Arguments(dst).
		Option("flush", opt.Flush).
		Exec(ctx, nil)
}

// Copy copies file in a specific Drive from one path to another.
// NOTE: Does not do actual copy, only copies reference(a.k.a SymLink).
// That way file is not duplicated on a disk, but accesible from different paths.
func (api *apiDriveFS) Copy(ctx context.Context, id drive.ID, src string, dst string, opts ...iapi.DriveOption) error {
	opt := iapi.ParseDriveOptions(opts...)
	return api.apiHttp().NewRequest("drive/cp").
		Arguments(id.String()).
		Arguments(src).
		Arguments(dst).
		Option("flush", opt.Flush).
		Exec(ctx, nil)
}

// MakeDir created new directory in a specific Drive at a given path.
func (api *apiDriveFS) MakeDir(ctx context.Context, id drive.ID, path string, opts ...iapi.DriveOption) error {
	opt := iapi.ParseDriveOptions(opts...)
	return api.apiHttp().NewRequest("drive/mkdir").
		Arguments(id.String()).
		Arguments(path).
		Option("flush", opt.Flush).
		Exec(ctx, nil)
}

// Ls lists all the files and directories of a specific Drive and information about them at a given path.
func (api *apiDriveFS) Ls(ctx context.Context, id drive.ID, path string) ([]os.FileInfo, error) {
	out := &lsResponse{}
	err := api.apiHttp().NewRequest("drive/ls").
		Arguments(id.String()).
		Arguments(path).
		Exec(ctx, out)
	if err != nil {
		return nil, err
	}

	return out.toFileInfo(), nil
}

// Stat returns information about a file or directory at a given path of a specific Drive
func (api *apiDriveFS) Stat(ctx context.Context, id drive.ID, path string) (os.FileInfo, error) {
	out := &statResponse{&Stat{}}
	return out.toFileInfo(), api.apiHttp().NewRequest("drive/stat").
		Arguments(id.String()).
		Arguments(path).
		Exec(ctx, out)
}

// Flush uploads the state of a Drive to all contract members.
func (api *apiDriveFS) Flush(ctx context.Context, id drive.ID) error {
	return api.apiHttp().NewRequest("drive/flush").
		Arguments(id.String()).
		Exec(ctx, nil)
}

// Clear clears all files locally
func (api *apiDriveFS) Clear(ctx context.Context, opts ...iapi.DriveOption) error {
	return api.apiHttp().NewRequest("drive/clear").
		Exec(ctx, nil)
}

func (api *apiDriveFS) apiHttp() *apiHttp {
	return (*apiHttp)(api)
}

type addResponse struct {
	Id string
}

type lsResponse struct {
	List []*Stat
}

func (ls *lsResponse) toFileInfo() []os.FileInfo {
	fis := make([]os.FileInfo, len(ls.List))
	for i, entry := range ls.List {
		fis[i] = &FileInfo{entry}
	}

	return fis
}

type statResponse struct {
	Stat *Stat
}

func (ls *statResponse) toFileInfo() os.FileInfo {
	return &FileInfo{ls.Stat}
}
