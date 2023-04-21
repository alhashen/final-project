package config

import (
	"finalproj/model"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Gorm struct {
	Username string
	Password string
	Port     string
	Address  string
	Database string

	DB *gorm.DB
}

type GormDB struct {
	*Gorm
}

var GORM *GormDB

func InitGorm() error {
	GORM = new(GormDB)

	GORM.Gorm = &Gorm{
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Address:  os.Getenv("POSTGRES_ADDRESS"),
		Database: os.Getenv("POSTGRES_DB"),
	}

	err := GORM.Gorm.OpenConnection()
	if err != nil {
		return err
	}

	return nil
}

func (p *Gorm) OpenConnection() error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable",
		p.Address, p.Port, p.Username, p.Database)

	dbConnection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	p.DB = dbConnection

	err = p.DB.Debug().AutoMigrate(model.User{})
	if err != nil {
		return err
	}

	fmt.Println("connected to database:", p.Database, "using gorm")

	return nil
}
