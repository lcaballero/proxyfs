package fs

import (
	. "github.com/lcaballero/exam/assert"
	"os"
	"path/filepath"
	"testing"
)

func mountAt() string {
	dir, _ := os.Getwd()
	mount := filepath.Join(dir, ".files")
	return mount
}

func Test_Cached_File_003(t *testing.T) {
	t.Log("number of cached files should increase as new files are opened.")

	fs, _ := NewCachedFileSystem(mountAt())
	fs.Open("example-file.txt")
	cfs, ok := fs.(*CachedFileSystem)

	IsTrue(t, ok)
	IsEqInt(t, len(cfs.files), 1)

	fs.Open("example-file-2.txt")
	IsEqInt(t, len(cfs.files), 2)

	fs.Open("example-file.txt")
	IsEqInt(t, len(cfs.files), 2)
}

func Test_Cached_File_002(t *testing.T) {
	t.Log("number of cached files should not increase as files are opened.")

	fs, _ := NewCachedFileSystem(mountAt())
	fs.Open("example-file.txt")
	cfs, ok := fs.(*CachedFileSystem)

	IsTrue(t, ok)
	IsEqInt(t, len(cfs.files), 1)

	fs.Open("example-file.txt")
	IsEqInt(t, len(cfs.files), 1)
}

func Test_Cached_File_001b(t *testing.T) {
	t.Log("NewCachedFiles.Open() shouldn't return err.")
	fs, _ := NewCachedFileSystem(mountAt())

	f, err := fs.Open("example-file.txt")
	IsNil(t, err)
	IsNotNil(t, f)
}

func Test_Cached_File_001(t *testing.T) {
	t.Log("NewCachedFiles shouldn't return err.")

	fs, err := NewCachedFileSystem(mountAt())
	IsNil(t, err)
	IsNotNil(t, fs)

	f, err := fs.Open("example-file.txt")
	IsNil(t, err)
	IsNotNil(t, f)
}
