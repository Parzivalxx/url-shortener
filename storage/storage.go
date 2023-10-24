package storage

import (
	"context"
	"log"
	"time"
	"url-shortener/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBStorage struct {
    client *mongo.Client
    db     *mongo.Database
    col    *mongo.Collection
}

func Storage(mongoURI, dbName, collectionName string) (*MongoDBStorage, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))

	if err != nil {
        return nil, err
    }

    return &MongoDBStorage{
        client: client,
        db:     client.Database(dbName),
        col:    client.Database(dbName).Collection(collectionName),
    }, nil
}

func (s *MongoDBStorage) Close() {
    if s.client != nil {
        err := s.client.Disconnect(context.Background())
        if err != nil {
            log.Fatal(err)
        }
    }
}

func (s *MongoDBStorage) Shorten(longURL string, numShortenedUrlChars int) (string, error) {
    shortURL := utils.GenerateRandomShortURL(numShortenedUrlChars)
    _, err := s.col.InsertOne(context.Background(), bson.M{"shortURL": shortURL, "longURL": longURL})
    if err != nil {
        return "", err
    }

    return shortURL, nil
}

func (s *MongoDBStorage) GetOriginal(shortURL string) (string, error) {
    var result struct {
        LongURL string `bson:"longURL"`
    }
    filter := bson.M{"shortURL": shortURL}
    err := s.col.FindOne(context.Background(), filter).Decode(&result)
    if err != nil {
        return "", err
    }

    return result.LongURL, nil
}
