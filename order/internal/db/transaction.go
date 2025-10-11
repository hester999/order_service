package db

import (
	"context"
	"database/sql"
	"database/sql/driver"
)

// TransactionManager интерфейс для управления транзакциями
type TransactionManager interface {
	// BeginTx начинает новую транзакцию
	BeginTx(ctx context.Context, opts *sql.TxOptions) (Transaction, error)
}

// Transaction интерфейс транзакции
type Transaction interface {
	driver.Tx
	// ExecContext выполняет запрос без возврата результата
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	// QueryContext выполняет запрос с возвратом строк
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	// QueryRowContext выполняет запрос с возвратом одной строки
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	// PrepareContext подготавливает запрос
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
}

// TxWrapper обертка над sql.Tx для реализации интерфейса Transaction
type TxWrapper struct {
	*sql.Tx
}

// NewTxWrapper создает новую обертку транзакции
func NewTxWrapper(tx *sql.Tx) *TxWrapper {
	return &TxWrapper{Tx: tx}
}

// TransactionManagerImpl реализация TransactionManager для sqlx.DB
type TransactionManagerImpl struct {
	db *sql.DB
}

// NewTransactionManager создает новый менеджер транзакций
func NewTransactionManager(db *sql.DB) TransactionManager {
	return &TransactionManagerImpl{db: db}
}

// BeginTx начинает новую транзакцию
func (tm *TransactionManagerImpl) BeginTx(ctx context.Context, opts *sql.TxOptions) (Transaction, error) {
	tx, err := tm.db.BeginTx(ctx, opts)
	if err != nil {
		return nil, err
	}
	return NewTxWrapper(tx), nil
}
