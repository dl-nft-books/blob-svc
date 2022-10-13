package service

import (
	"gitlab.com/distributed_lab/logan/v3"
	"net"
	"net/http"

	"gitlab.com/tokend/nft-books/blob-svc/internal/config"

	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type service struct {
	log       *logan.Entry
	copus     types.Copus
	listener  net.Listener
	mimeTypes *config.MimeTypes
	aws       *config.AWSConfig
	jwt       *config.JWT
}

func (s *service) run() error {
	s.log.Info("Service started")
	r := s.router()

	if err := s.copus.RegisterChi(r); err != nil {
		return errors.Wrap(err, "cop failed")
	}

	return http.Serve(s.listener, r)
}

func newService(cfg config.Config) *service {
	return &service{
		log:       cfg.Log(),
		copus:     cfg.Copus(),
		listener:  cfg.Listener(),
		mimeTypes: cfg.MimeTypes(),
		aws:       cfg.AWSConfig(),
		jwt:       cfg.JWT(),
	}
}

func Run(cfg config.Config) {
	if err := newService(cfg).run(); err != nil {
		panic(err)
	}
}