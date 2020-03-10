package dbconnector

import (

	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	conf "github.com/harik8/todo-list-service/config"
)

// GetClient : Get a DB instance
func GetClient() *mongo.Client {
	/**
		context.TODO() vs context.Background()
		In terms of functionality: none. They are exactly the same value, bit-by-bit.
		https://www.programming-books.io/essential/go/context-todo-vs-context-background-d5224e27ff724a33a79cb4e03a5eb333
	**/
	clientOptions := options.Client().ApplyURI(conf.MongoURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	} else {
		log.Println("Established DB connection")
	}
	return client
}

// GetCollection : Func to handle DB connection
func GetCollection(client *mongo.Client) *mongo.Collection {
	collection := client.Database(conf.MongoDatabase).Collection(conf.MongoCollection)
	return collection
}