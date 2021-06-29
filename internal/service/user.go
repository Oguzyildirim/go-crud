package service

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/trace"

	"github.com/Oguzyildirim/go-crud/internal"
)

// UserRepository defines the datastore handling persisting User records
type UserRepository interface {
	Create(ctx context.Context, name string, lastname string, username string, country string) (internal.User, error)
	Delete(ctx context.Context, id string) error
	Find(ctx context.Context, id string) (internal.User, error)
	FindByCountry(ctx context.Context, country string) ([]internal.User, error)
	Update(ctx context.Context, id string, name string, lastname string, username string, country string) error
}

// User defines the application service in charge of interacting with Users
type User struct {
	repo UserRepository
}

// NewUser
func NewUser(repo UserRepository) *User {
	return &User{
		repo: repo,
	}
}

// Create stores a new record
func (u *User) Create(ctx context.Context, name string, lastname string, username string, country string) (internal.User, error) {
	ctx, span := trace.SpanFromContext(ctx).TracerProvider().Tracer("UserTracer").Start(ctx, "User.Create")
	defer span.End()

	user, err := u.repo.Create(ctx, name, lastname, username, country)
	if err != nil {
		return internal.User{}, fmt.Errorf("repo create: %w", err)
	}

	return user, nil
}

// Delete removes an existing User from the datastore
func (u *User) Delete(ctx context.Context, id string) error {
	ctx, span := trace.SpanFromContext(ctx).TracerProvider().Tracer("UserTracer").Start(ctx, "User.Delete")
	defer span.End()

	if err := u.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("repo delete: %w", err)
	}

	return nil
}

// Find gets an existing User from the datastore
func (u *User) Find(ctx context.Context, id string) (internal.User, error) {
	ctx, span := trace.SpanFromContext(ctx).TracerProvider().Tracer("UserTracer").Start(ctx, "User.Find")
	defer span.End()

	user, err := u.repo.Find(ctx, id)
	if err != nil {
		return internal.User{}, fmt.Errorf("repo find: %w", err)
	}

	return user, nil
}

// FindByCountry gets an existing User from the datastore
func (u *User) FindByCountry(ctx context.Context, country string) ([]internal.User, error) {
	ctx, span := trace.SpanFromContext(ctx).TracerProvider().Tracer("UserTracer").Start(ctx, "User.FindByCountry")
	defer span.End()

	user, err := u.repo.FindByCountry(ctx, country)
	if err != nil {
		return []internal.User{}, fmt.Errorf("repo find: %w", err)
	}

	return user, nil
}

// Update updates an existing User in the datastore
func (u *User) Update(ctx context.Context, id string, name string, lastname string, username string, country string) error {
	ctx, span := trace.SpanFromContext(ctx).TracerProvider().Tracer("UserTracer").Start(ctx, "User.Update")
	defer span.End()

	if err := u.repo.Update(ctx, id, name, lastname, username, country); err != nil {
		return fmt.Errorf("repo update: %w", err)
	}

	return nil
}
