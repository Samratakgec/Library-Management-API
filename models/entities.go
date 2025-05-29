package models

import (
	"time"
)

type Book struct {
	ISBN   string `json:"isbn" bson:"isbn"`
	Name   string `json:"name" bson:"name"`
	Author string `json:"author" bson:"author"`
}

type TransactionLogs struct {
	TimestampIn    time.Time `json:"timestamp_in" bson:"timestamp_in"`
	TimestampOut   time.Time `json:"timestamp_out" bson:"timestamp_out"`
	ISBN           string    `json:"isbn" bson:"isbn"`
	StudentID      string    `json:"std_id" bson:"std_id"`
	FineCollected  float64   `json:"fine_collected" bson:"fine_collected"`
	IsReturned	   bool		 `json:"is_returned" bson:"is_returned"`
}


type User struct {
	EmpID    string             `bson:"emp_id,omitempty"`  // For librarian
	StdID    string             `bson:"std_id,omitempty"`  // For student
	Password string             `bson:"password"`
	Role     string             `bson:"role"`              // "librarian" or "student"
}

