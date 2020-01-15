package producer

import (
	"time"

	mapset "github.com/deckarep/golang-set"
)

type Release struct {
	Version   Version
	Published time.Time
	Author    string
}

type Releases []*Release

func (r Releases) Diff(other Releases) Releases {
	diff := r.ToSet().Difference(other.ToSet())
	res := make([]*Release, 0)
	for _, v := range r {
		if diff.Contains(v.Version.String()) {
			res = append(res, v)
		}
	}

	return res
}

func (r Releases) ToSet() mapset.Set {
	if len(r) == 0 {
		return mapset.NewSet(nil)
	}

	set := mapset.NewSet(r[0].Version.String())
	if len(r) > 1 {
		for _, v := range r[1:] {
			set.Add(v.Version.String())
		}
	}

	return set
}
