package config

import (
	"fmt"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"context"
)

var DB *mongo.Database

var Books *mongo.Collection

func init() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://bond:moneypenny007@localhost:27017/bookstore"))
	if err != nil {
		panic(err)
	}

	if err = client.Ping(context.TODO(), nil); err != nil {
		panic(err)
	}

	DB = client.Database("bookstore")
	Books = DB.Collection("books")

	fmt.Println("You are connected to your mongo database")
}