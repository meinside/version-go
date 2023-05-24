package version

import (
	"testing"
)

func TestAll(t *testing.T) {
	build := Build(All)
	if build == UnknownVersion {
		t.Errorf("build version is invalid")
	}
}
