package service

import "book-rest-api-showcase/model"

type BookService interface {
	CreateBook(in model.Book) (res model.Book, err error)
	GetAllBooks() (res []model.Book, err error)
	GetBookByID(id int64) (res model.Book, err error)
	UpdateBook(in model.Book) (res model.Book, err error)
	DeleteBook(id int64) (err error)
}

func (s *Service) CreateBook(in model.Book) (res model.Book, err error) {
	return s.repo.CreateBook(in)
}

func (s *Service) GetAllBooks() (res []model.Book, err error) {
	res, err = s.repo.GetAllBooks()
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *Service) GetBookByID(id int64) (res model.Book, err error) {
	return s.repo.GetBookByID(id)
}

func (s *Service) UpdateBook(in model.Book) (res model.Book, err error) {
	return s.repo.UpdateBook(in)
}
func (s *Service) DeleteBook(id int64) (err error) {
	return s.repo.DeleteBook(id)
}
