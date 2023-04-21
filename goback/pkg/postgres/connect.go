package postgres

import (
	"database/sql"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"context"

	"github.com/pkg/errors"
)

type Tranzactor interface {
	WithTx(ctx context.Context, fn func(ctx context.Context) error) (err error)
}

type Connector interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
	PrepareNamedContext(ctx context.Context, query string) (*sqlx.NamedStmt, error)
	PreparexContext(ctx context.Context, query string) (*sqlx.Stmt, error)
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error

	//problem = not implement to tx
	//NamedQueryContext(ctx context.Context, query string, arg interface{}) (*sqlx.Rows, error)

}

type PG struct {
	Db *sqlx.DB
}

func NewPG(dsn string) (*PG, error) {
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &PG{db}, nil
}

type ctxTxKey struct{}

func (pg *PG) WithTx(ctx context.Context, fn func(ctx context.Context) error) (err error) {
	tx, alreadyHasTx := ctx.Value(ctxTxKey{}).(*sqlx.Tx)
	if !alreadyHasTx {
		tx, err = pg.Db.BeginTxx(ctx, nil)
		if err != nil {

			return errors.WithStack(err)
		}
		ctx = context.WithValue(ctx, ctxTxKey{}, tx)
	}
	err = errors.WithStack(fn(ctx))

	if alreadyHasTx {

		return err
	}

	if err == nil {

		return errors.WithStack(tx.Commit())
	}

	tx.Rollback()

	return err
}

func (pg *PG) ExtractTx(ctx context.Context) Connector {
	tx, ok := ctx.Value(ctxTxKey{}).(*sqlx.Tx)
	if ok {
		return tx
	}
	return pg.Db
}
