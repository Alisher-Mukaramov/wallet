package service

import (
	"github.com/Alisher-Mukaramov/wallet/internal/pkg/repository"
	"go.uber.org/fx"
)

var Module = fx.Provide(newService)

type Iservice interface {
	Repository() repository.Irepository
}

type serviceParams struct {
	fx.In
	repository.Irepository
}

type service struct {
	repository.Irepository
}

func newService(sp serviceParams) Iservice {
	return &service{sp.Irepository}
}

func (s *service) Repository() repository.Irepository {
	return s
}
