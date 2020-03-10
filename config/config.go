package config

import (
	"os"
)

var (
// DB configs

// MongoHost : Mongo DB host
MongoHost = os.Getenv("MONGO_HOST")

// MongoPort : Mongo DB service port
MongoPort = os.Getenv("MONGO_PORT")

// MongoUsername : Mongo DB user
MongoUsername = os.Getenv("MONGO_USERNAME")

// MongoPassword : Mongo DB password
MongoPassword = os.Getenv("MONGO_PASSWORD")

// MongoDatabase : Mongo Database name
MongoDatabase = os.Getenv("MONGO_DATABASE")

// MongoCollection : Mongo Database collection
MongoCollection = os.Getenv("MONGO_COLLECTION")

// MongoURI : Mongo DB connection string
MongoURI = "mongodb://" + MongoUsername + ":" + MongoPassword + "@" + MongoHost + ":" + MongoPort

// Application config

// TodoServiceIP : Todo service IP address
TodoServiceIP = os.Getenv("TODO_SERVICE_IP")

// TodoServicePort : Todo service port
TodoServicePort = os.Getenv("TODO_SERVICE_PORT")
) 