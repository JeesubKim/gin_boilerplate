package main

import "log"

func main() {
	//Connect with DB

	var database Database
	mongo := &MongoDB{ip: "localhost", port: 27017, user: "root", password: "admin"}
	database = mongo
	database.Connect()
	log.Println(database.Read("boilerplate", "user"))
	router := initRoutes()

	router.Run()
}
