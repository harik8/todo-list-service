package dbexecutor

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	conn "github.com/harik8/todo-list-service/dbconnector"
	todo "github.com/harik8/todo-list-service/todo"
)

// AddTodo : Add a Todo
func AddTodo(t todo.Todo) *mongo.InsertOneResult {
	client := conn.GetClient()
	collection := conn.GetCollection(client)
	defer client.Disconnect(context.TODO())

	insertResult, err := collection.InsertOne(context.TODO(), t)
	if err != nil {
			log.Fatal(err)
	}
	return insertResult
}

// UpdateTodo : Update a given Todo
func UpdateTodo(tid primitive.ObjectID, t todo.Todo) error {
	client := conn.GetClient()
	collection := conn.GetCollection(client)
	defer client.Disconnect(context.TODO())

	filter := bson.D{primitive.E{Key: "_id", Value: tid}}
	set := bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{Key: "task", Value: t.Task},
		primitive.E{Key: "duedate", Value: t.Duedate},
		primitive.E{Key: "labels", Value: t.Labels},
		primitive.E{Key: "comments", Value: t.Comments},
	}}}

	err := collection.FindOneAndUpdate(context.TODO(), filter, set).Decode(&t)
	return err
}

// DeleteTodo : Delete a given Todo
func DeleteTodo(tid primitive.ObjectID) (*mongo.DeleteResult, error) {
	client := conn.GetClient()
	collection := conn.GetCollection(client)
	defer client.Disconnect(context.TODO())

	filter := bson.D{primitive.E{Key: "_id", Value: tid}}
	d, err := collection.DeleteOne(context.TODO(), filter)
	return d, err
}

// GetTodo : Read a given Todo
func GetTodo(tid primitive.ObjectID) (todo.Todo, error) {
	var td todo.Todo
	client := conn.GetClient()
	collection := conn.GetCollection(client)
	defer client.Disconnect(context.TODO())

	filter := bson.D{primitive.E{Key: "_id", Value: tid}}
	err := collection.FindOne(context.TODO(), filter).Decode(&td)
	return td, err
}

// GetTodos : Read all Todos
func GetTodos() ([]todo.Todo, error) {
	var td []todo.Todo
	client := conn.GetClient()
	collection := conn.GetCollection(client)
	cur, err := collection.Find(context.TODO(), bson.D{})
	for cur.Next(context.Background()) {
		var t todo.Todo
		cur.Decode(&t)
		td = append(td, t)
	}
	return td, err
}