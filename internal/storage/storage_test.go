package storage

import (
	"os"
	"path"
	"testing"
)

func TestWriteFile(t *testing.T) {
	err := WriteFile("readme.md", "./a/b", []byte("hello world"))
	if err != nil {
		t.Errorf("WriteFile err:%s", err)
	}
}

func TestPath(t *testing.T) {
	t.Logf("path.Join(a/b, c): %s", path.Join("a/b", "c"))
}

func TestOSMkdir(t *testing.T) {
	err := os.MkdirAll("./a/b/c", os.FileMode(0777))
	if err != nil {
		t.Errorf("os.MkdirAll err:%s", err)
	}
}
