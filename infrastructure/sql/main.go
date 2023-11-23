package sql

import (
	"base-plate/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Client struct {
	*gorm.DB
}

func NewClient() *Client {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Cfg.GetString("MYSQL_USER"), config.Cfg.GetString("MYSQL_PASS"), config.Cfg.GetString("MYSQL_HOST"), config.Cfg.GetString("MYSQL_PORT"), config.Cfg.GetString("MYSQL_DBNAME"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Cannot Connect to Database")
	}

	return &Client{db}
}

func NewPostgreSQLClient() *Client {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", config.Cfg.GetString("POSTGRES_HOST"), config.Cfg.GetString("POSTGRES_USER"), config.Cfg.GetString("POSTGRES_PASS"), config.Cfg.GetString("POSTGRES_DBNAME"), config.Cfg.GetInt("POSTGRES_PORT"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Cannot Connect to PostgreSQL")
	}

	return &Client{db}
}
