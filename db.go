package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Database interface which will be used in handlers,
//Only user can access through this interface
type Database interface {
	Connect()
	Create()
	Read(database string, tableOrCollection string) []map[string]interface{}
	Update()
	Delete()
}

//MongoDB struct that implements Database interface
type MongoDB struct {
	ip           string
	port         int
	user         string
	password     string
	client       *mongo.Client
	databaseName string
}

//Connect method of MongoDB
func (m *MongoDB) Connect() {
	credential := options.Credential{
		Username: m.user,
		Password: m.password,
	}
	uri := fmt.Sprintf("mongodb://%s:%v", m.ip, m.port)
	clientOptions := options.Client().ApplyURI(uri).SetAuth(credential)
	// client, err := mongo.NewClient(clientOptions)
	client, err := mongo.Connect(context.TODO(), clientOptions) //Use context.TODO if you don't know what context has to be delivered
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil) //Check connection
	if err != nil {
		log.Fatal(err)
	}
	m.client = client
	log.Println("Connection complete!", client)

}

//Create function of MongoDB
func (m *MongoDB) Create() {

}

//Read function of MongoDB
func (m *MongoDB) Read(database string, tableOrCollection string) []map[string]interface{} {
	collection := m.client.Database(database).Collection(tableOrCollection)
	cur, err := collection.Find(context.TODO(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}
	var results []map[string]interface{}
	for cur.Next(context.TODO()) {
		var val interface{}
		err := cur.Decode(&val)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, map[string]interface{}{"key": val})
	}
	return results
}

//Update function of MongoDB
func (m *MongoDB) Update() {

}

//Delete function of MongoDB
func (m *MongoDB) Delete() {

}
