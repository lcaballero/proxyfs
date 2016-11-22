package fs

import (
	"os"
	"time"
)

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
	return cf.info.IsDir()
}

// underlying data source (can return nil)
func (cf CachedFileInfo) Sys() interface{} {
	return cf.info.Sys()
}
