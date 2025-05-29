package services

import (
	"context"
	"errors"
	"time"

	"github.com/Samratakgec/library-management/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Samratakgec/library-management/config"
	//"go.mongodb.org/mongo-driver/bson"
)

func AddBook(book models.Book) error {
	collection := config.GetCollection("book")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := collection.InsertOne(ctx, book)
	return err
}

func GetBookByIsbn(isbn string) (*models.Book, error) {
	collection := config.GetCollection("book")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var book models.Book
	err := collection.FindOne(ctx, bson.M{"isbn": isbn}).Decode(&book)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("book not found")
	}

	return &book, err
}

func DeleteBookByIsbn(isbn string) error {

	book, err := GetBookByIsbn(isbn)
	if book != nil {
		collection := config.GetCollection("book")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		collection.DeleteOne(ctx, bson.M{"isbn": isbn})
	}
	if err != nil && err.Error() != "book not found" {
		return errors.New("internal server error")
	} else if err != nil && err.Error() == "book not found" {
		return errors.New("book not found")
	}
	return nil
}
