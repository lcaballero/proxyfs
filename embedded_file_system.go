package proxyfs

type EmbeddedFileSystem struct {
	files map[string]File
}

func NewEmbeddedFileSystem() (*EmbeddedFileSystem, error) {
	efs := &EmbeddedFileSystem{
		files: make(map[string]File),
	}
	return efs, nil
}

func (c *EmbeddedFileSystem) Open(name string) (File, error) {
	f, ok := c.files[name]
	if ok {
		return f, nil
	}

	bin, err := Asset(name)
	if err != nil {
		return nil, err
	}

	info, err := AssetInfo(name)
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
