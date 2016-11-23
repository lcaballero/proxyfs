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
	mark int64
}

// Close always returns nil since any error had to of occurred before when
// when Open on the FileSystem was called.
func (c *CachedFile) Close() error {
	return nil
}

// Read implements io.Read function but over a CachedFile.  After
// using this method to read a file the mark will be set at the end
// of the file, so in order to reset it use the Seek method below.
func (c *CachedFile) Read(p []byte) (n int, err error) {
	filesize := int64(len(c.bin))
	bufsize := int64(len(p))
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

// FromTheStart is a constant that can be used for 'whence' in the
// Seek method.
const FromTheStart = 0

// FromTheCurrentOffset is a constant that can be used for 'whence' in the
// Seek method.
const FromTheCurrentOffset = 1

// FromTheEnd is a constant that can be used for 'whence' in the Seek method.
const FromTheEnd = 2

// ErrNegativeReadOffsetAfterSeek occurs if the next read offset after a call
// to Seek is would be negative.
var ErrNegativeReadOffsetAfterSeek = errors.New("negative read offset after seek")

// ErrOffsetGreaterThanFileSizeAfterSeek occurs if the next read offset is
// beyond the end of the file after a Seek call.
var ErrOffsetGreaterThanFileSizeAfterSeek = errors.New("offset greater then file length after seek")

//
func (c *CachedFile) Seek(offset int64, whence int) (int64, error) {
	var pos int64
	switch whence {
	case FromTheStart:
		pos = offset
	case FromTheCurrentOffset:
		pos = c.mark + offset
	case FromTheEnd:
		pos = int64(len(c.bin)) - offset
	}

	if pos < 0 {
		return 0, ErrNegativeReadOffsetAfterSeek
	}
	if pos > int64(len(c.bin)) {
		return 0, ErrOffsetGreaterThanFileSizeAfterSeek
	}

	c.mark = pos
	return int64(c.mark), nil
}

func (c *CachedFile) Stat() (os.FileInfo, error) {
	return c.info, nil
}
