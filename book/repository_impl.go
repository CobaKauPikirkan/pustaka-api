package book

import (
	"gorm.io/gorm"
	"github.com/CobaKauPikirkan/pustaka-api/helper"
)

type repositoryImpl struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *repositoryImpl {
	return &repositoryImpl{db}
}

func (r *repositoryImpl)FindAll() ([]Book, error) {
	var books []Book
	
	err:= r.db.Find(&books).Error
	helper.PanicIfError(err)
	
	return books, err 
}
func (r *repositoryImpl)FindById(id int) (Book, error) {
	var book Book

	err:= r.db.Find(&book, id).Error
	helper.PanicIfError(err)
	
	return book, err 
}

func (r *repositoryImpl)Create(book Book) (Book, error) {
	err:= r.db.Create(&book).Error
	helper.PanicIfError(err)
	return book, err
}

func (r *repositoryImpl)Update(book Book) (Book, error) {
	err:= r.db.Save(&book).Error
	helper.PanicIfError(err)

	return book, err
}

func (r *repositoryImpl)Delete(book Book) (Book, error) {
	err:= r.db.Delete(&book).Error
	helper.PanicIfError(err)
	
	return book, err
}