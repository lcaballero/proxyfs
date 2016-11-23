package fs

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// CachedFileSystem implements the FileSystem interface.
type CachedFileSystem struct {
	root  string
	files map[string]File
}

// NewCachedFileSystem creates a FileSystem over the given directory, which
// must be an absolute path.  If the directory provided is not absolute
// an error.
func NewCachedFileSystem(root string) (FileSystem, error) {
	if !filepath.IsAbs(root) {
		return nil, ErrPathNotAbsolute
	}
	cf := &CachedFileSystem{
		root:  root,
		files: make(map[string]File),
	}
	return cf, nil
}

// Open maps the given name to a file rooted at the mount point of the
// CachedFileSystem returning a File if it exists or an error if one
// occurs while reading the file.
func (c *CachedFileSystem) Open(name string) (File, error) {
	path := filepath.Join(c.root, name)
	if !filepath.IsAbs(path) {
		return nil, ErrPathNotAbsolute
	}

	f, ok := c.files[path]
	if ok {
		return f, nil
	}

	rawFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer rawFile.Close()

	info, err := rawFile.Stat()
	if err != nil {
		return nil, err
	}

	bin, err := ioutil.ReadAll(rawFile)
	if err != nil {
		return nil, err
	}

	cf := &CachedFile{
		path: path,
		bin:  bin,
		info: info,
	}
	c.files[path] = cf
	return cf, nil
}
