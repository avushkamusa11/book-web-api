package main

import (
	"BooksWebApi/db"
	"BooksWebApi/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Response map[string]any

func main() {

	db.Init()

	app := gin.Default()

	app.GET("/books", func(context *gin.Context) {

		result, err := models.GetAllBooks()
		if err != nil {
			context.JSON(400, Response{
				"message": "Cannot serve your request",
			})
			return
		}

		context.JSON(200, Response{
			"message": "All books in the database",
			"books":   result,
		})
	})

	app.GET("/books/:bookId", func(context *gin.Context) {
		bookIdStr := context.Param("bookId")

		bookId, err := strconv.ParseInt(bookIdStr, 10, 64)
		if err != nil {
			context.JSON(400, Response{
				"message": "Invalid book ID",
			})
			return
		}
		result, err := models.GetBookById(bookId)
		if err != nil {
			context.JSON(400, Response{
				"message": "Cannot serve your request",
			})
			return
		}
		context.JSON(200, Response{
			"message": "Book found",
			"book":    result,
		})
	})

	app.POST("/books", func(context *gin.Context) {

		var bookObject models.Book
		err := context.ShouldBindJSON(&bookObject)
		if err != nil {
			context.JSON(400, Response{
				"message": "Invalid object",
			})
			return
		}
		err = bookObject.Save()

		if err != nil {
			context.JSON(400, Response{
				"message": "Cannot insert book object",
			})
			return
		}

		context.JSON(200, Response{
			"message": "Book created successfully",
			"object":  bookObject,
		})
	})
	app.PUT("/books/:bookId", func(context *gin.Context) {
		bookIdStr := context.Param("bookId")
		bookId, err := strconv.ParseInt(bookIdStr, 10, 64)
		if err != nil {
			context.JSON(400, Response{
				"message": "Invalid book ID",
			})
			return
		}

		var bookObject models.Book
		bookObject.Id = bookId
		err = context.ShouldBindJSON(&bookObject)
		if err != nil {
			context.JSON(400, Response{
				"message": "Invalid object",
			})
			return
		}
		err = bookObject.UpdateBook()

		if err != nil {
			context.JSON(400, Response{
				"message": "Cannot update book object",
			})
			return
		}

		context.JSON(200, Response{
			"message": "Book updated successfully",
			"object":  bookObject,
		})
	})

	app.DELETE("/books/:bookId", func(context *gin.Context) {
		bookIdStr := context.Param("bookId")
		bookId, err := strconv.ParseInt(bookIdStr, 10, 64)
		if err != nil {
			context.JSON(400, Response{
				"message": "Invalid book ID",
			})
			return
		}

		err = models.DeleteBook(bookId)
		if err != nil {
			context.JSON(400, Response{
				"message": err.Error(),
			})
			return
		}

		context.JSON(200, Response{
			"message": "Book deleted successfully",
		})
	})

	err := app.Run(":8087")
	if err != nil {
		fmt.Println("SERVER exception")
		fmt.Println(err)
	}
}
