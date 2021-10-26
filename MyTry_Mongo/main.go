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
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func dbtry() *mongo.Client {
	// Create a Client and execute a ListDatabases operation.

	client, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		//log.Fatal(err)
		log.Fatal("Data log")
		fmt.Println("Program finished")
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	return client
}

type Employee struct {
	Name        string `json:"Employee_Name"`
	EmpId       int    `json:"Employee_Id"`
	Designation string `json:"Designation"`
}

var userdb = dbtry().Database("employee").Collection("gotest")

func PostEmp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var emp Employee

	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		fmt.Print(err)
	}
	insertemp, err := userdb.InsertOne(context.TODO(), emp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted Single Employee Record", insertemp)
	json.NewEncoder(w).Encode(insertemp.InsertedID)
}

func GetEmp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var empget Employee

	eg := json.NewDecoder(r.Body).Decode(&empget)
	if eg != nil {
		fmt.Println(eg)
	}
	var result primitive.M
	err := userdb.FindOne(context.TODO(), bson.D{primitive.E{Key: "name", Value: empget.Name}}).Decode(&result)
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(result)
}
func GetAllEmp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var allemp []primitive.M
	cursor, err := userdb.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(context.TODO()) {
		var all primitive.M
		err := cursor.Decode(&all)
		if err != nil {
			log.Fatal(err)
		}
		allemp = append(allemp, all)
	}
	json.NewEncoder(w).Encode(allemp)
}
func DelEmp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	urlid := mux.Vars(r)["id"]
	_id, err := primitive.ObjectIDFromHex(urlid)
	if err != nil {
		fmt.Printf(err.Error())
	}
	opt := options.Delete().SetCollation(&options.Collation{})
	res, err := userdb.DeleteOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: _id}}, opt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents \n", res.DeletedCount)
	json.NewEncoder(w).Encode(res.DeletedCount)
}

func UpdateEmp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	type updateEmployee struct {
		EmpId       int    `json:"Employee_Id"`
		Designation string `json:"Designation"`
	}
	// var putemp updateEmployee
	// pemp := json.NewDecoder(r.Body).Decode(&putemp)
	// if pemp != nil {
	// 	fmt.Print(pemp)
	// }
	// putfilter := bson.D{primitive.E{Key: "Employee_Id", Value: putemp.EmpId}}
	// after := options.After
	// returnop := options.FindOneAndUpdateOptions{
	// 	ReturnDocument: &after,
	// }
	// update := bson.D{primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "Designation", Value: putemp.Designation}}}}
	// updateResult := userdb.FindOneAndUpdate(context.TODO(), putfilter, update, &returnop)

	// var result primitive.M
	// _ = updateResult.Decode(&result)
	// json.NewEncoder(w).Encode(result)
	var body updateEmployee
	e := json.NewDecoder(r.Body).Decode(&body)
	if e != nil {

		fmt.Print(e)
	}
	filter := bson.D{primitive.E{Key: "Employee_Id", Value: body.EmpId}} // converting value to BSON type
	after := options.After                                               // for returning updated document
	returnOpt := options.FindOneAndUpdateOptions{

		ReturnDocument: &after,
	}
	update := bson.D{primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "Designation", Value: body.Designation}}}}
	// update := bson.D{{"$set", bson.D{{"city", body.City}}}}
	updateResult := userdb.FindOneAndUpdate(context.TODO(), filter, update, &returnOpt)

	var result primitive.M
	_ = updateResult.Decode(&result)

	json.NewEncoder(w).Encode(result)
}
func main() {
	apiroute := mux.NewRouter()
	basepath := apiroute.PathPrefix("/api").Subrouter()
	basepath.HandleFunc("/postemp", PostEmp).Methods(http.MethodPost)
	basepath.HandleFunc("/getemp", GetEmp).Methods(http.MethodGet)
	basepath.HandleFunc("/getallemp", GetAllEmp).Methods(http.MethodGet)
	basepath.HandleFunc("/delemp/{id}", DelEmp).Methods(http.MethodDelete)
	basepath.HandleFunc("/putemp", UpdateEmp).Methods(http.MethodPut)
	log.Fatal(http.ListenAndServe("localhost:8080", basepath))
}
