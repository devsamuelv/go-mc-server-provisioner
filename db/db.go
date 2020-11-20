package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"main/types"
)

// ListDocs Lists a Doc in a collection
func ListDocs(ctx context.Context, client *mongo.Client, collection string) {
	var docs []bson.D
	// var returnDocs []string

	cusor, err := client.Database("global").Collection(collection).Find(ctx, bson.M{})

	if err != nil {
		backend.PrintError(err.Error())
		return
	}

	if err = cusor.All(ctx, &docs); err != nil {
		backend.PrintError(err.Error())
		return
	}

	// return returnDocs

}

// AddDocument Adds a document to a collection and returns a true if compleated for false if failed
func AddDocument(ctx context.Context, client *mongo.Client, collection string, data bson.D) bool {
	_, err := client.Database("global").Collection(collection).InsertOne(ctx, data)

	if err != nil {
		backend.PrintError(err.Error())

		return false
	}

	return true
}

// AddDocumentServer Adds a document to a collection and returns a true if compleated for false if failed
func AddDocumentServer(ctx context.Context, client *mongo.Client, collection string, data bson.D) bool {
	_, err := client.Database("server").Collection(collection).InsertOne(ctx, data)

	if err != nil {
		backend.PrintError(err.Error())

		return false
	}

	return true
}

// GetDocs returns a list of docs in json
func GetDocs(ctx context.Context, client *mongo.Client, collection string) []bson.M {
	cursor, err := client.Database("global").Collection(collection).Find(ctx, bson.M{})

	if err != nil {
		log.Fatal(err.Error())
	}

	var docs []bson.M

	// FIXME this offers less memory usage but has the error cannot decode document into []primitive.M
	// defer cursor.Close(ctx)
	// for cursor.Next(ctx) {
	// 	if err = cursor.Decode(&docs); err != nil {
	// 		log.Fatal("test" + err.Error())
	// 	}
	// }

	if err = cursor.All(ctx, &docs); err != nil {
		log.Fatal(err.Error())
	}

	return docs
}

// GetDocsServer returns a list of docs in json
func GetDocsServer(ctx context.Context, client *mongo.Client, collection string) []bson.M {
	cursor, err := client.Database("server").Collection(collection).Find(ctx, bson.M{})

	if err != nil {
		log.Fatal(err.Error())
	}

	var docs []bson.M

	// FIXME this offers less memory usage but has the error cannot decode document into []primitive.M
	// defer cursor.Close(ctx)
	// for cursor.Next(ctx) {
	// 	if err = cursor.Decode(&docs); err != nil {
	// 		log.Fatal("test" + err.Error())
	// 	}
	// }

	if err = cursor.All(ctx, &docs); err != nil {
		log.Fatal(err.Error())
	}

	return docs
}

// UpdateDocument Updates a doc with the id spcifyed
func UpdateDocument(ctx context.Context, client *mongo.Client, collection string, id string, data interface{}) bool {

	filter := backend.UpdateFilter{
		Id: id,
	}

	_, err := client.Database("global").Collection(collection).UpdateOne(ctx, filter, data)

	if err != nil {
		return false
	}

	return true
}

// UpdateDocumentByType test update function
func UpdateDocumentByType(ctx context.Context, client *mongo.Client, collection string, filter bson.M, data bson.M) bool {
	parsedData := bson.M{
		"$set": data,
	}
	_, err := client.Database("global").Collection(collection).UpdateOne(ctx, filter, parsedData)

	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	return true
}

// FindDocByType finds and returns a doc by a filter
func FindDocByType(ctx context.Context, client *mongo.Client, collection string, filter bson.M) []bson.M {
	// data, err := client.Database("global").Collection("user").Find(ctx, bson.M{"uid": parms.Uid})

	cursor, err := client.Database("global").Collection(collection).Find(ctx, filter)

	if err != nil {
		log.Fatal(err.Error())
	}

	var docs []bson.M

	// FIXME this offers less memory usage but has the error cannot decode document into []primitive.M
	// defer cursor.Close(ctx)
	// for cursor.Next(ctx) {
	// 	if err = cursor.Decode(&docs); err != nil {
	// 		log.Fatal("test" + err.Error())
	// 	}
	// }

	if err = cursor.All(ctx, &docs); err != nil {
		log.Fatal(err.Error())
	}

	return docs
}

// FindDocByTypeContiains finds and returns a doc by a filter
func FindDocByTypeContiains(ctx context.Context, client *mongo.Client, collection string, filter bson.D) []bson.M {
	// data, err := client.Database("global").Collection("user").Find(ctx, bson.M{"uid": parms.Uid})

	cursor, err := client.Database("global").Collection(collection).Find(ctx, filter)

	if err != nil {
		log.Fatal(err.Error())
	}

	var docs []bson.M

	// FIXME this offers less memory usage but has the error cannot decode document into []primitive.M
	// defer cursor.Close(ctx)
	// for cursor.Next(ctx) {
	// 	if err = cursor.Decode(&docs); err != nil {
	// 		log.Fatal("test" + err.Error())
	// 	}
	// }

	if err = cursor.All(ctx, &docs); err != nil {
		log.Fatal(err.Error())
	}

	return docs
}

// DeleteDoc Deletes a Doc in a collection
func DeleteDoc(ctx context.Context, client *mongo.Client, collection string, filter bson.D) bool {
	_, err := client.Database("global").Collection(collection).DeleteOne(ctx, filter)

	if err != nil {
		return false
	}

	return true
}

// DeleteDocServer Deletes a Doc in a collection
func DeleteDocServer(ctx context.Context, client *mongo.Client, collection string, filter bson.M) bool {
	_, err := client.Database("server").Collection(collection).DeleteOne(ctx, filter)

	if err != nil {
		return false
	}

	return true
}

// AddIndex add a value to a index
func AddIndex(ctx context.Context, client *mongo.Client, collection string, data mongo.IndexModel) string {
	returnData, err := client.Database("global").Collection(collection).Indexes().CreateOne(ctx, data)

	if err != nil {
		fmt.Println(err.Error())
	}

	return returnData
}

// FindIndex find value in an index
func FindIndex(ctx context.Context, client *mongo.Client, collection string, filter backend.SearchIndex) bool {
	cursor, err := client.Database("global").Collection(collection).Indexes().List(ctx)

	var data []bson.M
	var returnData bool = false

	if err != nil {
		fmt.Println(err.Error())
	}

	if err = cursor.All(ctx, &data); err != nil {
		log.Fatal(err.Error())
	}

	for _, doc := range data {
		fmt.Println(doc)

		if doc[filter.FilterType] == filter.Content {
			returnData = true
		}
	}

	return returnData
}
