package controllers

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"optician-rest-api/database"
	"optician-rest-api/models"
	"time"
)

func CreateBillData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var bill models.Bills

	err := json.NewDecoder(r.Body).Decode(&bill)
	if err != nil {
		return
	}

	// Assign curernt date and time to the bill
	bill.Date = time.Now().Format(time.DateTime)

	collection := database.GetBillsCollection()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := collection.InsertOne(ctx, bill)

	err1 := json.NewEncoder(w).Encode(result)
	if err1 != nil {
		return
	}
}

func GetBillsData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var bills []models.Bills
	collection := database.GetBillsCollection()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(`{"message": "` + err.Error() + `"}`))
		if err != nil {
			return
		}
		return
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			log.Fatal(err)
			return
		}
	}(cursor, ctx)
	for cursor.Next(ctx) {
		var bill models.Bills
		err := cursor.Decode(&bill)
		if err != nil {
			return
		}
		bills = append(bills, bill)
	}
	if err := cursor.Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err1 := w.Write([]byte(`{"message": "` + err.Error() + `"}`))
		if err1 != nil {
			log.Fatal(err1)
			return
		}
		return
	}

	err2 := json.NewEncoder(w).Encode(bills)
	if err2 != nil {
		log.Fatal(err2)
		return
	}
}

func GetBillData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["imageId"])
	var bill models.Bills
	collection := database.GetBillsCollection()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := collection.FindOne(ctx, models.Bills{Id: id}).Decode(&bill)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err1 := w.Write([]byte(`{"message": "` + err.Error() + `"}`))
		if err1 != nil {
			log.Fatal(err1)
			return
		}
		return
	}

	err2 := json.NewEncoder(w).Encode(bill)
	if err2 != nil {
		log.Fatal(err2)
		return
	}
}

func UpdateBillData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["imageId"])
	var bill models.Bills
	collection := database.GetBillsCollection()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := collection.FindOne(ctx, models.Bills{Id: id}).Decode(&bill)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err1 := w.Write([]byte(`{"message": "` + err.Error() + `"}`))
		if err1 != nil {
			log.Fatal(err1)
			return
		}
		return
	}

	// Decode the request body into an ImageData struct
	err = json.NewDecoder(r.Body).Decode(&bill)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err2 := w.Write([]byte(`{"message": "Invalid request payload"}`))
		if err2 != nil {
			log.Fatal(err2)
			return
		}
		return
	}
	// Update the image data
	update := bson.M{
		"$set": bson.M{
			"date":         bill.Date,
			"customerbill": bill.SpectaclePrescription,
		},
	}
	err = collection.FindOneAndUpdate(ctx, bson.M{"_id": id}, update).Decode(&bill)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err3 := w.Write([]byte(`{"message": "` + err.Error() + `"}`))
		if err3 != nil {
			log.Fatal(err3)
			return
		}
		return
	}

	err4 := json.NewEncoder(w).Encode(bill)
	if err4 != nil {
		return
	}
}

func DeleteBillData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["billId"])
	collection := database.GetBillsCollection()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err1 := w.Write([]byte(`{"message": "` + err.Error() + `"}`))
		if err1 != nil {
			log.Fatal(err1)
			return
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err2 := w.Write([]byte(`{"message": "Bill delete successfully"}`))
	if err2 != nil {
		log.Fatal(err2)
		return
	}
}
