package postgresql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"github.com/Oguzyildirim/go-crud/internal"
)

// User represents the repository used for interacting with User records
type User struct {
	q *Queries
}

// NewUser instantiates the User repository
func NewUser(db *sql.DB) *User {
	return &User{
		q: New(db),
	}
}

// Create inserts a new ıser record
func (t *User) Create(ctx context.Context, name string, lastname string, username string, country string) (internal.User, error) {
	ctx, span := trace.SpanFromContext(ctx).Tracer().Start(ctx, "User.Create")
	span.SetAttributes(attribute.String("db.system", "postgresql"))
	defer span.End()
	id, err := t.q.InsertUser(ctx, InsertUserParams{
		Name:     name,
		Lastname: lastname,
		Username: username,
		Country:  country,
	})
	if err != nil {
		return internal.User{}, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "insert user")
	}

	return internal.User{
		ID:       id.String(),
		Name:     name,
		Lastname: lastname,
		Username: username,
		Country:  country,
	}, nil
}

// Delete deletes the existing record matching the id.
func (t *User) Delete(ctx context.Context, id string) error {
	ctx, span := trace.SpanFromContext(ctx).Tracer().Start(ctx, "User.Delete")
	span.SetAttributes(attribute.String("db.system", "postgresql"))
	defer span.End()
	val, err := uuid.Parse(id)
	if err != nil {
		return internal.WrapErrorf(err, internal.ErrorCodeInvalidArgument, "invalid uuid")
	}
	_, err = t.q.DeleteUser(ctx, val)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return internal.WrapErrorf(err, internal.ErrorCodeNotFound, "user not found")
		}

		return internal.WrapErrorf(err, internal.ErrorCodeUnknown, "delete user")
	}
	return nil
}

// Find returns the requested User by searching its id.
func (t *User) Find(ctx context.Context, id string) (internal.User, error) {
	ctx, span := trace.SpanFromContext(ctx).Tracer().Start(ctx, "User.Find")
	span.SetAttributes(attribute.String("db.system", "postgresql"))
	defer span.End()
	val, err := uuid.Parse(id)
	if err != nil {
		return internal.User{}, internal.WrapErrorf(err, internal.ErrorCodeInvalidArgument, "invalid uuid")
	}
	res, err := t.q.SelectUser(ctx, val)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return internal.User{}, internal.WrapErrorf(err, internal.ErrorCodeNotFound, "user not found")
		}

		return internal.User{}, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "select User")
	}
	return internal.User{
		ID:       res.ID.String(),
		Name:     res.Name,
		Lastname: res.Lastname,
		Username: res.Username,
		Country:  res.Country,
	}, nil
}

// Find returns the requested User by searching its id.
func (t *User) FindByCountry(ctx context.Context, country string) ([]internal.User, error) {
	ctx, span := trace.SpanFromContext(ctx).Tracer().Start(ctx, "User.Find")
	span.SetAttributes(attribute.String("db.system", "postgresql"))
	defer span.End()
	res, err := t.q.SelectUsersByCountry(ctx, country)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []internal.User{}, internal.WrapErrorf(err, internal.ErrorCodeNotFound, "user not found")
		}

		return []internal.User{}, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "select User")
	}
	var users []internal.User
	for _, value := range res {
		user := internal.User{
			ID:       value.ID.String(),
			Name:     value.Name,
			Lastname: value.Lastname,
			Username: value.Username,
			Country:  value.Country,
		}
		users = append(users, user)
	}
	return users, nil
}

// Update updates the existing record with new values.
func (t *User) Update(ctx context.Context, id string, name string, lastname string, username string, country string) error {
	ctx, span := trace.SpanFromContext(ctx).Tracer().Start(ctx, "User.Update")
	span.SetAttributes(attribute.String("db.system", "postgresql"))
	defer span.End()
	val, err := uuid.Parse(id)
	if err != nil {
		return internal.WrapErrorf(err, internal.ErrorCodeInvalidArgument, "invalid uuid")
	}
	if _, err := t.q.UpdateUser(ctx, UpdateUserParams{
		ID:       val,
		Name:     name,
		Lastname: lastname,
		Username: username,
		Country:  country,
	}); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return internal.WrapErrorf(err, internal.ErrorCodeNotFound, "user not found")
		}
		return internal.WrapErrorf(err, internal.ErrorCodeUnknown, "update User")
	}
	return nil
}
