package repository

import (
	"book-rest-api-showcase/helper"
	"book-rest-api-showcase/model"
	"fmt"
)

type BookRepo interface {
	CreateBook(in model.Book) (res model.Book, err error)
	GetAllBooks() (res []model.Book, err error)
	GetBookByID(id int64) (res model.Book, err error)
	UpdateBook(in model.Book) (res model.Book, err error)
	DeleteBook(id int64) (err error)
}

func (r Repo) CreateBook(in model.Book) (res model.Book, err error) {
	err = r.gorm.Create(&in).Scan(&res).Error
	if err != nil {
		return res, err
	}
	return res, nil
}

func (r Repo) GetAllBooks() (res []model.Book, err error) {
	// err = r.gorm.Where("deleted_at is null").Find(&res).Error
	err = r.gorm.Find(&res).Error
	if err != nil {
		return res, err
	}
	return res, nil
}

func (r Repo) GetBookByID(id int64) (res model.Book, err error) {
	// err = r.gorm.Where("deleted_at is null").First(&res, id).Error
	err = r.gorm.First(&res, id).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) UpdateBook(in model.Book) (res model.Book, err error) {
	err = r.gorm.Model(&res).Where("id = ?", in.ID).Updates(model.Book{
		Namebook:  in.Namebook,
		Author:    in.Author,
		CreatedAt: in.CreatedAt,
		UpdatedAt: in.UpdatedAt,
	}).Scan(&res).Error
	if err != nil {
		return res, err
	}
	return res, nil
}

func (r Repo) DeleteBook(id int64) (err error) {
	book := model.Book{}
	result := r.gorm.Where("id = ?", id).Delete(&book)
	if result.RowsAffected < 1 {
		return fmt.Errorf("%s: %v", helper.ErrorNotFound, id)
	}
	return result.Error
}
