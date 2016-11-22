package fs

import (
	"errors"
	"os"
	"io"
)

// A CachedFile represents
type CachedFile struct {
	path string
	bin  []byte
	info os.FileInfo
	mark int
}

// Close always returns nil since any error had to of occurred before when
// when Open on the FileSystem was called.
func (c *CachedFile) Close() error {
	return nil
}

func (c *CachedFile) Read(p []byte) (n int, err error) {
	filesize := len(c.bin)
	bufsize := len(p)
	from := c.mark
	to := from + bufsize

	// Case: have transitioned mark to the file length
	if from >= filesize {
		return 0, io.EOF
	}

	// Case: last offset less then the file length
	if to <= filesize {
		src := c.bin[from:to]
		c.mark = to
		return copy(p, src), nil
	}

	src := c.bin[from:filesize]
	c.mark = to
	return copy(p, src), nil
}

// ReadAt fills the destination slice with bytes from the file starting
// at the provided offset, reading upto dest length, or the remainder of
// bytes in the file.  The number of bytes read and an error should one
// occur is returned.
func (c *CachedFile) ReadAt(p []byte, off int64) (n int, err error) {
	if off >= int64(len(c.bin)) {
		return -1, ErrOffsetBeyondBounds
	}
	a, b := int(off), int(len(p))
	src := c.bin[a : a+b]
	return copy(p, src), nil
}

func (c *CachedFile) Seek(offset int64, whence int) (int64, error) {
	return int64(0), errors.New("not implemented yet")
}

func (c *CachedFile) Stat() (os.FileInfo, error) {
	return c.info, nil
}
