package main

import (
	"fmt"
	"log"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/Jorik2018/gin-erp/handler"
	"github.com/Jorik2018/gin-erp/repository"
)

func main() {
	fmt.Println("setup database connections")
	if err := repository.SetupRepo(); err != nil {
		log.Fatal(err)
	}
	defer repository.CloseRepo()

	fmt.Println("handling roures")
	router := gin.Default()
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	router.GET("/book", handler.BookGetHandler)
	router.POST("/book", handler.BookPostHandler)
	router.PUT("/book", handler.BookPutHandler)
	router.DELETE("/book/:id", handler.BookDeleteHandler)

	router.GET("/student", handler.StudentGetHandler)
	router.POST("/student", handler.StudentPostHandler)
	router.PUT("/student", handler.StudentPutHandler)
	router.DELETE("/student/:id", handler.StudentDeleteHandler)



	router.GET("/ping", handler.ping)
	router.GET("/someJSON", handler.someJSON)
    router.GET("/getb", handler.GetDataB)
    router.GET("/getc", handler.GetDataC)
    router.GET("/getd", handler.GetDataD)
	router.GET("/hello", handler.hello)
	router.MaxMultipartMemory = 8 << 20  // 8 MiB
	router.POST("/upload", handler.upload)

	router.Run(":8080")
}
