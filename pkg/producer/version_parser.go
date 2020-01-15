package producer

import (
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/xerrors"
)

type Version struct {
	Major int
	Minor int
	Patch int
}

func (v Version) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}

func ParseVersionString(in string) (Version, error) {
	return ParseSemVerString(in)
}

func ParseSemVerString(in string) (Version, error) {
	if strings.HasPrefix(in, "v") {
		in = in[1:]
	}

	s := strings.Split(in, ".")
	if len(s) != 3 {
		return Version{}, xerrors.New("producer: not sem ver")
	}
	major, err := strconv.ParseInt(s[0], 10, 32)
	if err != nil {
		return Version{}, xerrors.Errorf(": %v", err)
	}
	minor, err := strconv.ParseInt(s[1], 10, 32)
	if err != nil {
		return Version{}, xerrors.Errorf(": %v", err)
	}
	patch, err := strconv.ParseInt(s[2], 10, 32)
	if err != nil {
		return Version{}, xerrors.Errorf(": %v", err)
	}

	return Version{Major: int(major), Minor: int(minor), Patch: int(patch)}, nil
}

func (v Version) Equal(right Version) bool {
	return v.Major == right.Major && v.Minor == right.Minor && v.Patch == right.Patch
}

func (v Version) Less(right Version) bool {
	return v.Major < right.Major || v.Minor < right.Minor || v.Patch < right.Patch
}
