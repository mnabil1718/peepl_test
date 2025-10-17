package main

import (
	"context"
	"fmt"

	"github.com/qdrant/go-client/qdrant"
)

const CollectionName = "articles"

func main() {

	db := NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})

	db.client.CreateCollection(context.Background(), &qdrant.CreateCollection{
		CollectionName: CollectionName,
		VectorsConfig: qdrant.NewVectorsConfig(&qdrant.VectorParams{
			Size:     4,
			Distance: qdrant.Distance_Cosine,
		}),
	})

	_, err := db.client.Upsert(context.Background(), &qdrant.UpsertPoints{
		CollectionName: CollectionName,
		Points: []*qdrant.PointStruct{
			{
				Id:      qdrant.NewIDNum(1),
				Vectors: qdrant.NewVectors(0.05, 0.61, 0.76, 0.74),
				Payload: qdrant.NewValueMap(map[string]any{"title": "Physics"}),
			},
			{
				Id:      qdrant.NewIDNum(2),
				Vectors: qdrant.NewVectors(0.19, 0.81, 0.75, 0.11),
				Payload: qdrant.NewValueMap(map[string]any{"title": "Economics"}),
			},
			{
				Id:      qdrant.NewIDNum(3),
				Vectors: qdrant.NewVectors(0.36, 0.55, 0.47, 0.94),
				Payload: qdrant.NewValueMap(map[string]any{"title": "Politics"}),
			},
		},
	})
	if err != nil {
		panic(err)
	}

	searchResult, err := db.client.Query(context.Background(), &qdrant.QueryPoints{
		CollectionName: CollectionName,
		Query:          qdrant.NewQuery(0.1, 0.2, 0.3, 0.4),
		WithPayload:    qdrant.NewWithPayload(true),
		Limit:          qdrant.PtrOf(uint64(1)),
	})
	if err != nil {
		panic(err)
	}

	for _, ent := range searchResult {
		p := ent.GetPayload()
		fmt.Println("title: ", p["title"].GetStringValue())
	}

}
