package producer

import (
	"bytes"
	"encoding/gob"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
	"golang.org/x/xerrors"
)

type DataSource struct {
	*leveldb.DB
}

func NewDataSource(filename string) (*DataSource, error) {
	db, err := leveldb.OpenFile(filename, nil)
	if err != nil {
		return nil, xerrors.Errorf(": %v", err)
	}

	return &DataSource{DB: db}, nil
}

func (d *DataSource) ListApp() ([]string, error) {
	v, err := d.Get([]byte("apps"), nil)
	if err != nil {
		return nil, xerrors.Errorf(": %v", err)
	}

	apps := make([]string, 0)
	if err := gob.NewDecoder(bytes.NewReader(v)).Decode(apps); err != nil {
		return nil, xerrors.Errorf(": %v", err)
	}

	return apps, nil
}

func (d *DataSource) List(app string) ([]*Release, error) {
	prefix := []byte(app + "/")

	res := make([]*Release, 0)
	iter := d.NewIterator(&util.Range{Start: prefix}, nil)
	for iter.Next() {
		if !bytes.HasPrefix(iter.Key(), prefix) {
			continue
		}

		v := &Release{}
		if err := gob.NewDecoder(bytes.NewReader(iter.Value())).Decode(&v); err != nil {
			continue
		}
		res = append(res, v)
	}

	return res, nil
}

func (d *DataSource) Set(app string, r *Release) error {
	buf := new(bytes.Buffer)
	if err := gob.NewEncoder(buf).Encode(r); err != nil {
		return xerrors.Errorf(": %v", err)
	}

	return d.Put(d.versionKey(app, r.Version), buf.Bytes(), nil)
}

func (d *DataSource) versionKey(app string, v Version) []byte {
	return []byte("version/" + app + "/" + v.String())
}
