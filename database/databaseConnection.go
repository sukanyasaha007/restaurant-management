package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go-mongodb.org/mongo-driver/mongo"
	"go-mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/tools/go/analysis/passes/nilfunc"
)


func DBInstance() *mongo.Client {
	MongoDB:= "mongo://localhost:27017"
	fmt.Print(MongoDB)
	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDB))
	if err!= nil{
		log.Fatal(err)
	}
	ctx, cancel:= context.WithTimeout((context.Background(), 5*time.Second))

	defer cancel()

	err:= client.Connect()
	if err!= nil{
		log.Fatal(err)
	}
	fmt.Print("Connection established")
	return client
}

var client: *mongo.Client:= DBInstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection{
	var collection *mongo.Collection:= client.Database("restaurant").Collection(collectionName)
	return collection
}