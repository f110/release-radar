package producer

import (
	"net/url"

	"golang.org/x/xerrors"
)

type Manufacturer interface {
	Stop() error
}

type Producer struct {
	db *DataSource
}

func NewProducer(filename string) (*Producer, error) {
	ds, err := NewDataSource(filename)
	if err != nil {
		return nil, xerrors.Errorf(": %v", err)
	}

	return &Producer{db: ds}, nil
}

func (p *Producer) Add(t, u string) error {
	addr, err := url.Parse(u)
	if err != nil {
		return xerrors.Errorf(": %v", err)
	}

	return nil
}

func (p *Producer) Start() error {
	return nil
}
