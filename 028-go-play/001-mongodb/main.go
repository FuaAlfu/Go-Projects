package main

import (
"context"
"fmt"

"go.mongodb.org/mongo-driver/mongo"
"go.mongodb.org/mongo-driver/mongo/options"

)

/*
a testing code
*/

func main() {

// Set client options for host and port of server
clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

// Connect to MongoDB server with client options set
client, err := mongo.Connect(context.TODO(), clientOptions)

if err != nil { //Error handling if connection fails
fmt.Println(err)
return // break out of function, do not continue without mongodb connection successfull }

err = client.Ping(context.TODO(), nil) // Ensure MongoDB server is reachable by pinging it

if err != nil { // Error handling if ping fails }
fmt.Println(err) return 
}

fmt.Println("Connected to MongoDB!") // Connection successful!

}
}