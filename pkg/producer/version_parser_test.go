package producer

import (
	"testing"
)

func TestParseVersionString(t *testing.T) {
	cases := []struct {
		In     string
		Expect Version
	}{
		{In: "v0.0.1", Expect: Version{Major: 0, Minor: 0, Patch: 1}},
		{In: "v1.2.3", Expect: Version{Major: 1, Minor: 2, Patch: 3}},
		{In: "4.5.6", Expect: Version{Major: 4, Minor: 5, Patch: 6}},
	}

	for _, v := range cases {
		got, err := ParseSemVerString(v.In)
		if err != nil {
			t.Fatal(err)
		}

		if !v.Expect.Equal(got) {
			t.Errorf("expect: %v got: %v", v.Expect, got)
		}
	}

	errCases := []string{"v0.1", "v1", "0.1", "1"}
	for _, v := range errCases {
		_, err := ParseSemVerString(v)
		if err == nil {
			t.Errorf("%s: expect to fail parse but not", v)
		}
	}
}

func TestVersion_Equal(t *testing.T) {
	if !(Version{1, 2, 3}.Equal(Version{1, 2, 3})) {
		t.Errorf("expect to return true")
	}

	if (Version{1, 2, 3}.Equal(Version{1, 3, 2})) {
		t.Errorf("expect to return false")
	}
}

func TestVersion_Less(t *testing.T) {
	if !(Version{1, 2, 3}.Less(Version{2, 0, 0})) {
		t.Errorf("expect to return true")
	}
	if !(Version{1, 2, 3}.Less(Version{1, 3, 0})) {
		t.Errorf("expect to return true")
	}
	if !(Version{1, 2, 3}.Less(Version{1, 2, 4})) {
		t.Errorf("expect to return true")
	}

	if (Version{1, 2, 3}.Less(Version{0, 2, 3})) {
		t.Errorf("expect to return false")
	}
	if (Version{1, 2, 3}.Less(Version{1, 1, 3})) {
		t.Errorf("expect to return false")
	}
	if (Version{1, 2, 3}.Less(Version{1, 2, 2})) {
		t.Errorf("expect to return false")
	}
}
