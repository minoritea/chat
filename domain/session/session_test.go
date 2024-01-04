package session_test

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/sessions"
	"github.com/minoritea/chat/database"
	"github.com/minoritea/chat/domain/session"
	dbmock "github.com/minoritea/chat/test/mock/database"
	sessionmock "github.com/minoritea/chat/test/mock/session"
	containerstub "github.com/minoritea/chat/test/stub/container"
	"go.uber.org/mock/gomock"
)

func TestStoreNewSession(t *testing.T) {
	ctrl := gomock.NewController(t)
	store := sessionmock.NewMockStore(ctrl)
	querier := dbmock.NewMockQuerier(ctrl)
	ctx := context.Background()
	c := struct {
		containerstub.QuerierContainer
		containerstub.SessionStoreContainer
	}{
		containerstub.NewQuerierContainer(querier),
		containerstub.NewSessionStoreContainer(store),
	}
	r := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	userID := database.NewID()
	var sess database.Session
	querier.EXPECT().CreateSession(ctx, gomock.Any()).DoAndReturn(func(ctx context.Context, params database.CreateSessionParams) (database.Session, error) {
		sess = database.Session{ID: params.ID, UserID: params.UserID}
		return sess, nil
	})
	storedSession := sessions.NewSession(store, session.SessionName)
	store.EXPECT().New(r, session.SessionName).Return(storedSession, nil)
	store.EXPECT().Save(r, w, storedSession).Return(nil)
	err := session.StoreNewSession(ctx, c, w, r, userID)
	if err != nil {
		t.Fatal(err)
	}
	sessionID, ok := storedSession.Values["session_id"].(string)
	if !ok {
		t.Errorf("session_id is not set")
	}

	if sess.ID != sessionID {
		t.Errorf("sess.ID = %s, want %s", sess.ID, sessionID)
	}

	if sess.UserID != userID {
		t.Errorf("sess.UserID = %s, want %s", sess.UserID, userID)
	}
}
