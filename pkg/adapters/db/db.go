package db

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"orcinator/pkg/domain/saga"
	"orcinator/pkg/dto"
)

type Adapter struct {
	db *gorm.DB
}

func NewAdapter(dsn string) (*Adapter, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("db connection failure: %v", err)
	}

	db.AutoMigrate(&dto.SagaDto{})

	return &Adapter{db}, nil
}

type User struct {
	gorm.Model
	Name string
}

func (a *Adapter) CreateSaga(saga saga.Saga) {
	dto := dto.NewCreateSagaDto(saga)
	result := a.db.Create(&dto)
	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Printf("Inserted saga with id %s", saga.Id())
}
