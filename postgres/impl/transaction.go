package appy_driver

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type TransactionFn func(tx Tx) error

type Tx interface {
	Exec(query string, args ...any) (ExecResult, error)
	QueryRow(query string, args ...any) RowResult
	Query(query string, args ...any) (RowsResult, error)
	Commit() error
	Rollback() error
	CommitOrRollback()
}

type txScoped struct {
	internal pgx.Tx
}

type txDirect struct {
}

type ExecResult struct {
	res pgconn.CommandTag
}

type RowResult struct {
	isRead bool
	row    pgx.Row
}

type RowsResult struct {
	rows pgx.Rows
}

type Scannable interface {
	Scan(dest ...any) error
	Close()
	Next() bool
	Err() error
}

func StartTransaction() (Tx, error) {
	tx, err := gDatabaseConnection.Begin(context.TODO())
	if err != nil {
		return nil, err
	}

	return &txScoped{
		internal: tx,
	}, nil
}

func RunTransaction(fn TransactionFn) error {
	tx, err := StartTransaction()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = fn(tx)

	if err == nil {
		return tx.Commit()
	}

	return err
}

func DirectConnection() Tx {
	return &txDirect{}
}

func (tx *txScoped) Exec(query string, args ...any) (ExecResult, error) {
	res, err := tx.internal.Exec(context.TODO(), query, args...)
	if err != nil {
		return ExecResult{}, err
	}

	return ExecResult{
		res: res,
	}, nil
}

func (tx *txScoped) QueryRow(query string, args ...any) RowResult {
	res := tx.internal.QueryRow(context.TODO(), query, args...)
	return RowResult{
		isRead: false,
		row:    res,
	}
}

func (tx *txScoped) Query(query string, args ...any) (RowsResult, error) {
	res, err := tx.internal.Query(context.TODO(), query, args...)
	return RowsResult{
		rows: res,
	}, err
}

func (tx *txScoped) Commit() error {
	return tx.internal.Commit(context.TODO())
}

func (tx *txScoped) Rollback() error {
	return tx.internal.Rollback(context.TODO())
}

// Commit the transaction or rollback on failure
func (tx *txScoped) CommitOrRollback() {
	if err := tx.Commit(); err != nil {
		tx.Rollback()
	}
}

func (tx *txDirect) Exec(query string, args ...any) (ExecResult, error) {
	res, err := gDatabaseConnection.Exec(context.TODO(), query, args...)
	if err != nil {
		return ExecResult{}, err
	}

	return ExecResult{
		res: res,
	}, nil
}

func (tx *txDirect) QueryRow(query string, args ...any) RowResult {
	res := gDatabaseConnection.QueryRow(context.TODO(), query, args...)
	return RowResult{
		isRead: false,
		row:    res,
	}
}

func (tx *txDirect) Query(query string, args ...any) (RowsResult, error) {
	res, err := gDatabaseConnection.Query(context.TODO(), query, args...)
	return RowsResult{
		rows: res,
	}, err
}

func (tx *txDirect) Commit() error {
	return nil
}

func (tx *txDirect) Rollback() error {
	return nil
}

// Commit the transaction or rollback on failure
func (tx *txDirect) CommitOrRollback() {
	if err := tx.Commit(); err != nil {
		tx.Rollback()
	}
}

// Compatibility only
func (rr RowResult) Err() error {
	return nil
}

func (rr RowResult) Scan(dest ...any) error {
	rr.isRead = true
	return rr.row.Scan(dest...)
}

func (rr RowResult) Close() {
	// Do nothing
}

func (rr RowResult) Next() bool {
	return rr.isRead
}

func (rr RowsResult) Scan(dest ...any) error {
	return rr.rows.Scan(dest...)
}

func (rr RowsResult) Err() error {
	return rr.rows.Err()
}

func (rr RowsResult) Close() {
	rr.rows.Close()
}

func (rr RowsResult) Next() bool {
	return rr.rows.Next()
}

func (er ExecResult) RowsAffected() int64 {
	return er.res.RowsAffected()
}
