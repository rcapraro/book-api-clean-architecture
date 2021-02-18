package repository

import (
	"github.com/dgraph-io/badger/v3"
	"log"
)

func NewDB(dir string) *badger.DB {
	opts := badger.DefaultOptions("./badger")
	opts.Dir = dir
	opts.ValueDir = dir
	opts.Logger = nil

	db, err := badger.Open(opts)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
