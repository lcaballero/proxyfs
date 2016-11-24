package proxyfs

import "errors"

// ErrProviderIsNil occurs when a provider to NewEmbeddedFileSystem is nil
var ErrProviderIsNil = errors.New("provider is nil error")

// EmbeddedFileSystem provides a FileSystem over embedded assets.
type EmbeddedFileSystem struct {
	open  OpenProvider
	info  InfoProvider
	files map[string]File
}

// NewEmbeddedFileSystem creates a new FileSystem with the given functions
// that open and provide file-info.
func NewEmbeddedFileSystem(
	open OpenProvider,
	info InfoProvider) (*EmbeddedFileSystem, error) {

	if open == nil {
		return nil, ErrProviderIsNil
	}
	if info == nil {
		return nil, ErrProviderIsNil
	}

	efs := &EmbeddedFileSystem{
		open:  open,
		info:  info,
		files: make(map[string]File),
	}
	return efs, nil
}

// Open provides a file from the EmbeddedFileSystem else and error if
// the given file name could not be found.
func (c *EmbeddedFileSystem) Open(name string) (File, error) {
	f, ok := c.files[name]
	if ok {
		return f, nil
	}

	bin, err := c.open(name)
	if err != nil {
		return nil, err
	}

	info, err := c.info(name)
	if err != nil {
		return nil, err
	}

	cf := &CachedFile{
		path: name,
		bin:  bin,
		info: info,
	}
	c.files[name] = cf
	return cf, nil
}
