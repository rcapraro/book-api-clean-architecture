package controller

import "github.com/dgraph-io/badger/v3"

type registry struct {
	db *badger.DB
}

type Registry interface {
	NewAppController() AppController
}

func NewRegistry(db *badger.DB) Registry {
	return &registry{db}
}

func (r registry) NewAppController() AppController {
	return r.NewAppController()
}
