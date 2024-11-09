package domain

import (
	"context"
	"time"
)

type Document struct {
	ID       uint64    `db:"id"`
	Name     string    `db:"name"`
	Mime     string    `db:"mime"`
	IsFile   bool      `db:"is_file"`
	IsPublic bool      `db:"is_public"`
	Modified time.Time `db:"modified"`
	Created  time.Time `db:"created"`
}

type DocumentRepository interface {
	CreateDocument(ctx context.Context, document *Document) (*Document, error)
	GetDocument(ctx context.Context, id uint64) (*Document, error)
	GetDocuments(ctx context.Context) ([]Document, error)
	UpdateDocument(ctx context.Context, document *Document) (*Document, error)
	DeleteDocument(ctx context.Context, id uint64) error
}
