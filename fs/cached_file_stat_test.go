package fs

import (
	"testing"
	. "github.com/lcaballero/exam/assert"
)

func Test_Stat_001(t *testing.T) {
	t.Log("Stat should provide the name of the flie")

	fs, _ := NewCachedFiles(mountAt())
	f, err := fs.Open("example-file.txt")

	IsNil(t, err)

	info, err := f.Stat()
	IsNil(t, err)
	IsEqStrings(t, info.Name(), "example-file.txt")
	IsFalse(t, info.IsDir())
}
