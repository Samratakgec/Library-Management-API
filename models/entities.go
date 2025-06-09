package models

import (
	"time"
)

// defining book model as entity
type Book struct {
	ISBN   string `json:"isbn" bson:"isbn"`
	Name   string `json:"name" bson:"name"`
	Author string `json:"author" bson:"author"`
}

// defining TransactionLogs model as entity
type TransactionLogs struct {
	TimestampIn    time.Time `json:"timestamp_in" bson:"timestamp_in"`
	TimestampOut   time.Time `json:"timestamp_out" bson:"timestamp_out"`
	ISBN           string    `json:"isbn" bson:"isbn"`
	StudentID      string    `json:"std_id" bson:"std_id"`
	FineCollected  float64   `json:"fine_collected" bson:"fine_collected"`
	IsReturned	   bool		 `json:"is_returned" bson:"is_returned"`
}

// defining User model as entity
type User struct {
	EmpID    string             `bson:"emp_id,omitempty"`  // For librarian
	StdID    string             `bson:"std_id,omitempty"`  // For student
	Password string             `bson:"password"`
	Role     string             `bson:"role"`              // "librarian" or "student"
}

