package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenDB() *gorm.DB {
	// 環境変数取得
	if os.Getenv("GO_ENV") == "dev" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalln(err)
		}
	}
	// DB接続
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",os.Getenv("POSTGRES_USER"),
	os.Getenv("POSTGRES_PW"),os.Getenv("POSTGRES_HOST"),os.Getenv("POSTGRES_PORT"),os.Getenv("POSTGRES_DB"))

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("DB接続成功")
		return db
}

func CloseDB(db *gorm.DB) {
	dbSQL, _ := db.DB()
	if err := dbSQL.Close(); err != nil {
		log.Fatalln(err)
	}
}
