package db

import (
	"errors"
	"log"
	"net"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/leegeobuk/financial-ledger/testutil"
)

var (
	mockContainer *testContainer
	mockDB        DB
)

func TestMain(m *testing.M) {
	// setup
	if err := testutil.SetupConfig("dev"); err != nil {
		log.Fatalf("Error setting up config: %v", err)
	}

	container, err := setupTestContainer(testCtx)
	if err != nil {
		log.Fatalf("Error setting up mock container: %v", err)
	}
	mockContainer = container

	testDB, err := NewMySQL(container.dsn())
	if err != nil {
		log.Fatalf("Error setting up mock db: %v", err)
	}
	mockDB = testDB

	// run tests
	code := m.Run()

	// tear down
	mockDB.Close()
	mockContainer.Terminate(testCtx)
	os.Exit(code)
}

func TestNewMySQL(t *testing.T) {
	tests := []struct {
		name    string
		replace string
		wantErr error
	}{
		{
			name:    "fail case: no /",
			replace: "/",
			wantErr: errors.New(""),
		},
		{
			name:    "fail case: no (",
			replace: "(",
			wantErr: errors.New(""),
		},
		{
			name:    "fail case: no )",
			replace: ")",
			wantErr: errors.New(""),
		},
		{
			name:    "success case: valid dsn",
			replace: "",
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// when
			dsn := strings.Replace(mockContainer.dsn(), tt.replace, "", -1)

			// then
			_, err := NewMySQL(dsn)
			if !reflect.DeepEqual(reflect.TypeOf(err), reflect.TypeOf(tt.wantErr)) {
				t.Errorf("NewMySQL() error = %T, wantErr %T", err, tt.wantErr)
			}
		})
	}
}

func TestMySQL_Ping(t *testing.T) {
	tests := []struct {
		name    string
		config  testContainerConfig
		wantErr error
	}{
		{
			name: "fail case: incorrect user",
			config: testContainerConfig{
				user:     "use",
				password: "1111",
				proto:    "tcp",
				addr:     mockContainer.config.addr,
				schema:   "ledger",
			},
			wantErr: &mysql.MySQLError{},
		},
		{
			name: "fail case: incorrect password",
			config: testContainerConfig{
				user:     "user",
				password: "0000",
				proto:    "tcp",
				addr:     mockContainer.config.addr,
				schema:   "ledger",
			},
			wantErr: &mysql.MySQLError{},
		},
		{
			name: "fail case: incorrect protocol",
			config: testContainerConfig{
				user:     "user",
				password: "1111",
				proto:    "ip",
				addr:     mockContainer.config.addr,
				schema:   "ledger",
			},
			wantErr: &net.OpError{},
		},
		{
			name: "fail case: incorrect address",
			config: testContainerConfig{
				user:     "user",
				password: "1111",
				proto:    "tcp",
				addr:     "localhost:0",
				schema:   "ledger",
			},
			wantErr: &net.OpError{},
		},
		{
			name: "fail case: incorrect schema",
			config: testContainerConfig{
				user:     "user",
				password: "1111",
				proto:    "tcp",
				addr:     mockContainer.config.addr,
				schema:   "unknown",
			},
			wantErr: &mysql.MySQLError{},
		},
		{
			name: "success case: correct user, password, proto, addr, schema",
			config: testContainerConfig{
				user:     "user",
				password: "1111",
				proto:    "tcp",
				addr:     mockContainer.config.addr,
				schema:   "ledger",
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		// when
		testDB, err := NewMySQL(tt.config.dsn())
		if err != nil {
			t.Fatalf("Error setting up test db: %v", err)
		}

		// then
		t.Run(tt.name, func(t *testing.T) {
			err = testDB.Ping()
			if !reflect.DeepEqual(reflect.TypeOf(err), reflect.TypeOf(tt.wantErr)) {
				t.Errorf("Ping() error = %T, wantErr %T", err, tt.wantErr)
			}
		})

		// after each
		t.Cleanup(func() {
			testDB.Close()
		})
	}
}

func TestMySQL_RetryPing(t *testing.T) {
	tests := []struct {
		name     string
		failed   bool
		config   testContainerConfig
		interval time.Duration
		reps     int
		wantErr  error
	}{
		{
			name:   "fail case: incorrect user",
			failed: true,
			config: testContainerConfig{
				user:     "use",
				password: "1111",
				proto:    "tcp",
				addr:     mockContainer.config.addr,
				schema:   "ledger",
			},
			interval: time.Second,
			reps:     3,
			wantErr:  &mysql.MySQLError{},
		},
		{
			name:     "success case: correct user, password, proto, addr, schema",
			failed:   false,
			config:   mockContainer.config,
			interval: time.Second,
			reps:     3,
			wantErr:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// given
			testDB, err := NewMySQL(tt.config.dsn())
			if err != nil {
				t.Fatalf("Error creating test db: %v", err)
			}

			// when
			start := time.Now()
			err = testDB.RetryPing(tt.interval, tt.reps)
			elapsed := time.Now().Sub(start).Round(time.Second)

			// then
			if tt.failed {
				wantDur := tt.interval * time.Duration(tt.reps)
				if elapsed != wantDur {
					t.Errorf("RetryPing() elapsed = %v, wantDur %v", elapsed, wantDur)
				}
			}

			if !reflect.DeepEqual(reflect.TypeOf(err), reflect.TypeOf(tt.wantErr)) {
				t.Errorf("RetryPing() error = %T, wantErr %T", err, tt.wantErr)
			}

			// after each
			t.Cleanup(func() {
				testDB.Close()
			})
		})
	}
}

func TestMySQL_Close(t *testing.T) {
	tests := []struct {
		name    string
		config  testContainerConfig
		wantErr error
	}{
		{
			name:    "success case: db closes without error",
			config:  mockContainer.config,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// given
			testDB, err := NewMySQL(tt.config.dsn())
			if err != nil {
				t.Fatalf("Error creating test db: %v", err)
			}

			// when
			err = testDB.Close()

			// then
			if !reflect.DeepEqual(reflect.TypeOf(err), reflect.TypeOf(tt.wantErr)) {
				t.Errorf("Close() error = %T, wantErr %T", err, tt.wantErr)
			}
		})
	}
}
