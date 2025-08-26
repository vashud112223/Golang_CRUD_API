package storage

import "github.com/ashutosh/students-api/internal/types"

type Storage interface {
	CreateStudent(name string, email string, age int) (int64, error)
	GetStudentById(id int64) (types.Student, error)
	GetStudents() ([]types.Student, error)
	UpdateStudent(student types.Student) error
	DeleteStudent(id int64) error
}
