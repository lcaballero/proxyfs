package fs

import (
	"testing"
	. "github.com/lcaballero/exam/assert"
)

func Test_CachedFile_Read_004(t *testing.T) {
	t.Log("final read should return less than buffer size (when remaining bytes to read are fewer than buffer size)")

	cf := &CachedFile{
		bin: []byte("abc def xyz"),
	}

	read := make([]byte, 4)
	numBytesRead, _ := cf.Read(read)
	numBytesRead, _ = cf.Read(read)
	numBytesRead, _ = cf.Read(read)
	IsEqInt(t, numBytesRead, 3)
	IsEqBytes(t, read[0:3], []byte("xyz"))
}

func Test_CachedFile_Read_003(t *testing.T) {
	t.Log("an incomplete read followed by another read should return the next bytes")

	cf := &CachedFile{
		bin: []byte("abc def xyz"),
	}

	read := make([]byte, 4)
	numBytesRead, _ := cf.Read(read)
	numBytesRead, _ = cf.Read(read)
	IsEqInt(t, numBytesRead, len(read))
	IsEqBytes(t, read, []byte("def "))
}

func Test_CachedFile_Read_002(t *testing.T) {
	t.Log("the first read call should fill the buffer with bytes from the file")

	cf := &CachedFile{
		bin: []byte("abc def xyz"),
	}

	read := make([]byte, 4)
	numBytesRead, err := cf.Read(read)

	IsNil(t, err)
	IsEqInt(t, numBytesRead, len(read))
	IsEqBytes(t, read, []byte("abc "))
}

func Test_CachedFile_Read_001(t *testing.T) {
	t.Log("reading the first slice of bytes from the file should match file contents")

	fs, _ := NewCachedFileSystem(mountAt())
	f, _ := fs.Open("example-file.txt")

	ct := "example-file"
	b := make([]byte, len(ct))

	n, err := f.Read(b)
	IsNil(t, err)
	IsEqInt(t, n, len(b))
	IsEqStrings(t, string(b), ct)
}
