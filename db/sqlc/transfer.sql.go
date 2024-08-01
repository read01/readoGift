// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: transfer.sql

package db

import (
	"context"
	"database/sql"
)

const createTransfer = `-- name: CreateTransfer :one
INSERT INTO transfers (
  from_account_id,
  to_account_id,
  amount
) VALUES (
  $1, $2, $3
) RETURNING id, from_account_id, to_account_id, amount, created_at
`

type CreateTransferParams struct {
	FromAccountID sql.NullInt64 `json:"from_account_id"`
	ToAccountID   sql.NullInt64 `json:"to_account_id"`
	Amount        int64         `json:"amount"`
}

func (q *Queries) CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfers, error) {
	row := q.db.QueryRowContext(ctx, createTransfer, arg.FromAccountID, arg.ToAccountID, arg.Amount)
	var i Transfers
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const deleteTransfer = `-- name: DeleteTransfer :one
DELETE FROM transfers
WHERE id = $1
RETURNING id, from_account_id, to_account_id, amount, created_at
`

func (q *Queries) DeleteTransfer(ctx context.Context, id int64) (Transfers, error) {
	row := q.db.QueryRowContext(ctx, deleteTransfer, id)
	var i Transfers
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const getTransfer = `-- name: GetTransfer :one
SELECT id, from_account_id, to_account_id, amount, created_at FROM transfers
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTransfer(ctx context.Context, id int64) (Transfers, error) {
	row := q.db.QueryRowContext(ctx, getTransfer, id)
	var i Transfers
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const listTransfer = `-- name: ListTransfer :many
SELECT id, from_account_id, to_account_id, amount, created_at FROM transfers
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListTransferParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListTransfer(ctx context.Context, arg ListTransferParams) ([]Transfers, error) {
	rows, err := q.db.QueryContext(ctx, listTransfer, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Transfers{}
	for rows.Next() {
		var i Transfers
		if err := rows.Scan(
			&i.ID,
			&i.FromAccountID,
			&i.ToAccountID,
			&i.Amount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTransfer = `-- name: UpdateTransfer :exec
UPDATE transfers
SET amount = $2
WHERE id = $1
`

type UpdateTransferParams struct {
	ID     int64 `json:"id"`
	Amount int64 `json:"amount"`
}

func (q *Queries) UpdateTransfer(ctx context.Context, arg UpdateTransferParams) error {
	_, err := q.db.ExecContext(ctx, updateTransfer, arg.ID, arg.Amount)
	return err
}
