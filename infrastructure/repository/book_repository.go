package repository

import (
	"book-api/model"
	"book-api/usecase/repository"
	"bytes"
	"encoding/gob"
	"github.com/dgraph-io/badger/v3"
)

type bookRepository struct {
	db *badger.DB
}

func NewBookRepository(db *badger.DB) repository.BookRepository {
	return &bookRepository{db}
}

func (br *bookRepository) FindAll() ([]*model.Book, error) {

	var books []*model.Book

	err := br.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			book, err := getBookEntry(it.Item())
			if err != nil {
				return err
			}
			books = append(books, book)
		}
		return nil
	})
	return books, err
}

func (br *bookRepository) Save(b *model.Book) error {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(b)
	if err != nil {
		return err
	}

	return br.db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(b.ISBN), buffer.Bytes())
	})
}

func getBookEntry(item *badger.Item) (*model.Book, error) {
	var bookEntry model.Book
	var buffer bytes.Buffer

	err := item.Value(func(val []byte) error {
		_, err := buffer.Write(val)
		return err
	})

	decoder := gob.NewDecoder(&buffer)
	err = decoder.Decode(&bookEntry)

	return &bookEntry, err
}
