package httpclient

import (
	"os"
	"time"

	"github.com/ipfs/go-cid"
)

type Stat struct {
	Name string
	Size int64
	Type string
	Cid  cid.Cid
}

type FileInfo struct {
	stat Stat
}

func (ref *FileInfo) Name() string {
	return ref.stat.Name
}

func (ref *FileInfo) Size() int64 {
	return ref.stat.Size
}

func (ref *FileInfo) Mode() os.FileMode {
	return 0
}

func (ref *FileInfo) ModTime() time.Time {
	return time.Time{}
}

func (ref *FileInfo) IsDir() bool {
	return ref.stat.Type == "dir"
}

func (ref *FileInfo) Sys() interface{} {
	return &ref.stat
}
