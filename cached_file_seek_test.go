package proxyfs

import (
	. "github.com/lcaballero/exam/assert"
	"testing"
)

const abc = "abc edf xyz"

func aCachedFile(bin string) *CachedFile {
	cf := &CachedFile{
		bin: []byte(bin),
	}
	return cf
}

func abcCachedFile() *CachedFile {
	return aCachedFile(abc)
}

func Test_CachedFile_Seek_005(t *testing.T) {
	t.Log("seek from the end to a negative offset returns an err")
	cf := abcCachedFile()
	offset := int64(len(abc) + 3)
	mark, err := cf.Seek(offset, FromTheEnd)
	IsNotNil(t, err)
	IsTrue(t, mark == 0)
}

func Test_CachedFile_Seek_004(t *testing.T) {
	t.Log("seek to a negative offset returns an err")
	cf := abcCachedFile()
	mark, err := cf.Seek(-11, FromTheStart)
	IsNotNil(t, err)
	IsTrue(t, mark == 0)
}

func Test_CachedFile_Seek_003(t *testing.T) {
	t.Log("seek from the end of the file puts the offset at the right spot")
	cf := abcCachedFile()
	mark, err := cf.Seek(3, FromTheEnd)
	IsNil(t, err)
	IsTrue(t, mark == int64(len(abc)-3))
}

func Test_CachedFile_Seek_002(t *testing.T) {
	t.Log("seek from current mark puts the offset further into the file")
	cf := abcCachedFile()
	mark, err := cf.Seek(3, FromTheStart)
	mark, err = cf.Seek(3, FromTheCurrentOffset)
	IsNil(t, err)
	IsTrue(t, mark == 6)
}

func Test_CachedFile_Seek_001(t *testing.T) {
	t.Log("simple seek from the beginning sets the mark to the offset")
	cf := abcCachedFile()
	mark, err := cf.Seek(3, FromTheStart)
	IsNil(t, err)
	IsTrue(t, mark == 3)
}
