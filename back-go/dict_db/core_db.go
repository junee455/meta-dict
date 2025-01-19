package dict_db

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DictDBConnection struct {
	DB *mongo.Database
}

var DB_CONNECTION DictDBConnection

func InitDB() {
	uri := os.Getenv("MONGO_URI")

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)

	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	// defer func() {
	// 	if err := client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()

	DB_CONNECTION = DictDBConnection{DB: client.Database("meta_dict_local")}
}

func Disconnect() {
	DB_CONNECTION.DB.Client().Disconnect(context.TODO())
}

func GetDB() DictDBConnection {
	return DB_CONNECTION
}
