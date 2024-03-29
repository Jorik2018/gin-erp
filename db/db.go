package db

import (
	"database/sql"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"encoding/json"
	"fmt"
	"os"
	"log"
	_ "github.com/lib/pq"
	"github.com/Jorik2018/gin-erp/models"
	//_ "github.com/go-sql-driver/mysql" //mysql driver
)

//config for server
type serverConfig struct {
	User         string `json:"user"`
	Password     string `json:"password"`
	DatabaseName string `json:"db-name"`
	HostName     string `json:"host"`
	Port         int    `json:"port"`
}

var ORM *gorm.DB

//GetDatabase - returns a Database object
func GetDatabase() (*sql.DB, error) {
	//setup configuration
	config := loadConfiguration()
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", 
		config.HostName, config.Port, config.User, config.Password, config.DatabaseName)
	//dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", config.User, config.Password, config.HostName, config.Port, config.DatabaseName)
	fmt.Println(dsn)
	db, err := sql.Open("postgres", dsn)
	if err == nil {
		db.SetMaxIdleConns(10)
		db.SetMaxOpenConns(20)
		//defer db.Close()
		if err := db.Ping(); err != nil {
			log.Fatal("Could not connect to the database:", err)
		}
	} else {
		log.Fatal(err)
	}
	//orm, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	orm, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//orm, err := gorm.Open(db, &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	orm.AutoMigrate(&models.Product{})
	orm.AutoMigrate(&models.Book{})
	
	ORM = orm
	return db, err
}

func loadConfiguration() serverConfig {
	file := "./config.json"
	var config serverConfig
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}
