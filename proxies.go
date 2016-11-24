package proxyfs

import (
	"io"
	"os"
	"time"
)

// FileSystem provides the Open method to obtain File instances that
// abstract a file system, much like os.Open provides FileInfo structs.
type FileSystem interface {
	Open(name string) (File, error)
}

// File is an abstraction over actual Files on the system provided by
// a FileSystem implementation.
type File interface {
	io.Closer
	io.Reader
	io.ReaderAt
	io.Seeker
	Stat() (FileInfo, error)
}

// FileInfo interface provided to match the os.FileInfo object.
type FileInfo interface {
	Name() string
	Size() int64
	Mode() os.FileMode
	ModTime() time.Time
	IsDir() bool
	Sys() interface{}
}

// OpenProvider is a function signature for opening files to be read from
// the file system.
type OpenProvider func(string) ([]byte, error)

// InfoProvider is a function signature for obtaining FileInfo instances.
type InfoProvider func(string) (os.FileInfo, error)
