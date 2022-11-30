package postgres

import (
	"fmt"

	"github.com/ibrat-muslim/students-api/storage/repo"
	"github.com/jmoiron/sqlx"
)

type studentRepo struct {
	db *sqlx.DB
}

func NewStudent(db *sqlx.DB) repo.StudentStorageI {
	return &studentRepo{
		db: db,
	}
}

func (sr *studentRepo) Create(s []*repo.Student) error {
	query := `
		INSERT INTO students (
			first_name,
			last_name,
			username,
			email,
			phone_number
		) VALUES($1, $2, $3, $4, $5)
	`

	for _, student := range s {
		_, err := sr.db.Exec(
			query,
			student.FirstName,
			student.LastName,
			student.UserName,
			student.Email,
			student.PhoneNumber,
		)

		if err != nil {
			return err
		}
	}

	return nil
}

func (sr *studentRepo) GetAll(params *repo.GetAllStudentsParams) (*repo.GetAllStudentsResult, error) {
	result := repo.GetAllStudentsResult{
		Students: make([]*repo.Student, 0),
		Count:    0,
	}

	offset := (params.Page - 1) * params.Limit

	limit := fmt.Sprintf(" LIMIT %d OFFSET %d ", params.Limit, offset)

	filter := ""
	if params.Search != "" {
		str := "%" + params.Search + "%"
		filter += fmt.Sprintf(`
				WHERE first_name ILIKE '%s' OR last_name ILIKE '%s' OR 
				username ILIKE '%s' OR email ILIKE '%s' OR phone_number ILIKE '%s'`,
			str, str, str, str, str,
		)
	}

	query := `
		SELECT 
			id,
			first_name,
			last_name,
			username,
			email,
			phone_number,
			created_at
		FROM students
		` + filter + `
		ORDER BY created_at DESC 
	    ` + limit

	rows, err := sr.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var student repo.Student

		err := rows.Scan(
			&student.ID,
			&student.FirstName,
			&student.LastName,
			&student.UserName,
			&student.Email,
			&student.PhoneNumber,
			&student.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		result.Students = append(result.Students, &student)
	}

	quertCount := `SELECT count(1) FROM students ` + filter

	err = sr.db.QueryRow(quertCount).Scan(&result.Count)

	if err != nil {
		return nil, err
	}

	return &result, nil
}
