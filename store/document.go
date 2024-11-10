package store

import (
	"context"
	"digdocs-back/domain"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/jackc/pgx/v5"
)

const (
	// TODO: попробовать *
	sqlDocumentSelect = `
SELECT  id,
		name,
		mime,
		is_file,
		is_public,
		modified,
		created
FROM documents
`
	sqlDocumentInsert = `
INSERT INTO documents ("name", "mime", "is_file", "is_public")
VALUES (@name, @mime, @is_file, @is_public)
RETURNING "id";
`

	sqlDocumentUpdate = `
UPDATE documents
SET "name" = @name,
    "mime" = @mime,
    "is_file" = @is_file,
    "is_public" = @is_public,
    "modified" = current_timestamp
`

	sqlDocumentDelete = `DELETE FROM documents WHERE id = @id`
)

func (r *registry) CreateDocument(ctx context.Context, document domain.Document) (domain.Document, error) {
	logger := log.With(r.logger, "method", "CreateDocument")

	rows, err := r.Query(ctx, sqlDocumentInsert, pgx.NamedArgs{
		"name":      document.Name,
		"mime":      document.Mime,
		"is_file":   document.IsFile,
		"is_public": document.IsPublic,
	})
	if err != nil {
		level.Error(logger).Log("msg", "Query error", "err", err)
		return domain.Document{}, err
	}

	id, err := pgx.CollectOneRow(rows, pgx.RowTo[uint64])
	if err != nil {
		level.Error(logger).Log("msg", "CollectOneRow error", "err", err)
		return domain.Document{}, err
	}

	return r.GetDocument(ctx, id)
}

func (r *registry) GetDocument(ctx context.Context, id uint64) (domain.Document, error) {
	logger := log.With(r.logger, "method", "GetDocument")

	rows, err := r.Query(ctx, sqlDocumentSelect+`WHERE id = @id`, pgx.NamedArgs{"id": id})
	if err != nil {
		level.Error(logger).Log("msg", "Query error", "err", err)
		return domain.Document{}, err
	}

	res, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[domain.Document])
	if err != nil {
		level.Error(logger).Log("msg", "CollectOneRow error", "err", err)
		return domain.Document{}, err
	}

	return res, err
}

func (r *registry) GetDocuments(ctx context.Context) ([]domain.Document, error) {
	logger := log.With(r.logger, "method", "GetDocuments")

	rows, err := r.Query(ctx, sqlDocumentSelect)
	if err != nil {
		level.Error(logger).Log("msg", "Query error", "err", err)
		return nil, err
	}

	res, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.Document])
	if err != nil {
		level.Error(logger).Log("msg", "CollectRows error", "err", err)
		return nil, err
	}

	return res, err
}

func (r *registry) UpdateDocument(ctx context.Context, document domain.Document) (domain.Document, error) {
	logger := log.With(r.logger, "method", "UpdateDocument")

	_, err := r.Query(ctx, sqlDocumentUpdate+`WHERE id = @id`, pgx.NamedArgs{"id": document.ID})
	if err != nil {
		level.Error(logger).Log("msg", "Query error", "err", err)
		return domain.Document{}, err
	}

	return r.GetDocument(ctx, document.ID)
}

func (r *registry) DeleteDocument(ctx context.Context, id uint64) error {
	logger := log.With(r.logger, "method", "DeleteDocument")

	_, err := r.Query(ctx, sqlDocumentDelete, pgx.NamedArgs{"id": id})
	if err != nil {
		level.Error(logger).Log("msg", "Query error", "err", err)
		return err
	}

	return nil
}
