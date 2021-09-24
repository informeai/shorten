package store

import (
	"context"
	"github.com/informeai/shorten/entities"
	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

//uri for connection database
const uri = "mongodb://localhost:27017/"

//StoreMongodb is struct for db mongodb
type StoreMongodb struct {
	client mongo.Client
	ctx    context.Context
}

//NewStoreMongodb return instance the StoreMongodb
func NewStoreMongodb() *StoreMongodb {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	c, err := mongo.Connect(options.Client().ApllyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	return &StoreMongodb{client: c, ctx: ctx}
}

//Ping execute test of connection in database.
func (sm *StoreMongodb) Ping() error {
	if err := client.Ping(sm.ctx, readpref.Primary()); err != nil {
		return err
	}
	return nil
}

//Get return first Shorten from database
func (sm *StoreMongodb) Get(id string) (entities.Shorten, error) {
	return entities.Shorten{}, nil
}

//Insert add new Shorten to database
func (sm *StoreMongodb) Insert(srt entities.Shorten) error {
	return nil
}

//Update change the shorten and save in database.
func (sm *StoreMongodb) Update(srt entities.Shorten) error {
	return nil
}

//Disconnect execute disconnect for database.
func (sm *StoreMongodb) Disconnect() error {
	if err := sm.client.Disconnect(sm.ctx); err != nil {
		return err
	}
	return nil
}
