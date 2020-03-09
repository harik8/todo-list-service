package dbconnector

import (

	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetClient() *mongo.Client {
	/**
		export MONGO_URI=mongodb+srv://admin:admin@mongo0-e3lih.mongodb.net/test?retryWrites=true&w=majority

		context.TODO() vs context.Background()
		In terms of functionality: none. They are exactly the same value, bit-by-bit.
		https://www.programming-books.io/essential/go/context-todo-vs-context-background-d5224e27ff724a33a79cb4e03a5eb333
	**/
	//	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))
	clientOptions := options.Client().ApplyURI("mongodb+srv://admin:admin@mongo0-e3lih.mongodb.net/test?retryWrites=true&w=majority")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}
	return client
}

func GetDBNames(client *mongo.Client) []string {
	dbs, err := client.ListDatabaseNames(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	return dbs
}

func GetCollection(client *mongo.Client) *mongo.Collection {
	collection := client.Database("todo").Collection("tasks")
	return collection
}