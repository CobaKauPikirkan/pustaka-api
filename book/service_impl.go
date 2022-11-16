package book

import "github.com/CobaKauPikirkan/pustaka-api/helper"

type service struct {
	repository Repository
}

func NewBookService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	books, err := s.repository.FindAll()
	helper.PanicIfError(err)
	
	return books, err
}
func (s *service) FindById(id int) (Book, error) {
	book, err := s.repository.FindById(id)
	helper.PanicIfError(err)

	return book, err
}

func (s *service) Create(bookRequest BookRequest) (Book, error) {
	price, _ := bookRequest.Price.Int64()
	rating, _ := bookRequest.Rating.Int64()
	discount,_ := bookRequest.Discount.Int64()
	
	book:= Book{
		Title: bookRequest.Title,
		Price: int(price),
		Description: bookRequest.Description,
		Rating: int(rating),
		Discount: int(discount),
	}
	newBook, err := s.repository.Create(book)
	helper.PanicIfError(err)

	return newBook, err
}
func (s *service) Update(Id int,bookRequest BookRequest) (Book, error) {
	book, err := s.repository.FindById(Id)
	helper.PanicIfError(err)

	price, _ := bookRequest.Price.Int64()
	rating, _ := bookRequest.Rating.Int64()
	discount,_ := bookRequest.Discount.Int64()
	
		book.Title= bookRequest.Title
		book.Price= int(price)
		book.Description= bookRequest.Description
		book.Rating= int(rating)
		book.Discount= int(discount)
	
	newBook, err := s.repository.Update(book)
	helper.PanicIfError(err)

	return newBook, err
}
func (s *service) Delete(Id int) (Book, error) {
	book, err := s.repository.FindById(Id)
	helper.PanicIfError(err)
	
	newBook, err := s.repository.Delete(book)
	helper.PanicIfError(err)

	return newBook, err
}