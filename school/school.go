package school

import (
	"context"
	"errors"
	"my-app/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type School struct {
	RDB *redis.Client
	PDB *pgxpool.Pool
}

// FUNCTION
func NewSchool(rdb *redis.Client, pdb *pgxpool.Pool) *School {
	return &School{
		RDB: rdb,
		PDB: pdb,
	}
}

// METHOD
func (s *School) GetStudent() (stu model.Student, err error) {
	ctx := context.Background()
	key := "student:1001"
	s.RDB.HGetAll(ctx, key).Scan(&stu)
	return
}

// METHOD
func (s *School) GetStudentP() (student model.Student, err error) {
	ctx := context.Background()
	id := 1
	query := "SELECT * FROM students WHERE ID = $1"
	row := s.PDB.QueryRow(ctx, query, id)
	err = row.Scan(&student.ID, &student.Name, &student.Age, &student.Job, &student.Score)
	if err == pgx.ErrNoRows {
		err = errors.New("not found")
	}
	return
}
