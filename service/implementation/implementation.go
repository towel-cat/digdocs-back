package implementation

import (
	"digdocs-back/domain"
	serviceInterface "digdocs-back/service"
	"digdocs-back/store"
	"github.com/go-kit/log"
)

type service struct {
	repository domain.Repository
	logger     log.Logger
}

func NewService(repository domain.Repository, logger log.Logger) serviceInterface.Service {
	return &service{
		repository: store.NewRegistry(repository, logger),
		logger:     logger,
	}
}
