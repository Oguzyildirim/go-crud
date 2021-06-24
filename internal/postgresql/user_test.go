package postgresql_test

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net"
	"net/url"
	"runtime"
	"testing"
	"time"

	migrate "github.com/golang-migrate/migrate/v4"
	migratepostgres "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/google/go-cmp/cmp"
	_ "github.com/jackc/pgx/v4/stdlib" // to initialize "pgx"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"

	"github.com/Oguzyildirim/go-crud/internal"
	"github.com/Oguzyildirim/go-crud/internal/postgresql"
)

func TestUser_Create(t *testing.T) {
	t.Parallel()

	t.Run("Create: OK", func(t *testing.T) {
		t.Parallel()

		user, err := postgresql.NewUser(newDB(t)).Create(context.Background(), "name", "lastname", "username", "UK")
		if err != nil {
			t.Fatalf("expected no error, got %s", err)
		}

		if user.ID == "" {
			t.Fatalf("expected valid record, got empty value")
		}
	})
}

func TestUser_Delete(t *testing.T) {
	t.Parallel()

	t.Run("Delete: OK", func(t *testing.T) {
		t.Parallel()

		store := postgresql.NewUser(newDB(t))

		createdUser, err := store.Create(context.Background(), "name", "lastname", "username", "UK")
		if err != nil {
			t.Fatalf("expected no error, got %s", err)
		}

		if err := store.Delete(context.Background(), createdUser.ID); err != nil {
			t.Fatalf("expected no error, got %s", err)
		}

		if _, err = store.Find(context.Background(), createdUser.ID); !errors.Is(err, sql.ErrNoRows) {
			t.Fatalf("expected no error, got %s", err)
		}
	})

	t.Run("Update: ERR uuid", func(t *testing.T) {
		t.Parallel()

		err := postgresql.NewUser(newDB(t)).Delete(context.Background(), "x")

		if err == nil {
			t.Fatalf("expected error, got not value")
		}

		var ierr *internal.Error
		if !errors.As(err, &ierr) || ierr.Code() != internal.ErrorCodeInvalidArgument {
			t.Fatalf("expected %T error, got %T : %v", ierr, err, err)
		}
	})

	t.Run("Delete: ERR not found", func(t *testing.T) {
		t.Parallel()

		err := postgresql.NewUser(newDB(t)).Delete(context.Background(), "44633fe3-b039-4fb3-a35f-a57fe3c906c7")

		var ierr *internal.Error
		if !errors.As(err, &ierr) || ierr.Code() != internal.ErrorCodeNotFound {
			t.Fatalf("expected %T error, got %T : %v", ierr, err, err)
		}
	})
}

func TestUser_Find(t *testing.T) {
	t.Parallel()

	t.Run("Find: OK", func(t *testing.T) {
		t.Parallel()

		store := postgresql.NewUser(newDB(t))

		originalUser, err := store.Create(context.Background(), "name", "lastname", "username", "UK")
		if err != nil {
			t.Fatalf("expected no error, got %s", err)
		}

		actualUser, err := store.Find(context.Background(), originalUser.ID)
		if err != nil {
			t.Fatalf("expected no error, got %s", err)
		}

		if !cmp.Equal(originalUser, actualUser) {
			t.Fatalf("expected result does not match: %s", cmp.Diff(originalUser, actualUser))
		}
	})

	t.Run("Find: ERR uuid", func(t *testing.T) {
		t.Parallel()

		_, err := postgresql.NewUser(newDB(t)).Find(context.Background(), "x")
		if err == nil {
			t.Fatalf("expected error, got not value")
		}

		var ierr *internal.Error
		if !errors.As(err, &ierr) || ierr.Code() != internal.ErrorCodeInvalidArgument {
			t.Fatalf("expected %T error, got %T : %v", ierr, err, err)
		}
	})

	t.Run("Find: ERR not found", func(t *testing.T) {
		t.Parallel()

		_, err := postgresql.NewUser(newDB(t)).Find(context.Background(), "44633fe3-b039-4fb3-a35f-a57fe3c906c7")
		if err == nil {
			t.Fatalf("expected error, got not value")
		}

		var ierr *internal.Error
		if !errors.As(err, &ierr) || ierr.Code() != internal.ErrorCodeNotFound {
			t.Fatalf("expected %T error, got %T : %v", ierr, err, err)
		}
	})
}

func TestUser_FindByCountry(t *testing.T) {
	t.Parallel()

	t.Run("Find: OK", func(t *testing.T) {
		t.Parallel()

		store := postgresql.NewUser(newDB(t))

		userOne, err := store.Create(context.Background(), "name", "lastname", "username", "UK")
		userTwo, err := store.Create(context.Background(), "name", "lastname", "username", "UK")
		userThird, err := store.Create(context.Background(), "name", "lastname", "username", "UK")
		userFourth, err := store.Create(context.Background(), "name", "lastname", "username", "UK")
		if err != nil {
			t.Fatalf("expected no error, got %s", err)
		}
		var testUsers []internal.User
		testUsers = append(testUsers, userOne)
		testUsers = append(testUsers, userTwo)
		testUsers = append(testUsers, userThird)
		testUsers = append(testUsers, userFourth)
		users, err := store.FindByCountry(context.Background(), "UK")
		if err != nil {
			t.Fatalf("expected no error, got %s", err)
		}

		if !cmp.Equal(len(users), len(testUsers)) {
			t.Fatalf("expected result does not match: %s", cmp.Diff(users, testUsers))
		}
	})
}

func TestUser_Update(t *testing.T) {
	t.Parallel()

	t.Run("Update: OK", func(t *testing.T) {
		t.Parallel()

		store := postgresql.NewUser(newDB(t))

		originalUser, err := store.Create(context.Background(), "name", "lastname", "username", "UK")
		if err != nil {
			t.Fatalf("expected no error, got %s", err)
		}

		originalUser.Name = "changed"
		originalUser.Lastname = "changed"
		originalUser.Username = "changed"
		originalUser.Country = "changed"

		if err := store.Update(context.Background(),
			originalUser.ID,
			originalUser.Name,
			originalUser.Lastname,
			originalUser.Username,
			originalUser.Country); err != nil {
			t.Fatalf("expected no error, got %s", err)
		}

		actualUser, err := store.Find(context.Background(), originalUser.ID)
		if err != nil {
			t.Fatalf("expected no error, got %s", err)
		}

		if !cmp.Equal(originalUser, actualUser) {
			t.Fatalf("expected result does not match: %s", cmp.Diff(originalUser, actualUser))
		}
	})

	t.Run("Update: ERR uuid", func(t *testing.T) {
		t.Parallel()

		err := postgresql.NewUser(newDB(t)).Update(context.Background(),
			"x",
			"",
			"",
			"",
			"")
		if err == nil {
			t.Fatalf("expected error, got not value")
		}

		var ierr *internal.Error
		if !errors.As(err, &ierr) || ierr.Code() != internal.ErrorCodeInvalidArgument {
			t.Fatalf("expected %T error, got %T : %v", ierr, err, err)
		}
	})

	t.Run("Update: ERR not found", func(t *testing.T) {
		t.Parallel()

		err := postgresql.NewUser(newDB(t)).Update(context.Background(),
			"44633fe3-b039-4fb3-a35f-a57fe3c906c7",
			"",
			"",
			"",
			"")
		if err == nil {
			t.Fatalf("expected error, got not value")
		}

		var ierr *internal.Error
		if !errors.As(err, &ierr) || ierr.Code() != internal.ErrorCodeNotFound {
			t.Fatalf("expected %T error, got %T : %v", ierr, err, err)
		}
	})
}

func newDB(tb testing.TB) *sql.DB {
	dsn := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword("username", "password"),
		Path:   "todo",
	}

	q := dsn.Query()
	q.Add("sslmode", "disable")

	dsn.RawQuery = q.Encode()

	//-

	pool, err := dockertest.NewPool("")
	if err != nil {
		tb.Fatalf("Couldn't connect to docker: %s", err)
	}

	pool.MaxWait = 10 * time.Second

	pw, _ := dsn.User.Password()

	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "12.5-alpine",
		Env: []string{
			fmt.Sprintf("POSTGRES_USER=%s", dsn.User.Username()),
			fmt.Sprintf("POSTGRES_PASSWORD=%s", pw),
			fmt.Sprintf("POSTGRES_DB=%s", dsn.Path),
		},
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{
			Name: "no",
		}
	})

	if err != nil {
		tb.Fatalf("Couldn't start resource: %s", err)
	}

	resource.Expire(60)

	tb.Cleanup(func() {
		if err := pool.Purge(resource); err != nil {
			tb.Fatalf("Couldn't purge container: %v", err)
		}
	})

	dsn.Host = fmt.Sprintf("%s:5432", resource.Container.NetworkSettings.IPAddress)
	if runtime.GOOS == "darwin" { // MacOS-specific
		dsn.Host = net.JoinHostPort(resource.GetBoundIP("5432/tcp"), resource.GetPort("5432/tcp"))
	}

	db, err := sql.Open("pgx", dsn.String())
	if err != nil {
		tb.Fatalf("Couldn't open DB: %s", err)
	}

	tb.Cleanup(func() {
		if err := db.Close(); err != nil {
			tb.Fatalf("Couldn't close DB: %s", err)
		}
	})

	if err := pool.Retry(func() (err error) {
		return db.Ping()
	}); err != nil {
		tb.Fatalf("Couldn't ping DB: %s", err)
	}

	//-

	instance, err := migratepostgres.WithInstance(db, &migratepostgres.Config{})
	if err != nil {
		tb.Fatalf("Couldn't migrate (1): %s", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://../../db/migrations/", "postgres", instance)
	if err != nil {
		tb.Fatalf("Couldn't migrate (2): %s", err)
	}

	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		tb.Fatalf("Couldnt' migrate (3): %s", err)
	}

	return db
}
