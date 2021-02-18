package interactor

import (
	"book-api/model"
	"book-api/usecase/output"
	"book-api/usecase/presenter"
	"book-api/usecase/repository"
)

type bookInteractor struct {
	BookRepository repository.BookRepository
	BookPresenter  presenter.BookPresenter
}

type BookInteractor interface {
	Get(b []*model.Book) ([]*output.BookOutput, error)
}

func NewBookInteractor(r repository.BookRepository, p presenter.BookPresenter) BookInteractor {
	return &bookInteractor{
		BookRepository: r,
		BookPresenter:  p,
	}
}

func (bi *bookInteractor) Get(b []*model.Book) ([]*output.BookOutput, error) {
	b, err := bi.BookRepository.FindAll(b)
	if err != nil {
		return nil, err
	}
	return bi.BookPresenter.PresentBooks(b), nil
}
