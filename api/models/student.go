package models

import "time"

type Student struct {
	ID          int64     `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	UserName    string    `json:"username"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
}

type CreateStudentRequest struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	UserName    string `json:"username"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

type StudentsArray []*CreateStudentRequest

type GetAllStudentsParams struct {
	Limit  int32  `json:"limit" binding:"required" default:"10"`
	Page   int32  `json:"page" binding:"required" default:"1"`
	Search string `json:"search"`
}

type GetAllStudentsResponse struct {
	Students []*Student `json:"students"`
	Count    int32      `json:"count"`
}
