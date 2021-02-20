package interactor

import (
	"book-api/usecase/input"
	"book-api/usecase/output"
	"book-api/usecase/presenter"
	"book-api/usecase/repository"
)

type bookInteractor struct {
	BookRepository repository.BookRepository
	BookPresenter  presenter.BookPresenter
}

type BookInteractor interface {
	FindAll() ([]*output.BookOutput, error)
	Create(b *input.BookInput) error
}

func NewBookInteractor(r repository.BookRepository, p presenter.BookPresenter) BookInteractor {
	return &bookInteractor{
		BookRepository: r,
		BookPresenter:  p,
	}
}

func (bi *bookInteractor) FindAll() ([]*output.BookOutput, error) {
	b, err := bi.BookRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return bi.BookPresenter.PresentBooks(b), nil
}

func (bi *bookInteractor) Create(b *input.BookInput) error {
	err := bi.BookRepository.Save(b.ToModel())
	return err
}
