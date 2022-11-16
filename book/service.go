package book

type Service interface{
	FindAll()([]Book, error)
	FindById(id int)(Book, error)
	Create(book BookRequest)(Book, error)
	Update(Id int,book BookRequest)(Book, error)
	Delete(Id int)(Book, error)
}
