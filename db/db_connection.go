package db

import (
	"database/sql"
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

const projectDirName = "PRACTICE"

func LoadEnv() {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))
	err := godotenv.Load(string(rootPath) + `./config/.env`)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error loading .env file")
	}
}

func CreateCon() *sql.DB {
	LoadEnv()
	cfg := mysql.Config{
		User:                 os.Getenv("dbuser"),
		Passwd:               os.Getenv("dbpassword"),
		Net:                  "tcp",
		Addr:                 os.Getenv("dbhost"),
		DBName:               os.Getenv("db"),
		AllowNativePasswords: true,
	}
	fmt.Println("cfg", cfg)
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("db is connected")
	}
	db.SetMaxOpenConns(1000)
	db.SetMaxIdleConns(1000)
	db.SetConnMaxLifetime(5 * time.Minute)

	//defer db.Close()
	// make sure connection is available
	fmt.Println("Configuration:", cfg)
	fmt.Println("DSN:", cfg.FormatDSN())
	err = db.Ping()
	if err != nil {
		fmt.Println("db is not connected")
		fmt.Println(err.Error())
	} else {
		fmt.Println("DB connection is available")
	}

	return db

}
