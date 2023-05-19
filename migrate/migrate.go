package main

import (
	"fmt"
	"go-cleanArchitecture/db"
	"go-cleanArchitecture/model"
)

func main() {
	doConn := db.OpenDB()
	defer fmt.Println("マイグレーション成功")
	defer db.CloseDB(doConn)
	doConn.AutoMigrate(&model.User{},&model.Task{})
}
