package registry

import (
	"book-api/infrastructure/controller"
	ip "book-api/infrastructure/presenter"
	ir "book-api/infrastructure/repository"
	"book-api/usecase/interactor"
	"book-api/usecase/presenter"
	"github.com/dgraph-io/badger/v3"
)

type registry struct {
	db *badger.DB
}

type Registry interface {
	NewBookController() controller.BookController
}

func NewRegistry(db *badger.DB) Registry {
	return &registry{db}
}

func (r *registry) NewBookController() controller.BookController {
	return controller.NewBookController(r.NewBookInteractor())
}

func (r *registry) NewBookInteractor() interactor.BookInteractor {
	return interactor.NewBookInteractor(ir.NewBookRepository(r.db), r.NewBookPresenter())
}

func (r *registry) NewBookPresenter() presenter.BookPresenter {
	return ip.NewBookPresenter()
}
