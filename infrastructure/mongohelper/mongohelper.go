package mongohelper

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect connection to mongodb
func Connect() (*mongo.Database, error) {
	fmt.Println("init mongodb")
	dbname := viper.GetString("MONGO_DB")

	// connection string
	uri := "mongodb://"
	if viper.GetBool("MONGO_AUTH") {
		uri += viper.GetString("MONGO_USER") + ":" + viper.GetString("MONGO_PASSWORD") + "@"
	}
	uri += viper.GetString("MONGO_HOST") + ":" + viper.GetString("MONGO_PORT")

	replicaSet := viper.GetString("MONGO_REPLICASET")
	if replicaSet != "" {
		uri += "/" + dbname + "?replicaSet=" + replicaSet // + "&authSource=admin"
	}

	fmt.Println("uri", uri)

	// create connection
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("err")
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return client.Database(dbname), nil
}
