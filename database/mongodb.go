package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// DataClient Struct
type DataClient struct {
	client       *mongo.Client
	databaseName string
	ctx          context.Context
}

// DataClient Struct
type DataCollection struct {
	ctx        context.Context
	collection *mongo.Collection
}

// Client mongoDb connection Client(Database string) DataClient
func Client(Database string) DataClient {
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

	return DataClient{
		client:       client,
		databaseName: Database,
		ctx:          ctx,
	}
}

// Collection connection (x DataClient) Collection(Collection string) DataCollection
func (x DataClient) Collection(Collection string) DataCollection {
	collection := x.client.Database(x.databaseName).Collection(Collection)
	return DataCollection{
		ctx:        x.ctx,
		collection: collection,
	}
}

// InsertOne mongoDb Insert Data (x DataCollection) InsertOne(data bson.D) (*mongo.InsertOneResult, error)
func (x DataCollection) InsertOne(data bson.D) (*mongo.InsertOneResult, error) {
	return x.collection.InsertOne(x.ctx, data)
}

// InsertMany mongoDb Insert Multiple Data (x DataCollection) InsertMany(dataset []interface{}) (*mongo.InsertManyResult, error)
func (x DataCollection) InsertMany(dataset []interface{}) (*mongo.InsertManyResult, error) {
	return x.collection.InsertMany(x.ctx, dataset)
}

// FindOne mongoDb Find One Data (x DataCollection) FindOne(query bson.M) (bson.M, error)
func (x DataCollection) FindOne(query bson.M) (bson.M, error) {
	if err := x.collection.FindOne(x.ctx, query).Decode(&query); err != nil {
		return bson.M{}, err
	}
	return query, nil
}

// FindById mongoDb Find By Id One Data (x DataCollection) FindById(id string) (bson.M, error)
func (x DataCollection) FindById(id string) (bson.M, error) {
	objectId, err1 := primitive.ObjectIDFromHex(id)
	if err1 != nil {
		return bson.M{}, err1
	}

	var data bson.M
	if err := x.collection.FindOne(x.ctx, bson.M{"_id": objectId}).Decode(&data); err != nil {
		return bson.M{}, err
	}
	return data, nil
}

// Find mongoDb Find All Data (x DataCollection) Find(query bson.M) ([]bson.M, error)
func (x DataCollection) Find(query bson.M) ([]bson.M, error) {
	var filtered []bson.M
	filter, err := x.collection.Find(x.ctx, query)
	if err != nil {
		return filtered, err
	}

	if err = filter.All(x.ctx, &filtered); err != nil {
		return filtered, err
	}

	return filtered, nil
}
