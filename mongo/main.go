package main

import (
	"context"
	"flag"
	"log"
	"strings"
	"time"

	"github.com/GerryLon/learn-go/assert"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type ST struct {
	Name string `bson:"name"`
}


func main() {
	hosts := ""
	username := ""
	password := ""
	database := ""
	flag.StringVar(&hosts,"hosts", "", "hosts,多个用半角逗号分隔")
	flag.StringVar(&username,"username", "", "username")
	flag.StringVar(&password,"password", "", "password")
	flag.StringVar(&database,"database", "", "database")
	flag.Parse()

	logrus.SetReportCaller(true)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	clientOptions := &options.ClientOptions{
		Hosts: strings.Split(hosts, ","),
		Auth: &options.Credential{
			AuthMechanism: "SCRAM-SHA-1",
			AuthSource:    "admin",
			Username:      username,
			Password:      password,
		},
	}
	client, err := mongo.Connect(
		ctx,
		clientOptions,
	)
	assert.Nil(err)

	err = client.Ping(ctx, readpref.Primary())
	assert.Nil(err)

	db := client.Database(database)

	limit := int64(3)
	collection := db.Collection("role")

	// findOptions.SetSkip(0)
	// findOptions.SetSort(bson.M{"_id": -1})
	cur, err := collection.Find(ctx, bson.M{}, options.Find().SetLimit(limit))

	assert.Nil(err)
	defer cur.Close(ctx)

	sts := make([]ST, 0)
	// err = cur.All(ctx, &sts)
	// assert.Nil(err)
	for cur.Next(ctx) {
		var result ST
		if err := cur.Decode(&result); err != nil {
			log.Fatalln("Error getting record:", err)
		}
		sts = append(sts, result)
		// For some reason Decode doesn't work for the _id field. Extract separately.
		// result.ID = cur.Current.Lookup("_id").ObjectID().String()
	}
	if err := cur.Err(); err != nil {
		logrus.Panic(err)
	}

	logrus.Infof("len(sts)=%d", len(sts))
	if len(sts) > 0 {
		logrus.Infof("first ele: %+v", sts[0])
	}
}
