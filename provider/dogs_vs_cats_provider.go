package provider

import (
	"github.com/haormj/baize-api/idl/dogsvscats"
	"github.com/haormj/baize-api/option"
	"github.com/haormj/baize-api/service"

	"github.com/micro/go-micro/v2"
)

type DogsVsCatsProvider struct {
	opt option.DogsVsCatsProviderOpt
	svc *service.DogsVsCatsService
}

func NewDogsVsCatsProvider(opt option.DogsVsCatsProviderOpt,
	svc *service.DogsVsCatsService) (*DogsVsCatsProvider, error) {
	p := DogsVsCatsProvider{
		opt: opt,
		svc: svc,
	}
	return &p, nil
}

func (p *DogsVsCatsProvider) Start() error {
	s := micro.NewService(
		micro.Name("go.micro.api.DogsVsCatsService"),
		micro.Version("latest"),
	)
	s.Init()

	dogsvscats.RegisterDogsVsCatsServiceHandler(s.Server(), p.svc)

	if err := s.Run(); err != nil {
		return err
	}

	return nil
}

func (p *DogsVsCatsProvider) Stop() error {
	return nil
}
