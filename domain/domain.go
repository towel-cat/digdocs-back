package domain

import "digdocs-back/postgres"

type Repository interface {
	postgres.QueryController
	DocumentRepository
}
