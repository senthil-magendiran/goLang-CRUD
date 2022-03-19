package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)


type USER struct {
	Name string `json:name`
	Age  int    `json:age`
	City string `json:city`
}

var dbCollection = db().Database("goLangDB").Collection("records")

func createUser(w http.ResponseWriter, r *http.Request) {

	var user USER
	err := json.NewDecoder(r.Body).Decode(&user) // storing in person variable of type user
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(user.Name)
	dbCollection.InsertOne(context.TODO(), user)

	json.NewEncoder(w).Encode("User added successfully")

}

func getUsers(w http.ResponseWriter, r *http.Request) {
	var results []primitive.M
	cur, err := dbCollection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		fmt.Println(err)
	}
	for cur.Next(context.TODO()) {
		var elem primitive.M
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem)
	}

	json.NewEncoder(w).Encode(results)
}

func searchUser(w http.ResponseWriter, r *http.Request) {

	var body USER
	e := json.NewDecoder(r.Body).Decode(&body)
	if e != nil {
		fmt.Print(e)
	}
	var result primitive.M
	err := dbCollection.FindOne(context.TODO(), bson.D{{"name", body.Name}}).Decode(&result)
	if err != nil {

		fmt.Println(err)

	}
	json.NewEncoder(w).Encode(result)

}

func updateUser(w http.ResponseWriter, r *http.Request) {

	type updateBody struct {
		Name string `json:"name"` //value that has to be matched
		City string `json:"city"` // value that has to be modified
	}
	var body updateBody
	e := json.NewDecoder(r.Body).Decode(&body)
	if e != nil {

		fmt.Print(e)
	}
	filter := bson.D{{"name", body.Name}}
	after := options.After
	returnOpt := options.FindOneAndUpdateOptions{

		ReturnDocument: &after,
	}
	update := bson.D{{"$set", bson.D{{"city", body.City}}}}
	updateResult := dbCollection.FindOneAndUpdate(context.TODO(), filter, update, &returnOpt)

	var result primitive.M
	_ = updateResult.Decode(&result)

	json.NewEncoder(w).Encode(result)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)["name"]

	 dbCollection.DeleteOne(context.TODO(), bson.D{{"name", params}})

	json.NewEncoder(w).Encode("User deleted successfully")

}
