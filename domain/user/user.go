package user

import (
	"context"
	"crypto/sha256"
	"errors"

	"github.com/minoritea/chat/container"
	"github.com/minoritea/chat/database"
)

type Container = container.Container
type User = database.User

func RegisterUser(ctx context.Context, c *Container, accountName, password string) (*User, error) {
	passwordHash := sha256.Sum256([]byte(password))
	q := c.GetQueries()
	user, err := q.CreateUser(ctx, database.CreateUserParams{
		ID:           database.NewID(),
		AccountName:  accountName,
		PasswordHash: string(passwordHash[:]),
	})
	return &user, err
}

type userKey struct{}

func FromContext(ctx context.Context) *User {
	user, ok := ctx.Value(userKey{}).(User)
	if !ok {
		panic("user not found in context")
	}
	return &user
}

func SetToContext(ctx context.Context, user *User) context.Context {
	return context.WithValue(ctx, userKey{}, *user)
}

var ErrPasswordMismatch = errors.New("password mismatch")

func GetByAccoutNameAndPassword(ctx context.Context, c *Container, accountName, password string) (*User, error) {
	q := c.GetQueries()
	user, err := q.GetUserByAccountName(ctx, accountName)
	if err != nil {
		return nil, err
	}
	passwordHash := sha256.Sum256([]byte(password))
	if user.PasswordHash != string(passwordHash[:]) {
		return nil, ErrPasswordMismatch
	}
	return &user, nil
}
