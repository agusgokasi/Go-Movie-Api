package config

import (
	"MovieApi/model"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	// db configuration
	Username string
	Password string
	Port     string
	Address  string
	Database string

	// db connection
	DB *gorm.DB
}

type PsqlDb struct {
	*Postgres
}

var (
	PSQL *PsqlDb
)

func InitPostgres() error {
	PSQL = new(PsqlDb)

	PSQL.Postgres = &Postgres{
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Address:  os.Getenv("POSTGRES_ADDRESS"),
		Database: os.Getenv("POSTGRES_DB"),
	}

	// connect to database
	err := PSQL.Postgres.OpenConnection()
	if err != nil {
		return err
	}

	// Migrate the schema
	err = PSQL.DB.AutoMigrate(&model.Movie{})
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) OpenConnection() error {
	// init dsn
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", p.Address, p.Port, p.Username, p.Database, p.Password)

	dbConnection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	p.DB = dbConnection

	// test connection
	db, err := p.DB.DB()
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	fmt.Println("Successfully connected to database")
	return nil
}

func GetDB() *gorm.DB {
	return PSQL.DB
}
