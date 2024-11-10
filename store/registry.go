package store

import (
	"digdocs-back/domain"
	"digdocs-back/postgres"

	"github.com/go-kit/kit/log"
)

type registry struct {
	postgres.QueryController
	logger log.Logger
}

func NewRegistry(controller postgres.QueryController, logger log.Logger) domain.Repository {
	return &registry{QueryController: postgres.NewController(controller), logger: logger}
}
