package user

import (
	"context"

	"github.com/minoritea/chat/database"
	"github.com/minoritea/chat/resource"
)

type Container = resource.Container
type User = database.User

func FindOrCreateUser(ctx context.Context, c Container, account string) (*User, error) {
	q := c.Queries()
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
