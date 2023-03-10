package mysql

import (
	"art/domain"
	"context"
	"database/sql"
	"fmt"
)

type mysqlAuthorRepo struct {
	DB *sql.DB
}

func NewMysqlAuthorRepository(db *sql.DB) domain.AuthorRepository {
	fmt.Println(".................")
	return &mysqlAuthorRepo{
		DB: db,
	}
}

func (m *mysqlAuthorRepo) getOne(ctx context.Context, query string, args ...interface{}) (res domain.Author, err error) {
	stmt, err := m.DB.PrepareContext(ctx, query)
	if err != nil {
		return domain.Author{}, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	res = domain.Author{}

	err = row.Scan(
		&res.ID,
		&res.Name,
	)
	return
}

func (m *mysqlAuthorRepo) GetByID(ctx context.Context, id int64) (domain.Author, error) {
	query := `SELECT id, name, created_at, updated_at FROM author WHERE id=?`
	return m.getOne(ctx, query, id)
}
