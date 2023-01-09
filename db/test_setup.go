package db

import (
	"context"
	"fmt"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

var testCtx = context.TODO()

// testContainer represents the test container.
type testContainer struct {
	testcontainers.Container
	config testContainerConfig
}

func (tc *testContainer) dsn() string {
	return tc.config.dsn()
}

type testContainerConfig struct {
	user, password, proto, addr, schema string
}

func (cfg testContainerConfig) dsn() string {
	return fmt.Sprintf("%s:%s@%s(%s)/%s", cfg.user, cfg.password, cfg.proto, cfg.addr, cfg.schema)
}

// setupTestContainer creates an instance of the testContainer type.
func setupTestContainer(ctx context.Context) (*testContainer, error) {
	const (
		user   = "user"
		pw     = "1111"
		schema = "ledger"
	)

	req := testcontainers.ContainerRequest{
		Image:        "mysql:5.7",
		Name:         "ledger-testdb",
		ExposedPorts: []string{"3306/tcp"},
		Env: map[string]string{
			"MYSQL_ROOT_PASSWORD": pw,
			"MYSQL_DATABASE":      schema,
			"MYSQL_USER":          user,
			"MYSQL_PASSWORD":      pw,
		},
		WaitingFor: wait.ForAll(
			wait.ForLog("port: 3306  MySQL Community Server (GPL)"),
			wait.ForListeningPort("3306/tcp"),
		),
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, fmt.Errorf("setup test container: %w", err)
	}

	host, err := container.Host(ctx)
	if err != nil {
		return nil, fmt.Errorf("get test container host: %w", err)
	}

	port, err := container.MappedPort(ctx, "3306/tcp")
	if err != nil {
		return nil, fmt.Errorf("get test container port: %w", err)
	}

	return &testContainer{
		Container: container,
		config: testContainerConfig{
			user:     user,
			password: pw,
			proto:    port.Proto(),
			addr:     fmt.Sprintf("%s:%s", host, port.Port()),
			schema:   schema,
		},
	}, nil
}
