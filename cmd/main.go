package main

import (
	"context"
	"digdocs-back/domain"
	"digdocs-back/postgres"
	"digdocs-back/service/implementation"
	"digdocs-back/store"
	"fmt"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"os"
	"strconv"
)

func main() {
	logger := level.NewFilter(log.NewSyncLogger(log.NewLogfmtLogger(os.Stdout)), level.AllowAll())
	logger = log.With(logger, "ts", log.DefaultTimestamp, "caller", log.DefaultCaller)
	logger = log.With(logger, "svc", "digdocs")

	var pool postgres.PgPool
	repo := store.NewRegistry(pool.GetPool("default"), logger)
	svc := implementation.NewService(repo, logger)

	ctx, ctxCancel := context.WithCancel(context.Background())
	defer ctxCancel()

	var document domain.Document
	var err error
	for i := 0; i < 5; i++ {
		document, err = svc.StoreDocument(ctx, domain.Document{
			Name:     "Test name" + strconv.Itoa(i),
			Mime:     "Test mime" + strconv.Itoa(i),
			IsFile:   true,
			IsPublic: true,
		})
		if err != nil {
			level.Error(logger).Log("err", err)
			return
		}
	}

	document.Name = "abracadabra"
	document, err = svc.StoreDocument(ctx, document)
	if err != nil {
		level.Error(logger).Log("err", err)
		return
	}

	err = svc.DeleteDocument(ctx, document.ID-1)
	if err != nil {
		level.Error(logger).Log("err", err)
		return
	}

	documents, err := svc.GetDocuments(ctx)
	if err != nil {
		level.Error(logger).Log("err", err)
		return
	}
	fmt.Println(documents)
}
