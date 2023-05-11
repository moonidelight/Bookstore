package repository

import (
	"fmt"
	"github.com/moonidelight/go_course/lab3/internal/domain/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	//"log"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository() *BookRepository {
	dsn := fmt.Sprintf("host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Almaty",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	//dsn := os.Getenv("DATABASE_URL")
	//dsn := "host=localhost user=postgres password=postgres dbname=gorm port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database =(. \n", err)
		os.Exit(2)
		return nil
	}
	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("running migrations")
	db.AutoMigrate(&entity.Book{})
	return &BookRepository{db: db}
}
func (repo *BookRepository) Create(title, desc string, cost float64) {
	repo.db.Create(&entity.Book{Title: title, Description: desc, Cost: cost})

}
func (repo *BookRepository) GetById(id int) (string, string, float64) {
	var book entity.Book
	repo.db.First(&book, uint(id))
	fmt.Println(book.Title)
	return book.Title, book.Description, book.Cost
}
func (repo *BookRepository) GetAll() []entity.Book {
	var books []entity.Book
	repo.db.Find(&books)
	return books
}
func (repo *BookRepository) UpdateTitleAndDesc(id int, title, desc string) (string, string) {
	var book entity.Book
	repo.db.First(&book, uint(id))
	repo.db.Model(&book).Updates(entity.Book{Title: title, Description: desc})
	return book.Title, book.Description
}
func (repo *BookRepository) DeleteById(id int) {
	var book entity.Book
	repo.db.Unscoped().Delete(&book, uint(id))
}
func (repo *BookRepository) SearchByTitle(title string) (string, string, float64) {
	var book entity.Book
	repo.db.Where("title = ?", title).Find(&book)
	return title, book.Description, book.Cost
}
func (repo *BookRepository) GetSortedBooksAsc() []entity.Book {
	var books []entity.Book
	repo.db.Order("cost asc").Find(&books)
	return books
}
func (repo *BookRepository) GetSortedBooksDesc() []entity.Book {
	var books []entity.Book
	repo.db.Order("cost desc").Find(&books)
	return books
}
