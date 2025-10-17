package main

import (
	"errors"

	"github.com/qdrant/go-client/qdrant"
)

type QdrantClient struct {
	cfg    *qdrant.Config
	client *qdrant.Client
}

func NewClient(cfg *qdrant.Config) *QdrantClient {

	client, err := qdrant.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	return &QdrantClient{
		cfg:    cfg,
		client: client,
	}
}

func (c *QdrantClient) BulkInsert(vectors [][]float32, payloads []map[string]any) error {
	return errors.New("TODO")
}
