package e2e

import (
	"context"
	"net/http/httptest"
	"os"
	"regexp"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/minoritea/chat/config"
	"github.com/minoritea/chat/domain/session"
	"github.com/minoritea/chat/domain/user"
	"github.com/minoritea/chat/resource"
	"github.com/minoritea/chat/router"
	"github.com/playwright-community/playwright-go"
)

func TestE2E(t *testing.T) {
	container, server, err := setupServer()
	t.Cleanup(server.Close)

	browser, err := setupBrowser()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { browser.Close() })

	assertion := playwright.NewPlaywrightAssertions()
	t.Run("top page requires session and redirects to auth page", func(t *testing.T) {
		bctx, err := browser.NewContext()
		if err != nil {
			t.Fatal(err)
		}
		page, err := bctx.NewPage()
		if err != nil {
			t.Fatal(err)
		}

		res, err := page.Goto(server.URL)
		if err != nil {
			t.Fatal(err)
		}

		res.Finished()
		err = assertion.Page(page).ToHaveURL(server.URL + "/auth")
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("authorized user can access top page", func(t *testing.T) {
		bctx, err := browser.NewContext()
		if err != nil {
			t.Fatal(err)
		}
		page, err := bctx.NewPage()
		if err != nil {
			t.Fatal(err)
		}
		err = setSessionToContext(*container, server, page.Context(), "test-user")
		if err != nil {
			t.Fatal(err)
		}

		res, err := page.Goto(server.URL)
		if err != nil {
			t.Fatal(err)
		}

		res.Finished()
		err = assertion.Page(page).ToHaveURL(server.URL + "/")
		if err != nil {
			t.Fatal(err)
		}

		err = assertion.Locator(page.Locator("div#messages")).ToHaveCount(1)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("authorized user can post message", func(t *testing.T) {
		bctx, err := browser.NewContext()
		if err != nil {
			t.Fatal(err)
		}
		page, err := bctx.NewPage()
		if err != nil {
			t.Fatal(err)
		}
		err = setSessionToContext(*container, server, page.Context(), "test-user")
		if err != nil {
			t.Fatal(err)
		}

		res, err := page.Goto(server.URL)
		if err != nil {
			t.Fatal(err)
		}

		res.Finished()
		err = assertion.Page(page).ToHaveURL(server.URL + "/")
		if err != nil {
			t.Fatal(err)
		}

		err = assertion.Locator(page.Locator("form")).ToHaveCount(1)
		if err != nil {
			t.Fatal(err)
		}

		err = page.Locator("form > input[type=text]").Fill("test message")
		if err != nil {
			t.Fatal(err)
		}

		err = page.Locator("form > input[type=submit]").Click()
		if err != nil {
			t.Fatal(err)
		}

		err = assertion.Locator(page.Locator("div#messages > .message")).ToHaveCount(1)
		if err != nil {
			t.Fatal(err)
		}

		matchText := regexp.MustCompile(`test message`)
		err = assertion.Locator(page.Locator("div#messages > .message")).ToHaveText(matchText)
		if err != nil {
			t.Fatal(err)
		}
	})
}

func setupServer() (*resource.Container, *httptest.Server, error) {
	conf := config.Config{
		SessionSecret:  "test",
		DatabaseDriver: "sqlite3",
		DatabasePath:   ":memory:",
		CSRFSecret:     "test",
	}
	container, err := resource.New(conf)
	if err != nil {
		return nil, nil, err
	}
	schema, err := os.ReadFile("../../database/schema.sql")
	if err != nil {
		return nil, nil, err
	}
	_, err = container.DB().Exec(string(schema))
	if err != nil {
		return nil, nil, err
	}
	server := httptest.NewServer(router.New(*container))
	return container, server, nil
}

func setupBrowser() (playwright.Browser, error) {
	playwright.Install(&playwright.RunOptions{
		Browsers: []string{"chromium"},
		Verbose:  true,
	})

	pw, err := playwright.Run()
	if err != nil {
		return nil, err
	}

	return pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(true),
	})
}

func setSessionToContext(container resource.Container, server *httptest.Server, bctx playwright.BrowserContext, userName string) error {
	dbUser, err := user.FindOrCreateUser(context.Background(), container, userName)
	if err != nil {
		return err
	}

	dbSession, err := session.PerpetuateSession(context.Background(), container, dbUser.ID)
	if err != nil {
		return err
	}

	r := httptest.NewRequest("GET", server.URL, nil)
	w := httptest.NewRecorder()
	s := session.MustNew(container, r)
	session.SetSessionID(s, dbSession.ID)
	session.MustSave(s, r, w)
	cookies := w.Result().Cookies()

	var pcookies []playwright.OptionalCookie
	for _, c := range cookies {
		pcookies = append(pcookies, playwright.OptionalCookie{
			Name:   c.Name,
			Value:  c.Value,
			Path:   playwright.String("/"),
			Domain: playwright.String("127.0.0.1"),
		})
	}

	return bctx.AddCookies(pcookies)
}
