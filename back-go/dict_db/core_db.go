package dict_db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

type DictDBConnection struct {
	DB *mongo.Database
}

type ConnectionWithUser struct {
	DictDBConnection
	UsersCollection *mongo.Collection
	User            UserSchema
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

func GetDBWithUser(userTgId string) (*ConnectionWithUser, error) {
	db := GetDB()

	filter := bson.M{"tgID": userTgId}

	var userDoc UserSchema

	usersCollection := db.DB.Collection("users")
	err := usersCollection.FindOne(context.Background(), filter).Decode(&userDoc)

	if err != nil {
		// failed to find user, then create a new one
		var newUser UserSchema

		newUser.TgID = userTgId

		_, err := usersCollection.InsertOne(context.Background(), newUser)

		if err != nil {
			return nil, err
		}

		err = usersCollection.FindOne(context.Background(), filter).Decode(&userDoc)

		if err != nil {
			return nil, err
		}
	}

	return &ConnectionWithUser{
		DictDBConnection: db,
		UsersCollection:  usersCollection,
		User:             userDoc,
	}, nil
}
