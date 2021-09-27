package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Database interface which will be used in handlers,
//Only user can access through this interface
type Database interface {
	Connect()
	Create()
	Read() []map[string]interface{}
	Update()
	Delete()
}

//MongoDB struct that implements Database interface
type MongoDB struct {
	ip       string
	port     int
	user     string
	password string
}

//Connect method of MongoDB
func (m *MongoDB) Connect() {
	credential := options.Credential{
		Username: m.user,
		Password: m.password,
	}
	uri := fmt.Sprintf("mongodb://%s:%v", m.ip, m.port)
	clientOptions := options.Client().ApplyURI(uri).SetAuth(credential)
	client, err := mongo.Connect(context.TODO(), clientOptions) //Use context.TODO if you don't know what context has to be delivered
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil) //Check connection
	if err != nil {
		log.Fatal(err)
	}
}

//Create function of MongoDB
func (m *MongoDB) Create() {

}

//Read function of MongoDB
func (m *MongoDB) Read() {

}

//Update function of MongoDB
func (m *MongoDB) Update() {

}

//Delete function of MongoDB
func (m *MongoDB) Delete() {

}
