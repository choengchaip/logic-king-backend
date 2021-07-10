package cores

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

type Mongo struct {
	Client   *mongo.Client
	Database *mongo.Database
}

type MongoOption struct {
	DBName   string
	Username string
	Password string
	URL      string
	Port     string
}

func NewMongo(opts *MongoOption) *Mongo {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%s", opts.Username, opts.Password, opts.URL, opts.Port)))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	return &Mongo{
		Client:   client,
		Database: client.Database(opts.DBName),
	}
}

func (m *Mongo) Disconnect(ctx context.Context) error {
	err := m.Client.Disconnect(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (m *Mongo) GetContext() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	return ctx
}
