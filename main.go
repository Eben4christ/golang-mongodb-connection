package main

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoField struct {
	FieldStr  string `json: "field Str"`
	FieldInt  int    `json: "field Int"`
	FieldBool bool   `json: "field Bool"`
}

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	fmt.Println("clientOption TYPE", reflect.TypeOf(clientOptions), "\n")

	client, err := mongo.Connect(context.TODO(),clientOptions)
	if err != nil {
		fmt.Println("Mongo.connect() ERROR: ", err)
		os.Exit(1)
	}
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	col := client.Database("First_Database").Collection("First Collection")
	fmt.Println("Collection Type: ", reflect.TypeOf(col), "\n")

	oneDoc := MongoField{
		FieldStr: "This is first data and its very important",
		FieldInt: 826482746,
		FieldBool: true,
	}

	fmt.Println("oneDoc Type: ", reflect.TypeOf(oneDoc), "\n")

	result, insertErr := col.InsertOne(ctx, oneDoc)
	if insertErr != nil {
		fmt.Println("InsertONE Error:", insertErr)
		os.Exit(1)
	}else {
		fmt.Println("InsertOne() result type: ", reflect.TypeOf(result))
		fmt.Println("InsertOne() api result type: ", result)

		newID := result.InsertedID
		fmt.Println("InsertOne(), newID", newID)
		fmt.Println("InsertOne() newID", reflect.TypeOf(newID))
	}

}