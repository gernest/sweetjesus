package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"os"
	"fmt"
)

func ConnectLocalDB() (*gorm.DB, error) {
	dns:=os.Getenv("LOCAL_POSTGRES_URL")
	fmt.Println("connecting to ..",dns)
	db, err := gorm.Open("postgres",dns )
	db.SingularTable(true)
	db.LogMode(true)
	return &db, err
}

func RemoteDB() (*gorm.DB, error) {
	dns:=os.Getenv("REMOTE_POSTGRES_URL")
	fmt.Println("connecting to..", dns)
	db, err := gorm.Open("postgres", dns)
	db.LogMode(true)
	db.SingularTable(true)
	return &db, err
}
