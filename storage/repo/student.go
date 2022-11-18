package repo

import "time"

type Student struct {
	ID int64
	FirstName string
	LastName string
	UserName string
	Email string
	PhoneNumber string
	CreatedAt time.Time
}

type GetAllStudentsParams struct {
	Limit int32
	Page int32
	Search string
}

type GetAllStudentsResult struct {
	Students []*Student
	Count int32
}

type StudentStorageI interface {
	Create(s []*Student) (error)
	GetAll(params *GetAllStudentsParams) (*GetAllStudentsResult, error)
}