package interfaces

import (
	"context"
	"digdocs-back/domain"
)

type DocumentService interface {
	StoreDocument(ctx context.Context, document domain.Document) (domain.Document, error)
	GetDocument(ctx context.Context, id uint64) (domain.Document, error)
	GetDocuments(ctx context.Context) ([]domain.Document, error)
	DeleteDocument(ctx context.Context, id uint64) error
}
