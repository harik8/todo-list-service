package dbexecutor

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	conn "github.com/harik8/todo-list-service/dbconnector"
)

type Todo struct {
	//	TID			primitive.ObjectID 	`json:"_id" 			bson:"_id"`
	task     string `json:"task" 			bson:"task"`
	duedate  string `json:"duedate" 		bson:"duedate"`
	labels   string `json:"labels"			bson:"labels"`
	comments string `json:"comments"  		bson:"comments"`
}

// func AddTodo(t Todo) *mongo.InsertOneResult {
// 	client := conn.GetClient()
// 	collection := conn.GetCollection(client)
// 	defer client.Disconnect(context.TODO())
// 	log.Println(t)
// 	insertResult, err := collection.InsertOne(context.TODO(), t)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return insertResult
// }

func AddTodo(t Todo) *mongo.InsertOneResult {
	client := conn.GetClient()
	collection := conn.GetCollection(client)
	defer client.Disconnect(context.TODO())

	insertResult, err := collection.InsertOne(context.TODO(), t)
	if err != nil {
			log.Fatal(err)
	}
	return insertResult
}

func UpdateTodo(tid primitive.ObjectID, t Todo) error {
	client := conn.GetClient()
	collection := conn.GetCollection(client)
	defer client.Disconnect(context.TODO())

	filter := bson.D{primitive.E{Key: "_id", Value: tid}}
	set := bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{Key: "task", Value: t.task},
		primitive.E{Key: "duedate", Value: t.duedate},
		primitive.E{Key: "labels", Value: t.labels},
		primitive.E{Key: "comments", Value: t.comments},
	}}}

	err := collection.FindOneAndUpdate(context.TODO(), filter, set).Decode(&t)
	return err
}

func DeleteTodo(tid primitive.ObjectID) (*mongo.DeleteResult, error) {
	client := conn.GetClient()
	collection := conn.GetCollection(client)
	filter := bson.D{primitive.E{Key: "_id", Value: tid}}
	d, err := collection.DeleteOne(context.TODO(), filter)
	return d, err
}

func GetTodo(tid primitive.ObjectID) (Todo, error) {
	var td Todo
	client := conn.GetClient()
	collection := conn.GetCollection(client)
	defer client.Disconnect(context.TODO())

	filter := bson.D{primitive.E{Key: "_id", Value: tid}}
	err := collection.FindOne(context.TODO(), filter).Decode(&td)
	return td, err
}

func GetTodos() ([]Todo, error) {
	var td []Todo
	client := conn.GetClient()
	collection := conn.GetCollection(client)
	cur, err := collection.Find(context.TODO(), bson.D{})
	for cur.Next(context.Background()) {
		var t Todo
		cur.Decode(&t)
		td = append(td, t)
	}
	return td, err
}