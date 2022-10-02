package firego

import (
	"context"
	"errors"
	"os"
	"path/filepath"

	firestore "cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	log "github.com/sirupsen/logrus"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

var CREDENTIALS_FILE_PATH = os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")

//var client *firestore.Client = connect("./PRIVATE.json")
var client *firestore.Client = connect()

func connect() *firestore.Client {
	if CREDENTIALS_FILE_PATH == "" {
		log.Info("trying local credentials file")
		credentials_file_temp, err := filepath.Abs("./credentials/PRIVATE.json")
		if err != nil {
			log.Error("failed to locate file")
		}
		CREDENTIALS_FILE_PATH = credentials_file_temp
		log.Debug("local credentials file: ", CREDENTIALS_FILE_PATH)
	}

	sa := option.WithCredentialsFile(CREDENTIALS_FILE_PATH)
	app, err := firebase.NewApp(context.Background(), nil, sa)
	if err != nil {
		log.Error("failed to create firebase client app")
		log.Fatal(err)
	}
	client, err := app.Firestore(context.Background())
	if err != nil {
		log.Error("failed to create firestore client")
		log.Fatal(err)
	}
	return client
}

func Disconnect() {
	defer client.Close()
}

func Create(collection string, id string, doc map[string]interface{}) error {
	result, err := client.Collection(collection).Doc(id).Set(context.Background(), doc)
	if err != nil {
		log.Error("failed to create database document: ", doc, err)
		return err
	}
	log.Info("Successfully created database document: ", doc)
	log.Debug("firestore result response: ", result)
	return nil
}

func Update(collection string, id string, doc map[string]interface{}) error {
	result, err := client.Collection(collection).Doc(id).Set(context.Background(), doc)
	if err != nil {
		return err
	}
	log.Info("Successfully updated database document: ", doc)
	log.Debug("firestore result response: ", result)
	return nil
}

func Delete(collection string, id string) error {
	result, err := client.Collection(collection).Doc(id).Delete(context.Background())
	if err != nil {
		return err
	}
	log.Info("Successfully deleted database document: ", id)
	log.Debug("firestore result response: ", result)
	return nil
}

func Get(collection string, id string) (map[string]interface{}, error) {
	result, err := client.Collection(collection).Doc(id).Get(context.Background())
	if err != nil {
		return nil, errors.New("failed to get document")
	}
	log.Info("Successfully got database document: ", id)
	log.Debug("firestore result response: ", result)
	return result.Data(), nil
}

func ListAll(collection string) (*[]map[string]interface{}, error) {
	results := []map[string]interface{}{}
	iter := client.Collection(collection).Documents(context.Background())
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Error("failed to get document in list", err)
			continue
		}

		log.Debug("successfully added document to list: ", doc.Data())
		results = append(results, doc.Data())
	}
	log.Info("successfully got document list: ", results)
	return &results, nil
}
