package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/Jorik2018/gin-erp/models"
	"github.com/Jorik2018/gin-erp/repository"
)

func StudentPost(c *gin.Context) {
	studentRepo := repository.GetStudentRepository()
	var student models.Student
	if err := c.ShouldBindJSON(&student); err == nil {
		_, err = studentRepo.Insert(student)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, true)
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

func StudentGet(c *gin.Context) {
	studentRepo := repository.GetStudentRepository()
	students, err := studentRepo.Select()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, convertListToArray(students))
}

func StudentPut(c *gin.Context) {
	studentRepo := repository.GetStudentRepository()
	var student models.Student
	if err := c.ShouldBindJSON(&student); err == nil {
		_, err = studentRepo.Update(student)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, true)
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

func StudentDelete(c *gin.Context) {
	studentRepo := repository.GetStudentRepository()
	var student models.Student
	idInt, _ := strconv.Atoi(c.Param("id"))
    student.ID = uint(idInt)
	_, err := studentRepo.Remove(student)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, true)
}
