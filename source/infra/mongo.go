package infra

import (
	"context"
	"doYourLogin/source/configuration"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type MongoDB struct {
	Client     *mongo.Client
	Database   *mongo.Database
	Collection *mongo.Collection
}

func NewMongoDB(connectionString, databaseName, collectionName, username, password string) (*MongoDB, error) {
	clientOptions := options.Client().ApplyURI(connectionString)
	clientOptions.Auth = &options.Credential{
		Username: username,
		Password: password,
	}

	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		return nil, err
	}

	db := client.Database(databaseName)
	coll := db.Collection(collectionName)

	return &MongoDB{
		Client:     client,
		Database:   db,
		Collection: coll,
	}, nil

}

func (m *MongoDB) Close() {
	err := m.Client.Disconnect(context.Background())
	if err != nil {
		log.Println("Error disconnection from MongoDB", err)
	}
}

func (m *MongoDB) InsertDocument(document interface{}) error {
	_, err := m.Collection.InsertOne(context.Background(), document)
	return err
}

func (m *MongoDB) FindDocuments(filter interface{}) ([]interface{}, error) {
	cur, err := m.Collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	defer cur.Close(context.Background())

	var results []interface{}
	for cur.Next(context.Background()) {
		var doc interface{}
		if err := cur.Decode(&doc); err != nil {
			log.Println("Error, decoding document", err)
			continue
		}
		results = append(results, doc)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return nil, err
}

func InitMongo() {

	connectionString := configuration.MONGO_CONNECTION_STRING.ValueAsString()
	databaseName := configuration.MONGO_DATABASE_NAME.ValueAsString()
	collectionName := configuration.MONGO_COLLECTION_NAME.ValueAsString()
	username := configuration.MONGO_DATABASE_NAME.ValueAsString()
	password := configuration.MONGO_PASSWORD.ValueAsString()

	db, err := NewMongoDB(connectionString, databaseName, collectionName, username, password)

	if err != nil {
		log.Fatal("Error connection to MongoDB: ", err)
	}

	defer db.Close()
	//document := map[string]interface{}{}
	//
	//err = db.InsertDocument(document)
	//if err != nil {
	//	log.Println("Error inserting document", err)
	//}
	//
	//filter := map[string]interface{}{
	//	"name": "John Doe",
	//}
	//
	//results, err := db.FindDocuments(filter)
	//
	//if err != nil {
	//	log.Println("Error finding documents", err)
	//}
	//
	//fmt.Println("Results", results)
}
