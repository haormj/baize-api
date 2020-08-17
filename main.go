package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/haormj/baize-api/option"
	"github.com/haormj/baize-api/provider"
	"github.com/haormj/baize-api/service"

	"github.com/haormj/log"
	"github.com/haormj/version"
)

func main() {
	v := flag.Bool("v", false, "show version")
	c := flag.Bool("c", false, "show config")
	flag.Parse()

	if *v {
		fmt.Println(version.FullVersion())
		return
	}

	if *c {
		opt, err := option.New()
		if err != nil {
			fmt.Printf("[ERR] option.NewOption error %v\n", err)
			return
		}
		fmt.Println(opt)
		return
	}

	opt, err := option.New()
	if err != nil {
		log.Logger.Errorw("NewOption error", "err", err)
		return
	}
	s, err := service.New(opt.ServiceOpt)
	if err != nil {
		log.Logger.Errorw("NewService error", "err", err)
		return
	}
	p, err := provider.New(opt.ProviderOpt, s)
	if err != nil {
		log.Logger.Errorw("NewProvider error", "err", err)
		return
	}
	pch := p.Run()
	defer p.Close()

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	select {
	case sig := <-ch:
		log.Logger.Infow("receive exit signal", "signal", sig)
	case err := <-pch:
		log.Logger.Errorw("Run error", "err", err)
	}
}
