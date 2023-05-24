package version

import (
	"testing"
)

func TestAll(t *testing.T) {
	full := Full()
	if full == UnknownVersion {
		t.Errorf("full build version is invalid")
	}

	min := Minimum()
	if min == UnknownVersion {
		t.Errorf("minimum build version is invalid")
	}

	build := Build(All)
	if build == UnknownVersion {
		t.Errorf("build version is invalid")
	}
}
