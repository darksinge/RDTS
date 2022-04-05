package main

import (
	"fmt"
	"orcinator/pkg/adapters/db"
	"orcinator/pkg/domain/saga"
	"os"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	dbAdapter, err := db.NewAdapter(dsn)
	if err != nil {
		panic(err)
	}

	svc := saga.NewService("test svc")
	e := saga.NewEvent("test", svc)
	dbAdapter.CreateSaga(saga.New(uuid.New(), e))
}
