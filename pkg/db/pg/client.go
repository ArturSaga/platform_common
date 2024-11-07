package pg

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"

	"github.com/ArturSaga/chat-server/internal/client/db"
)

type pgClient struct {
	masterDBC db.DB
}

// New - публичный метод, создающий сущность клиента бд
func New(ctx context.Context, dsn string) (db.Client, error) {
	dbc, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		return nil, errors.Errorf("failed to connect to db: %v", err)
	}

	return &pgClient{
		masterDBC: &pg{dbc: dbc},
	}, nil
}

// DB - публичный метод, возвращающий соединение с бд
func (c *pgClient) DB() db.DB {
	return c.masterDBC
}

// Close - публичный метод, закрывающий соединение с бд
func (c *pgClient) Close() error {
	if c.masterDBC != nil {
		c.masterDBC.Close()
	}

	return nil
}
