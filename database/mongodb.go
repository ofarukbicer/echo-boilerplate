package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

// DataClient Struct
type DataClient struct {
	mongoClient *mongo.Client
	mongoCtx context.Context
	mongoErr error
	mongoCollection *mongo.Collection
}

// Client mongoDb connection Client(Database string, Collection string) DataClient
func Client(Database string, Collection string) DataClient {
	client, err := mongo.NewClient() // local mongodb
	// client, err := mongo.NewClient(options.Client().ApplyURI("MongoDB-Atlas-URL-String"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database(Database).Collection(Collection)
	return DataClient{
		mongoClient: client,
		mongoCtx: ctx,
		mongoErr: err,
		mongoCollection: collection,
	}
}

// InsertOne mongoDb Insert Data (x DataClient) InsertOne(data bson.D) (*mongo.InsertOneResult, error)
func (x DataClient) InsertOne(data bson.D) (*mongo.InsertOneResult, error) {
	return x.mongoCollection.InsertOne(x.mongoCtx, data)
}

// FindOne mongoDb Find One Data (x DataClient) FindOne(query bson.M) bson.M
func (x DataClient) FindOne(query bson.M) bson.M {
	if err := x.mongoCollection.FindOne(x.mongoCtx, query).Decode(&query); err != nil {
		log.Fatal(err)
	}
	return query
}

// Find mongoDb Find All Data (x DataClient) Find(query bson.M) []bson.M
func (x DataClient) Find(query bson.M) []bson.M {
	filter, err := x.mongoCollection.Find(x.mongoCtx, query)
	if err != nil {
		log.Fatal(err)
	}
	var filtered []bson.M
	if err = filter.All(x.mongoCtx, &filtered); err != nil {
		log.Fatal(err)
	}

	return filtered
}