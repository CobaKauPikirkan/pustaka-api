package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/CobaKauPikirkan/pustaka-api/book"

	"github.com/CobaKauPikirkan/pustaka-api/helper"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type BookHandler struct{
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *BookHandler {
	return &BookHandler{bookService}
}

func (h *BookHandler)GetBookList(c *gin.Context)  {
	books, err := h.bookService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
	})
	return
	}

	// var booksResponse []book.BookResponse
	
	// for _, b := range books{
	// 	bookResponse := ConvertToBookResponse(b)

	// 	bookResponse = append(booksResponse, bookResponse)
	// }

	c.JSON(http.StatusOK, gin.H{
		"data": books,
	})
}

func (h *BookHandler)GetBookById(c *gin.Context)  {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)//convert int to string
	helper.PanicIfError(err)

	b, err := h.bookService.FindById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return 
	}
	
	bookResponse := ConvertToBookResponse(b)

	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}

func (h *BookHandler)CreateBook(c *gin.Context) {
	var bookRequest book.BookRequest

	err := c.ShouldBindJSON(&bookRequest)
	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("error on filed %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return
	}

	book, err := h.bookService.Create(bookRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": ConvertToBookResponse(book),
		
	})
}
func (h *BookHandler)UpdateBook(c *gin.Context) {
	var bookRequest book.BookRequest

	err := c.ShouldBindJSON(&bookRequest)
	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("error on filed %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return
	}

	idString := c.Param("id")
	id, err := strconv.Atoi(idString)//convert int to string
	helper.PanicIfError(err)

	book, err := h.bookService.Update(id, bookRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": ConvertToBookResponse(book),
		
	})
}

func (h *BookHandler)DeleteBook(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)//convert int to string
	helper.PanicIfError(err)

	book, err := h.bookService.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": ConvertToBookResponse(book),
		
	})
}

//refactor function
func ConvertToBookResponse(b book.Book) book.BookResponse {
	return book.BookResponse{
		ID: b.ID,
		Title: b.Title,
		Price: b.Price,
		Description: b.Description,
		Discount: b.Discount,
		Rating: b.Rating,
	}
}
