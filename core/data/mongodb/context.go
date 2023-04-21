package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Struct used for maintaining HTTP Request Context
type MongoClient struct {
	Context context.Context
}

var client *mongo.Client
var dummy string = "User-Rels"
var dbName *string = &dummy

// Create a new mogodb client object
func NewMongoClient(ctx context.Context, mongoDBHost string, database string) (MongoDatabase, error) {
	fmt.Println("at new mongoclient function")
	mongoClient := &MongoClient{Context: ctx}
	fmt.Println("mongo client has got the value")
	var err error
	fmt.Println("at line no 24")
	client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://asb:asbdesigngo@cluster0.ninbdmn.mongodb.net/?retryWrites=true&w=majority"))

	fmt.Println("go the client")
	// *dbName = string("User-Rels")

	fmt.Println("db name is assigned with user rels")
	fmt.Println("new mongo client has been created")
	return mongoClient, err
}

// GetContext get mongodb context
func (c *MongoClient) GetContext() (context.Context, error) {
	if client == nil {
		return nil, fmt.Errorf("MongoDB client is nil, be sure you have invoked NewClient() function already!")
	}
	return c.Context, nil
}

// Close Client
func (c *MongoClient) Close() error {
	return client.Disconnect(c.Context)
}

// GetCollection gets database collection
func (c *MongoClient) GetCollection(collectionName string) (*mongo.Collection, error) {
	if client == nil {
		return nil, fmt.Errorf("MongoDB client is nil, be sure you have invoked NewClient() function already!")
	}
	collection := client.Database(*dbName).Collection(collectionName)

	return collection, nil
}

// Create database session
func (c *MongoClient) GetDb() (*mongo.Database, error) {
	if client == nil {
		return nil, fmt.Errorf("MongoDB client is nil, be sure you have invoked NewClient() function already!")
	}
	db := client.Database(*dbName)

	return db, nil
}

// Ping database
func (c *MongoClient) Ping() error {
	if client == nil {
		_ = client.Ping(c.Context, readpref.Primary())
	}

	return nil
}
