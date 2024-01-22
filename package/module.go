package _package

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllCatatan(db *mongo.Database, col string) (data []Catatan) {
	catat := db.Collection(col)
	filter := bson.M{}
	cursor, err := catat.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func InsertCatatan(db *mongo.Database, col string, catatan Catatan) (insertedID primitive.ObjectID, err error) {
	result, err := db.Collection(col).InsertOne(context.Background(), catatan)
	if err != nil {
		fmt.Printf("Insert Catatan: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func GetLastCatatan(db *mongo.Database, col string) (lastCatatan Catatan, err error) {
	// Membuat filter untuk mendapatkan data terakhir
	filter := bson.M{}
	options := options.Find().SetSort(bson.D{{"id", -1}}).SetLimit(1)

	// Mengeksekusi query
	cursor, err := db.Collection(col).Find(context.TODO(), filter, options)
	if err != nil {
		return lastCatatan, err
	}
	defer cursor.Close(context.TODO())

	// Mengambil data terakhir
	if cursor.Next(context.TODO()) {
		err := cursor.Decode(&lastCatatan)
		if err != nil {
			return lastCatatan, err
		}
	}

	return lastCatatan, nil
}

func UpdateCatatan(db *mongo.Database, col string, id int, updatedCatatan Catatan) (err error) {
	filter := bson.M{"id": id}
	update := bson.M{
		"$set": bson.M{
			// "id":          updatedCatatan.ID2,
			"title":       updatedCatatan.Title,
			"note":        updatedCatatan.Note,
			"date":        updatedCatatan.Date,
			"startTime":   updatedCatatan.StartTime,
			"endTime":     updatedCatatan.EndTime,
			"remind":      updatedCatatan.Remind,
			"repeat":      updatedCatatan.Repeat,
			"isCompleted": updatedCatatan.IsCompleted,
			"completedAt": updatedCatatan.CompletedAt,
			"createdAt":   updatedCatatan.CreatedAt,
			"updatedAt":   updatedCatatan.UpdatedAt,
			"color":       updatedCatatan.Color,
		},
	}
	fmt.Printf("Filter: %v\n", filter)
	fmt.Printf("Update: %v\n", update)
	result, err := db.Collection(col).UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Printf("UpdateCatatan: %v\n", err)
		return
	}
	if result.ModifiedCount == 0 {
		err = errors.New("No data has been changed with the specified ID")
		return
	}
	return nil
}

func UpdateStatus(db *mongo.Database, col string, id int, updatedCatatan Catatan) (err error) {
	filter := bson.M{"id": id}
	update := bson.M{
		"$set": bson.M{
			"isCompleted": updatedCatatan.IsCompleted,
			"completedAt": updatedCatatan.CompletedAt,
		},
	}
	fmt.Printf("Filter: %v\n", filter)
	fmt.Printf("Update: %v\n", update)
	result, err := db.Collection(col).UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Printf("UpdateStatus: %v\n", err)
		return
	}
	if result.ModifiedCount == 0 {
		err = errors.New("No data has been changed with the specified ID")
		return
	}
	return nil
}

func DeleteCatatan(db *mongo.Database, collectionName string, id int) error {
	kontak := db.Collection(collectionName)
	filter := bson.M{"id": id}

	result, err := kontak.DeleteOne(context.Background(), filter)
	if err != nil {
		return fmt.Errorf("gagal menghapus data dengan ID %d: %s", id, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data dengan ID %d tidak ditemukan", id)
	}

	return nil
}
