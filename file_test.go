package webutils

import (
	"testing"
)

var noExistedFile = "/tmp/not_existed_file"

func TestSelfPath(t *testing.T) {
	path := SelfPath()
	if path == "" {
		t.Error("path cannot be empty")
	}
	t.Logf("SelfDir", "%s", path)
}
