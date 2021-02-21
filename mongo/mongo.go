package mongo

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var cli *mongo.Client

// Connect to mongo db
func Connect(addr string, port uint16, user string, password string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	url := "mongodb://" + user + ":" + password + "@" + addr + strconv.Itoa(int(port))
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	cli = client
	if err != nil {
		return
	}
	fmt.Println("success")
}

// GetCollection with database name
func GetCollection(databaseName string, collectionName string) *mongo.Collection {
	collection := cli.Database(databaseName).Collection(collectionName)
	return collection
}
