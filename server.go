package main

import (
	"fmt"
	"log"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/Jorik2018/gin-erp/handlers"
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

	router.GET("/book", handlers.BookGet)
	router.POST("/book", handlers.BookPost)
	router.PUT("/book", handlers.BookPut)
	router.DELETE("/book/:id", handlers.BookDelete)

	router.GET("/student", handlers.StudentGet)
	router.POST("/student", handlers.StudentPost)
	router.PUT("/student", handlers.StudentPut)
	router.DELETE("/student/:id", handlers.StudentDelete)

	router.GET("/ping", handlers.Ping)
	router.GET("/someJSON", handlers.JSON)
    router.GET("/getb", handlers.GetDataB)
    router.GET("/getc", handlers.GetDataC)
    router.GET("/getd", handlers.GetDataD)
	router.GET("/hello", handlers.Hello)
	router.MaxMultipartMemory = 8 << 20  // 8 MiB
	router.POST("/upload", handlers.Upload)

	router.Run(":7000")
}
