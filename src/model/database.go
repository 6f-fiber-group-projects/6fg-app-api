package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

func gormConnect() *gorm.DB {
	HOST := os.Getenv("DB_HOST")
	PORT := os.Getenv("DB_PORT")
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASSWORD")
	DBNAME := os.Getenv("DB_NAME")

	DBMS := "postgres"
	CONNECT := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", HOST, PORT, USER, DBNAME, PASS)

	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}

	return db
}

type authority struct {
	Id   int
	Name string
}

func main() {
	db := gormConnect()
	defer db.Close()

	authority := authority{}
	authority.Id = 1

	result := db.First(&authority)

	fmt.Println(result)
}
