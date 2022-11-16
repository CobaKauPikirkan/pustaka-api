package main

import (
	//"fmt"
	"log"

	"github.com/CobaKauPikirkan/pustaka-api/book"
	"github.com/CobaKauPikirkan/pustaka-api/handler"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/pustaka_api?charset=utf8mb4&parseTime=True&loc=Local"
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}	
	
	db.AutoMigrate(&book.Book{})
	
	bookRepository:= book.NewBookRepository(db)
	bookService:= book.NewBookService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)
	


	router := gin.Default()

	v1:= router.Group("/v1")

	v1.POST("/books", bookHandler.CreateBook)
	v1.GET("/books", bookHandler.GetBookList)
	v1.GET("/books/:id", bookHandler.GetBookById)
	v1.PUT("/books/:id", bookHandler.UpdateBook)
	v1.DELETE("/books/:id",bookHandler.DeleteBook)
	router.Run(":8080")
}

