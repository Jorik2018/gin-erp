package main

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/Jorik2018/go_crud/handlers"
	"github.com/Jorik2018/go_crud/commons"
)

const (
	dbHost     = "190.119.114.163"
	dbPort     = 3306 // Replace with the appropriate port number for your remote MySQL server
	dbUser     = "supersgtii"
	dbPassword = "A1_supabase"
	dbName     = "production"
)



func main() {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", dsn)
	
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatal("Could not connect to the database:", err)
	}
	
	orm, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})//gorm.Open(db, &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	orm.AutoMigrate(&commons.Product{})

	/*db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUser, dbPassword, dbHost, dbPort, dbName))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()*/

	// Check database connection


	http.HandleFunc("/users", handlers.GetHandler(db))

	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}