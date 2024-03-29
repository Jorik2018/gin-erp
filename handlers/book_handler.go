package handlers

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/Jorik2018/gin-erp/models"
	"github.com/Jorik2018/gin-erp/repository"
)

type BookHandler struct {
}

func NewBookHandler() *BookHandler {
    return &BookHandler{}
}

func (h *BookHandler) get(c *gin.Context) {
	bookRepo := repository.GetBookRepository()
	book, err := bookRepo.Find(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, book)
}

func (h *BookHandler) list(c *gin.Context) {
	bookRepo := repository.GetBookRepository()
	books, err := bookRepo.Select()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, convertListToArray(books))
}

func (h *BookHandler) Post(c *gin.Context) {
	bookRepo := repository.GetBookRepository()
	var book models.Book
	if err := c.ShouldBindJSON(&book); err == nil {
		_, err = bookRepo.Insert(&book)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, true)
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

func put(c *gin.Context) {
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
func delete(c *gin.Context) {
	bookRepo := repository.GetBookRepository()
	var book models.Book
	idInt, _ := strconv.Atoi(c.Param("id"))
    book.ID = uint(idInt)
	_, err := bookRepo.Remove(book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, true)
}


func (h *BookHandler) SetupRoutes(router *gin.Engine) {
	router.GET("/api/book", h.list)
	router.GET("/api/book/:id", h.get)
	router.POST("/api/book", h.Post)
	router.PUT("/api/book/:id", put)
	router.DELETE("/api/book/:id", delete)
}
