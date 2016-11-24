package proxyfs

import (
	. "github.com/lcaballero/exam/assert"
	"io/ioutil"
	"os"
	"testing"
)

func mockInfoProvider(name string) (os.FileInfo, error) {
	return nil, nil
}

func mockOpenProvider(name string) ([]byte, error) {
	return []byte{}, nil
}

func Test_Embedded_File_System_005(t *testing.T) {
	t.Log("Contents of the file should be as expected")
	fs, _ := NewEmbeddedFileSystem(Asset, AssetInfo)
	f, _ := fs.Open("file-1.txt")
	bin, err := ioutil.ReadAll(f)
	IsNil(t, err)
	IsEqBytes(t, bin, []byte("Here's the first file\n"))
}

func Test_Embedded_File_System_004(t *testing.T) {
	t.Log("Should not find a file that isn't embedded")
	fs, err := NewEmbeddedFileSystem(Asset, AssetInfo)
	IsNotNil(t, fs)
	IsNil(t, err)

	f, err := fs.Open("file-1.txt")
	IsNotNil(t, f)
	IsNil(t, err)
}

func Test_Embedded_File_System_003(t *testing.T) {
	t.Log("Should create the embedded file system with proper providers")
	fs, err := NewEmbeddedFileSystem(Asset, AssetInfo)
	IsNotNil(t, fs)
	IsNil(t, err)
}

func Test_Embedded_File_System_002(t *testing.T) {
	t.Log("Should return an error if InfoProvider is nil")
	fs, err := NewEmbeddedFileSystem(nil, mockInfoProvider)
	IsNil(t, fs)
	IsNotNil(t, err)
}

func Test_Embedded_File_System_001(t *testing.T) {
	t.Log("Should return an error if OpenProvider is nil")
	fs, err := NewEmbeddedFileSystem(mockOpenProvider, nil)
	IsNil(t, fs)
	IsNotNil(t, err)
}
