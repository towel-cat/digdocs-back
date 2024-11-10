package implementation

import (
	"context"
	"digdocs-back/domain"
)

func (s *service) StoreDocument(ctx context.Context, document domain.Document) (domain.Document, error) {
	if document.ID > 0 {
		return s.repository.UpdateDocument(ctx, document)
	}
	return s.repository.CreateDocument(ctx, document)
}

func (s *service) GetDocument(ctx context.Context, id uint64) (domain.Document, error) {
	return s.repository.GetDocument(ctx, id)
}

func (s *service) GetDocuments(ctx context.Context) ([]domain.Document, error) {
	return s.repository.GetDocuments(ctx)
}

func (s *service) DeleteDocument(ctx context.Context, id uint64) error {
	return s.repository.DeleteDocument(ctx, id)
}
