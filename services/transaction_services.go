package services

import (
	"context"
	"errors"
	"time"

	"github.com/Samratakgec/library-management/config"
	"github.com/Samratakgec/library-management/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetTransactionLogByIsbnAndStdID(isbn string, std_id string) (*models.TransactionLogs, error) {
	collection := config.GetCollection("TransactionLogs")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var txnlog models.TransactionLogs
	err := collection.FindOne(ctx, bson.M{
		"isbn":        isbn,
		"std_id":      std_id,
		"is_returned": false,
	}).Decode(&txnlog)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("record not found")
		}
		return nil, errors.New("internal server error")
	}

	return &txnlog, nil
}
func CanBookBeAlloted(std_id string) (bool,error)  {
	collection := config.GetCollection("TransactionLogs")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var txnlog models.TransactionLogs
	err := collection.FindOne(ctx, bson.M{
		"std_id" : std_id,
		"is_returned" : false,
	}).Decode(&txnlog)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return true,nil // allowed
		} else {
			return false,err // 500
		}
	}else {
		return false,nil // not-allowed
	}

}
func AllocateBook(transactionLogs models.TransactionLogs) error {
	collection := config.GetCollection("TransactionLogs")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()


	transactionLogs.TimestampOut = time.Now()
	transactionLogs.IsReturned = false

	_, err := collection.InsertOne(ctx, transactionLogs)
	return err
}

func DeAllocateBook(transactionLogs models.TransactionLogs) error {
	collection := config.GetCollection("TransactionLogs")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()


	filter := bson.M{
		"isbn":        transactionLogs.ISBN,
		"std_id":      transactionLogs.StudentID,
		"is_returned": false,
	}
	update := bson.M{
		"$set": bson.M{
			"timestamp_in":   time.Now(),
			"fine_collected": transactionLogs.FineCollected,
			"is_returned":    true,
		},
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	return err
}
