package provider

import (
	"github.com/haormj/baize-api/option"
	"github.com/haormj/baize-api/service"
)

type Provider struct {
	pch                chan error
	opt                option.ProviderOpt
	svc                *service.Service
	DogsVsCatsProvider *DogsVsCatsProvider
}

func New(opt option.ProviderOpt, svc *service.Service) (*Provider, error) {
	dvcp, err := NewDogsVsCatsProvider(opt.DogsVsCatsProviderOpt, svc.DogsVsCatsService)
	if err != nil {
		return nil, err
	}

	p := Provider{
		pch:                make(chan error),
		opt:                opt,
		svc:                svc,
		DogsVsCatsProvider: dvcp,
	}
	return &p, nil
}

func (p *Provider) Run() chan error {
	go func() {
		if err := p.DogsVsCatsProvider.Start(); err != nil {
			p.pch <- err
		}
	}()

	return p.pch
}

func (p *Provider) Close() error {
	close(p.pch)
	return nil
}
