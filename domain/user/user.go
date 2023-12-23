package user

import (
	"context"
	"crypto/sha256"

	"github.com/minoritea/chat/container"
	"github.com/minoritea/chat/database"
)

type Container = container.Container
type User = database.User

func RegisterUser(c *Container, accountName, password string) (user *User, err error) {
	passwordHash := sha256.Sum256([]byte(password))
	q := c.GetQuerier()
	err = q.CreateUser(accountName, string(passwordHash[:]))
	if err != nil {
		return nil, err
	}
	return q.GetUserByAccountName(accountName)
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
