package models

type LibrarianLoginPayload struct {
	EmpID    string `json:"emp_id" bson:"emp_id"`
	Password string `json:"password" bson:"password"`
}

type StudentLoginPayload struct {
	StdID    string `json:"std_id" bson:"std_id"`
	Password string `json:"password" bson:"password"`
}
