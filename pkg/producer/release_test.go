package producer

import "testing"

func TestReleases_Diff(t *testing.T) {
	this := Releases{
		{Version: Version{1, 2, 3}},
		{Version: Version{3, 4, 5}},
		{Version: Version{2, 6, 7}},
		{Version: Version{9, 8, 7}},
	}
	other := Releases{
		{Version: Version{1, 2, 3}},
		{Version: Version{2, 6, 7}},
		{Version: Version{9, 8, 7}},
		{Version: Version{6, 7, 8}},
	}

	diff := this.Diff(other)
	if len(diff) != 1 {
		t.Fatalf("expect diff element is one: %d", len(diff))
	}
	if !diff[0].Version.Equal(Version{3, 4, 5}) {
		t.Errorf("unexpected element: %s", diff[0].Version)
	}
}
