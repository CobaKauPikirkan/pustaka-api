package book

type Repository interface{
	FindAll()([]Book, error)
	FindById(id int)(Book, error)
	Create(book Book)(Book, error)
	Update(bokk Book)(Book, error)
	Delete(book Book)(Book, error)
}