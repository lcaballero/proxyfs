package proxyfs

import (
	. "github.com/lcaballero/exam/assert"
	"testing"
)

func Test_CachedFile_Read_At_002(t *testing.T) {
	t.Log("offset if beyond the bounds of the cached file binary")
}

func Test_CachedFile_Read_At_001(t *testing.T) {
	t.Log("reading the first slice of bytes from the file should match file contents")

	fs, _ := NewCachedFileSystem(mountAt())
	f, _ := fs.Open("example-file.txt")

	ct0 := "example-file\n"
	ct1 := "this should be readable"
	b := make([]byte, len(ct1))

	n, err := f.ReadAt(b, int64(len(ct0)))
	IsNil(t, err)
	IsEqInt(t, n, len(b))
	IsEqStrings(t, string(b), ct1)
}
