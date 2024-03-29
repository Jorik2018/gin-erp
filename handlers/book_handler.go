package handlers

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/Jorik2018/gin-erp/models"
	"github.com/Jorik2018/gin-erp/repository"
)

//BookGetHandler - handle book get requests
func BookGet(c *gin.Context) {
	bookRepo := repository.GetBookRepository()
	books, err := bookRepo.Select()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, convertListToArray(books))
}

func BookPost(c *gin.Context) {
	bookRepo := repository.GetBookRepository()

	var book models.Book
	if err := c.ShouldBindJSON(&book); err == nil {
		_, err = bookRepo.Insert(book)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, true)
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

func BookPut(c *gin.Context) {
	bookRepo := repository.GetBookRepository()
	var book models.Book
	if err := c.ShouldBindJSON(&book); err == nil {
		_, err = bookRepo.Update(book)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, true)
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

//BookDeleteHandler - handle book delete requests
func BookDelete(c *gin.Context) {
	bookRepo := repository.GetBookRepository()
	var book models.Book
	book.BookID, _ = strconv.Atoi(c.Param("id"))
	_, err := bookRepo.Remove(book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, true)
}
