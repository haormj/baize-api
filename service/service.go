package service

import (
	"github.com/haormj/baize-api/option"
)

type Service struct {
	opt               option.ServiceOpt
	DogsVsCatsService *DogsVsCatsService
}

func New(opt option.ServiceOpt) (*Service, error) {
	dvcs, err := NewDogsVsCatsService(opt.DogsVsCatsServiceOpt)
	if err != nil {
		return nil, err
	}

	s := Service{
		opt:               opt,
		DogsVsCatsService: dvcs,
	}

	return &s, nil
}
