package database

import (
    "fmt"
    "log"
    "os"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "github.com/joho/godotenv"
)

var Db *gorm.DB

func InitDb() *gorm.DB {
    Db = connectDB()
    return Db
}

func connectDB() *gorm.DB {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
    fmt.Println("dsn: ", dsn)
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

    if err != nil {
        fmt.Printf("Error connecting to database: error=%v", err)
        return nil
    }

    return db
}
