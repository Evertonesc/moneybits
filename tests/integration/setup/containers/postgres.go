//go:build integration

package containers

import (
	"context"
	"encoding/json"
	"fmt"
	"moneybits/tests/integration/setup"
	"os"

	"github.com/jackc/pgx/v5"
)

const (
	postgresContainer = "postgres"
)

type TestDatabase struct {
	DB *pgx.Conn
}

func PostgresDB(ctx context.Context) (*TestDatabase, error) {
	container, err := setup.ComposeTestStack.ServiceContainer(ctx, postgresContainer)
	if err != nil {
		return nil, err
	}
	containerState, err := container.State(ctx)
	if err != nil {
		return nil, err
	}
	if !containerState.Running {
		return nil, fmt.Errorf("%s container in invalid state", postgresContainer)
	}

	host, err := container.Host(ctx)
	if err != nil {
		return nil, err
	}

	connStr := fmt.Sprintf(
		"host=%s port=5432 user=postgres password=postgres dbname=moneybits sslmode=disable",
		host,
	)

	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	err = conn.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	return &TestDatabase{
		DB: conn,
	}, nil
}

func (td *TestDatabase) Load(ctx context.Context, fixturePath string, dest any) error {
	data, err := os.ReadFile(fixturePath)
	if err != nil {
		return fmt.Errorf("failed to read fixture file: %v", err)
	}

	if err := json.Unmarshal(data, dest); err != nil {
		return fmt.Errorf("failed to unmarshal fixture: %v", err)
	}

	return nil
}

func (td *TestDatabase) Execute(ctx context.Context, query string, args ...interface{}) error {
	_, err := td.DB.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to execute query: %v", err)
	}

	return nil
}

func (td *TestDatabase) Cleanup() error {
	// Add your cleanup logic here
	// For example, truncate all tables
	return nil
}
