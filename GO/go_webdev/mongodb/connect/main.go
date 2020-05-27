// following the tutorial at https://www.mongodb.com/blog/post/mongodb-go-driver-tutorial
package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Trainer struct {
	Name string
	Age  int
	City string
}

func main() {
	// Setting Client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	errCheck(err)

	// Check connection
	err = client.Ping(context.TODO(), nil)
	errCheck(err)

	fmt.Println("Connected to MongoDB! YOU DID IT!")

	collection := client.Database("test").Collection("trainers")



	ash := Trainer{"Ash", 10, "Pallet Town"}
	misty := Trainer{"Misty", 10, "Cerulean City"}
	brock := Trainer{"Brock", 15, "Pewter City"}

	// Insert one document
	insertResult, err := collection.InsertOne(context.TODO(), ash)
	errCheck(err)

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	// Insert many documents
	trainers := []interface{}{misty, brock}

	insertManyResult, err := collection.InsertMany(context.TODO(), trainers)
	errCheck(err)

	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)

	// Update Document
	filter := bson.D{{"name", "Ash"}}

	update := bson.D{
		{"$inc", bson.D{
			{"age", 1},
		}},
	}

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	errCheck(err)

	fmt.Printf("Matched %v documents and updated %v documnets.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	// Finding documents
	var result Trainer

	// One documnet
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	errCheck(err)

	fmt.Printf("Found one document: %v\n", result)

	// Many documents
	findOptions := options.Find()
	findOptions.SetLimit(2)

	var results []*Trainer

	// Passing bson.D{{}} as this filter will match all documents in the collection
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	errCheck(err)

	// Using a Cursor (a Cursor provides a stream of documents)
	for cur.Next(context.TODO()) {

		// Creating a vlaue to decode a single document
		var elem Trainer
		err := cur.Decode(&elem)
		errCheck(err)

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Delete Documents
	deleteResult, err := collection.DeleteMany(context.TODO(), bson.D{{}})
	errCheck(err)

	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)

	// Close the cursor once finished
	cur.Close(context.TODO())

	fmt.Printf("Found multiple documents (array of pointers): %v\n", results)

	// Close Connection
	err = client.Disconnect(context.TODO())
	errCheck(err)
	
	fmt.Println("Connection to MongoDB is CLOSED!")
}

func errCheck(err  error) {
	if err != nil {
		log.Fatal(err)
	}
}
