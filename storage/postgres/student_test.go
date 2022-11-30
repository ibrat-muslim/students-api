package postgres_test

import (
	"testing"

	"github.com/bxcodec/faker/v4"
	"github.com/ibrat-muslim/students-api/storage/repo"
	"github.com/stretchr/testify/require"
)

func TestCreateStudents(t *testing.T) {
	var students []*repo.Student

	for i := 0; i < 10; i++ {
		students = append(students, &repo.Student{
			FirstName: faker.FirstName(),
			LastName: faker.LastName(),
			UserName: faker.Username(),
			Email: faker.Email(),
			PhoneNumber: faker.Phonenumber(),
		})
	}

	err := strg.Student().Create(students)

	require.NoError(t, err)
}

func TestGetStudents(t *testing.T) {
	students, err := strg.Student().GetAll(&repo.GetAllStudentsParams{
		Limit: 10,
		Page: 1,
	})

	require.NoError(t, err)
	require.GreaterOrEqual(t, len(students.Students), 1)
	require.GreaterOrEqual(t, int(students.Count), 1)
}