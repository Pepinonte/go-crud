// models/setup.go
package models

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
    dsn := "host=localhost user=postgres dbname=go_blog port=5432 sslmode=disable timezone=Europe/Paris"
    database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})  // voir new ids avec culegs

    if err != nil {
        panic("Failed to connect to database!")
    }

    database.AutoMigrate(&Post{})  // register Post model

    DB = database
}