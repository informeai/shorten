package store

import (
	"github.com/informeai/shorten/entities"
)

//Shortener is interface for methods the store package.
type Shortener interface {
	Get(id string) (entities.Shorten, error)
	Insert(srt entities.Shorten) error
	Update(srt entities.Shorten) error
}
