package database

import(
	"context"
	"log"
	"time"
)

var connectionString string = ""

type(
	DB struct{
		client *mongo.Client
	}
)