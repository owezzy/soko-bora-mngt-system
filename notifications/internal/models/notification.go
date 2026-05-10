package models

type Notification struct {
	ID           string
	OrderID      string
	CustomerID   string
	Type         string
	SMSNumber    string
	Message      string
}
