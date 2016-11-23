package fs

import (
	"io"
	"os"
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
	Stat() (os.FileInfo, error)
}
