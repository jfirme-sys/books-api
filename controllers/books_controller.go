package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jfirme-sys/books-api/database"
	"github.com/jfirme-sys/books-api/models"
)

func ShowBook(context *gin.Context) {
	id := context.Param("id")

	integerID, err := strconv.Atoi(id)
	if err != nil {
		context.JSON(400, gin.H{
			"error": "ID has to be integer",
		})

		return
	}

	db := database.GetDB()

	var book models.Book
	err = db.First(&book, integerID).Error
	if err != nil {
		context.JSON(400, gin.H{
			"error": "cannot find book: " + err.Error(),
		})

		return
	}

	context.JSON(200, book)
}

func CreateBook(context *gin.Context) {
	db := database.GetDB()

	var book models.Book

	err := context.ShouldBindJSON(&book)

	if err != nil {
		context.JSON(400, gin.H{
			"error": "cannot bind JSON book: " + err.Error(),
		})

		return
	}

	err = db.Create(&book).Error

	if err != nil {
		context.JSON(400, gin.H{
			"error": "cannot create book: " + err.Error(),
		})

		return
	}

	context.JSON(200, book)
}

func ShowAllBooks(context *gin.Context) {
	db := database.GetDB()

	var books []models.Book

	err := db.Find(&books).Error
	if err != nil {
		context.JSON(400, gin.H{
			"error": "cannot list book: " + err.Error(),
		})

		return
	}

	context.JSON(200, books)
}

func UpdateBook(context *gin.Context) {
	db := database.GetDB()

	var book models.Book

	err := context.ShouldBindJSON(&book)

	if err != nil {
		context.JSON(400, gin.H{
			"error": "cannot bind JSON book: " + err.Error(),
		})

		return
	}

	err = db.Save(&book).Error

	if err != nil {
		context.JSON(400, gin.H{
			"error": "cannot update book: " + err.Error(),
		})

		return
	}

	context.JSON(200, book)
}

func DeleteBook(context *gin.Context) {
	id := context.Param("id")

	integerID, err := strconv.Atoi(id)
	if err != nil {
		context.JSON(400, gin.H{
			"error": "ID has to be integer",
		})

		return
	}

	db := database.GetDB()

	err = db.Delete(&models.Book{}, integerID).Error

	if err != nil {
		context.JSON(400, gin.H{
			"error": "cannot delete book",
		})

		return
	}

	context.Status(204)
}
