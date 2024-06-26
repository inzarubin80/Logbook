// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: category.sql

package db

import (
	"context"
)

const createCategory = `-- name: CreateCategory :one
INSERT INTO category (name, created_at) VALUES ($1, NOW()) RETURNING id, name, created_at, updated_at
`

func (q *Queries) CreateCategory(ctx context.Context, name string) (Category, error) {
	row := q.db.QueryRow(ctx, createCategory, name)
	var i Category
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteCategoryByIDs = `-- name: DeleteCategoryByIDs :exec
DELETE FROM category WHERE id = $1
`

func (q *Queries) DeleteCategoryByIDs(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteCategoryByIDs, id)
	return err
}

const findCategoryByIDs = `-- name: FindCategoryByIDs :one
SELECT id, name, created_at, updated_at FROM category WHERE  id = $1 LIMIT 1
`

func (q *Queries) FindCategoryByIDs(ctx context.Context, id int64) (Category, error) {
	row := q.db.QueryRow(ctx, findCategoryByIDs, id)
	var i Category
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getCategorys = `-- name: GetCategorys :many
SELECT id, name, created_at, updated_at FROM category Where name ILIKE  $1 OR $1 = '%%*%%' ORDER BY id DESC
`

func (q *Queries) GetCategorys(ctx context.Context, name string) ([]Category, error) {
	rows, err := q.db.Query(ctx, getCategorys, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Category{}
	for rows.Next() {
		var i Category
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateCategory = `-- name: UpdateCategory :one
UPDATE category SET name = $1,  updated_at = NOW() WHERE id = $2  RETURNING id, name, created_at, updated_at
`

type UpdateCategoryParams struct {
	Name string `json:"name"`
	ID   int64  `json:"id"`
}

func (q *Queries) UpdateCategory(ctx context.Context, arg UpdateCategoryParams) (Category, error) {
	row := q.db.QueryRow(ctx, updateCategory, arg.Name, arg.ID)
	var i Category
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
