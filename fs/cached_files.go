package fs

import (
	"errors"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

// AssetPathNotAbsoluteError occurs when a new FileSystem isn't rooted
// via an absolute path.
var AssetPathNotAbsoluteError = errors.New("asset mount not absolute path error")

// OffsetBeyondBoundsError occurs when attempting to read outside the
// bounds of the backing file bytes.
var OffsetBeyondBoundsError = errors.New("offset beyond bounds error")

// FileSystem provides the Open method to obtain File instances that
// abstract a file system, much like os.Open provides FileInfo structs.
type FileSystem interface {
	Open(name string) (File, error)
}

// CachedFileSystem implements the FileSystem interface.
type CachedFileSystem struct {
	root  string
	files map[string]File
}

// NewCachedFiles creates a FileSystem over the given directory, which
// must be an absolute path.  If the directory provided is not absolute
// an error.
func NewCachedFiles(root string) (FileSystem, error) {
	if !filepath.IsAbs(root) {
		return CachedFileSystem{}, AssetPathNotAbsoluteError
	}
	cf := CachedFileSystem{
		root:  root,
		files: make(map[string]File),
	}
	return cf, nil
}

// Open maps the given name to a file rooted at the mount point of the
// CachedFileSystem returning a File if it exists or an error if one
// occurs while reading the file.
func (c CachedFileSystem) Open(name string) (File, error) {
	path := filepath.Join(c.root, name)
	if !filepath.IsAbs(path) {
		return CachedFile{}, AssetPathNotAbsoluteError
	}

	f, ok := c.files[path]
	if ok {
		return f, nil
	}

	rawFile, err := os.Open(path)
	if err != nil {
		return CachedFile{}, err
	}
	defer rawFile.Close()

	info, err := rawFile.Stat()
	if err != nil {
		return CachedFile{}, err
	}

	bin, err := ioutil.ReadAll(rawFile)
	if err != nil {
		return CachedFile{}, err
	}

	cf := CachedFile{
		path: path,
		bin:  bin,
		info: info,
	}
	c.files[path] = cf
	return cf, nil
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

// A CachedFile represents
type CachedFile struct {
	path string
	bin  []byte
	info os.FileInfo
}

// Close always returns nil since any error had to of occurred before when
// when Open on the FileSystem was called.
func (c CachedFile) Close() error {
	return nil
}

func (c CachedFile) Read(p []byte) (n int, err error) {
	size := len(c.bin)
	if len(p) <= size {
		src := c.bin[0:len(c.bin)]
		return copy(p, src), nil
	}
	return 0, errors.New("not implemented yet")
}

// ReadAt fills the destination slice with bytes from the file starting
// at the provided offset, reading upto dest length, or the remainder of
// bytes in the file.  The number of bytes read and an error should one
// occur is returned.
func (c CachedFile) ReadAt(p []byte, off int64) (n int, err error) {
	if off >= int64(len(c.bin)) {
		return -1, OffsetBeyondBoundsError
	}
	a, b := int(off), int(len(p))
	src := c.bin[a: a + b]
	return copy(p, src), nil
}


func (c CachedFile) Seek(offset int64, whence int) (int64, error) {
	return int64(0), errors.New("not implemented yet")
}

func (c CachedFile) Stat() (os.FileInfo, error) {
	return c.info, nil
}

// CachedFileInfo provides an implementation of the FileInfo.
type CachedFileInfo struct {
	info os.FileInfo
}

// base name of the file
func (cf CachedFileInfo) Name() string {
	return cf.info.Name()
}

// length in bytes for regular files; system-dependent for others
func (cf CachedFileInfo) Size() int64 {
	return cf.info.Size()
}
// file mode bits
func (cf CachedFileInfo) Mode() os.FileMode {
	return cf.info.Mode()
}

// modification time
func (cf CachedFileInfo) ModTime() time.Time {
	return cf.info.ModTime()
}

// abbreviation for Mode().IsDir()
func (cf CachedFileInfo) IsDir() bool {
	return cf.IsDir()
}

// underlying data source (can return nil)
func (cf CachedFileInfo) Sys() interface{} {
	return cf.Sys()
}


