package user_test

import (
	"context"
	"testing"

	"github.com/minoritea/chat/database"
	"github.com/minoritea/chat/domain/user"
	dbmock "github.com/minoritea/chat/test/mock/database"
	containerstub "github.com/minoritea/chat/test/stub/container"
	"go.uber.org/mock/gomock"
)

func TestFindOrCreateUser(t *testing.T) {
	t.Run("the target user exists", func(t *testing.T) {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		querier := dbmock.NewMockQuerier(ctrl)
		c := containerstub.NewQuerierContainer(querier)
		u := database.User{ID: "1", Account: "account"}
		querier.EXPECT().GetUserByAccount(ctx, "account").Return(u, nil)
		foundUser, err := user.FindOrCreateUser(ctx, c, "account")
		if err != nil {
			t.Fatal(err)
		}
		if *foundUser != u {
			t.Errorf("foundUser = %+v, want %v", foundUser, u)
		}
	})

	t.Run("the target user does not exist", func(t *testing.T) {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		querier := dbmock.NewMockQuerier(ctrl)
		c := containerstub.NewQuerierContainer(querier)
		querier.EXPECT().GetUserByAccount(ctx, "account").Return(database.User{}, database.RecordNotFound)
		var u database.User
		querier.EXPECT().CreateUser(ctx, gomock.Any()).DoAndReturn(func(ctx context.Context, params database.CreateUserParams) (database.User, error) {
			u = database.User{ID: params.ID, Account: params.Account}
			return u, nil
		})
		foundUser, err := user.FindOrCreateUser(ctx, c, "account")
		if err != nil {
			t.Fatal(err)
		}
		if *foundUser != u {
			t.Errorf("foundUser = %+v, want %v", foundUser, u)
		}
	})
}
