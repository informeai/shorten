package store

import (
	"context"
	"github.com/informeai/shorten/config"
	"github.com/informeai/shorten/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
	"time"
)

//StoreMongodb is struct for db mongodb
type StoreMongodb struct {
	client     *mongo.Client
	ctx        context.Context
	collection *mongo.Collection
}

//NewStoreMongodb return instance the StoreMongodb
func NewStoreMongodb() *StoreMongodb {
	config.LoadEnvs()
	mongoUri := os.Getenv("MONGO_URI")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	c, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUri))
	if err != nil {
		log.Fatal(err)
	}
	coll := c.Database("shorten").Collection("shortlinks")
	return &StoreMongodb{client: c, ctx: ctx, collection: coll}
}

//Ping execute test of connection in database.
func (sm *StoreMongodb) Ping() error {
	if err := sm.client.Ping(sm.ctx, readpref.Primary()); err != nil {
		return err
	}
	return nil
}

//Get return first Shorten from database
func (sm *StoreMongodb) Get(id string) (entities.Shorten, error) {
	var shorten entities.Shorten
	if err := sm.collection.FindOne(sm.ctx, bson.D{{"key", id}}).Decode(&shorten); err != nil {
		return shorten, err
	}
	return shorten, nil
}

//Insert add new Shorten to database
func (sm *StoreMongodb) Insert(srt entities.Shorten) error {
	_, err := sm.collection.InsertOne(sm.ctx, srt)
	if err != nil {
		return err
	}
	err = sm.Disconnect()
	if err != nil {
		return err
	}
	return nil
}

//Update change the shorten and save in database.
func (sm *StoreMongodb) Update(srt entities.Shorten) error {
	filter := bson.D{{"key", srt.Id}}
	update := bson.D{{"$set", bson.D{{"url", srt.Url}, {"visits", srt.Visits + 1}}}}
	_, err := sm.collection.UpdateOne(sm.ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

//Disconnect execute disconnect for database.
func (sm *StoreMongodb) Disconnect() error {
	if err := sm.client.Disconnect(sm.ctx); err != nil {
		return err
	}
	return nil
}
