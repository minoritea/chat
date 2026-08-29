// Package user provides domain logic for users.
package user

import (
	"context"

	"github.com/minoritea/chat/database"
)

// Container provides a database.Querier.
type Container interface {
	Querier() database.Querier
}

// User is an alias for database.User.
type User = database.User

// FindOrCreateUser returns the user with the given account, creating it if it does not exist.
func FindOrCreateUser(ctx context.Context, c Container, account string) (*User, error) {
	q := c.Querier()
	user, err := q.GetUserByAccount(ctx, account)
	if err == nil {
		return &user, nil
	}
	if !database.IsRecordNotFound(err) {
		return nil, err
	}
	user, err = q.CreateUser(ctx, database.CreateUserParams{
		ID:      database.NewID(),
		Account: account,
	})
	return &user, err
}

type userKey struct{}

// FromContext returns the user stored in ctx. It panics if the user is not set.
func FromContext(ctx context.Context) *User {
	user, ok := ctx.Value(userKey{}).(User)
	if !ok {
		panic("user not found in context")
	}
	return &user
}

// SetToContext stores the user in ctx to mark the user as logged in.
func SetToContext(ctx context.Context, user *User) context.Context {
	return context.WithValue(ctx, userKey{}, *user)
}
