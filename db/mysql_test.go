package db

import (
	"log"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/leegeobuk/financial-ledger/testutil"
)

var (
	mockContainer *testContainer
	mockDB        DB
)

func TestMain(m *testing.M) {
	// setup
	if err := testutil.SetupConfig("local"); err != nil {
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
	_ = mockDB.Close()
	_ = mockContainer.Terminate(testCtx)
	os.Exit(code)
}

func TestNewMySQL(t *testing.T) {
	tests := []struct {
		name       string
		shouldFail bool
		replace    string
		wantErrStr string
	}{
		{
			name:       "fail case: no /",
			shouldFail: true,
			replace:    "/",
			wantErrStr: "new MySQL",
		},
		{
			name:       "fail case: no (",
			shouldFail: true,
			replace:    "(",
			wantErrStr: "new MySQL",
		},
		{
			name:       "fail case: no )",
			shouldFail: true,
			replace:    ")",
			wantErrStr: "new MySQL",
		},
		{
			name:       "success case: valid dsn",
			replace:    "",
			wantErrStr: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// given
			dsn := strings.ReplaceAll(mockContainer.dsn(), tt.replace, "")

			// when
			_, err := NewMySQL(dsn)

			// then
			if tt.shouldFail {
				if got := err.Error(); !strings.HasPrefix(got, tt.wantErrStr) {
					t.Errorf("NewMySQL() error = %v, wantErrStr %s", err, tt.wantErrStr)
				}

				return
			}

			if err != nil {
				t.Errorf("NewMySQL() error = %v, wantErrStr %s", err, tt.wantErrStr)
			}
		})
	}
}

func TestMySQL_Ping(t *testing.T) {
	tests := []struct {
		name       string
		shouldFail bool
		config     testContainerConfig
		wantErrStr string
	}{
		{
			name:       "fail case: incorrect user",
			shouldFail: true,
			config: testContainerConfig{
				user:     "use",
				password: "1111",
				proto:    "tcp",
				addr:     mockContainer.config.addr,
				schema:   "ledger",
			},
			wantErrStr: "ping MySQL",
		},
		{
			name:       "fail case: incorrect password",
			shouldFail: true,
			config: testContainerConfig{
				user:     "user",
				password: "0000",
				proto:    "tcp",
				addr:     mockContainer.config.addr,
				schema:   "ledger",
			},
			wantErrStr: "ping MySQL",
		},
		{
			name:       "fail case: incorrect protocol",
			shouldFail: true,
			config: testContainerConfig{
				user:     "user",
				password: "1111",
				proto:    "ip",
				addr:     mockContainer.config.addr,
				schema:   "ledger",
			},
			wantErrStr: "ping MySQL",
		},
		{
			name:       "fail case: incorrect address",
			shouldFail: true,
			config: testContainerConfig{
				user:     "user",
				password: "1111",
				proto:    "tcp",
				addr:     "localhost:0",
				schema:   "ledger",
			},
			wantErrStr: "ping MySQL",
		},
		{
			name:       "fail case: incorrect schema",
			shouldFail: true,
			config: testContainerConfig{
				user:     "user",
				password: "1111",
				proto:    "tcp",
				addr:     mockContainer.config.addr,
				schema:   "unknown",
			},
			wantErrStr: "ping MySQL",
		},
		{
			name:       "success case: correct user, password, proto, addr, schema",
			shouldFail: false,
			config: testContainerConfig{
				user:     "user",
				password: "1111",
				proto:    "tcp",
				addr:     mockContainer.config.addr,
				schema:   "ledger",
			},
			wantErrStr: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// given
			testDB, err := NewMySQL(tt.config.dsn())
			if err != nil {
				t.Fatalf("Error setting up test db: %v", err)
			}

			// when
			err = testDB.Ping()

			// then
			if tt.shouldFail {
				if got := err.Error(); !strings.HasPrefix(got, tt.wantErrStr) {
					t.Errorf("Ping() error = %v, wantErrStr %s", err, tt.wantErrStr)
				}

				return
			}

			if err != nil {
				t.Errorf("Ping() error = %v, wantErrStr %s", err, tt.wantErrStr)
			}

			// after each
			t.Cleanup(func() {
				testDB.Close()
			})
		})
	}
}

func TestMySQL_RetryPing(t *testing.T) {
	tests := []struct {
		name       string
		shouldFail bool
		config     testContainerConfig
		interval   time.Duration
		reps       int
		wantErrStr string
	}{
		{
			name:       "fail case: incorrect user",
			shouldFail: true,
			config: testContainerConfig{
				user:     "use",
				password: "1111",
				proto:    "tcp",
				addr:     mockContainer.config.addr,
				schema:   "ledger",
			},
			interval:   time.Second,
			reps:       3,
			wantErrStr: "ping MySQL",
		},
		{
			name:       "success case: correct user, password, proto, addr, schema",
			shouldFail: false,
			config:     mockContainer.config,
			interval:   time.Second,
			reps:       3,
			wantErrStr: "",
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
			elapsed := time.Since(start).Round(time.Second)

			// then
			if tt.shouldFail {
				wantDur := tt.interval * time.Duration(tt.reps)
				if elapsed != wantDur {
					t.Errorf("RetryPing() elapsed = %v, wantDur %v", elapsed, wantDur)
				}

				if got := err.Error(); !strings.HasPrefix(got, tt.wantErrStr) {
					t.Errorf("RetryPing() error = %v, wantErrStr %s", err, tt.wantErrStr)
				}

				return
			}

			if err != nil {
				t.Errorf("RetryPing() error = %v, wantErrStr %s", err, tt.wantErrStr)
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
				t.Errorf("Close() error = %T, wantErrStr %T", err, tt.wantErr)
			}
		})
	}
}
