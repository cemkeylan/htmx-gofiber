package htmx

import (
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWrite(t *testing.T) {
	assert := assert.New(t)
	r := require.New(t)

	app := fiber.New()
	app.Get("/", func(c fiber.Ctx) error {
		return NewResponse().
			StatusCode(StatusStopPolling).
			Location("/profiles").
			Redirect("/pull").
			PushURL("/push").
			Refresh(true).
			ReplaceURL("/water").
			Retarget("#world").
			Reselect("#hello").
			AddTrigger(Trigger("myEvent")).
			Reswap(SwapInnerHTML.ShowOn("#swappy", Top)).
			Write(c)
	})

	req := httptest.NewRequest("GET", "http://localhost", nil)
	resp, err := app.Test(req)
	r.NoError(err)
	assert.Equal(StatusStopPolling, resp.StatusCode, "wrong error code")

	expectedHeaders := map[string]string{
		HeaderTrigger:    "myEvent",
		HeaderLocation:   "/profiles",
		HeaderRedirect:   "/pull",
		HeaderPushURL:    "/push",
		HeaderRefresh:    "true",
		HeaderReplaceUrl: "/water",
		HeaderRetarget:   "#world",
		HeaderReselect:   "#hello",
		HeaderReswap:     "innerHTML show:#swappy:top",
	}

	for k, v := range expectedHeaders {
		assert.Equal(v, resp.Header.Get(k), "wrong value for header %q", k)
	}
}

func TestRenderHTML(t *testing.T) {
	text := `hello world!`
	assert := assert.New(t)
	r := require.New(t)

	app := fiber.New()
	app.Get("/", func(c fiber.Ctx) error {
		_, err := NewResponse().Location("/conversation/message").RenderHTML(c, template.HTML(text))
		return errors.Wrap(err, "an error occured writing html")
	})

	req := httptest.NewRequest("GET", "http://localhost", nil)
	resp, err := app.Test(req)
	r.NoError(err)
	body, err := io.ReadAll(resp.Body)
	r.NoError(err)
	assert.Equal("/conversation/message", resp.Header.Get(HeaderLocation), "wrong value for header %q", HeaderLocation)
	assert.Equal(string(body), text, "wrong response body")
}

func TestMustRenderHTML(t *testing.T) {
	text := `hello world!`

	app := fiber.New()
	app.Get("/", func(c fiber.Ctx) error {
		NewResponse().MustRenderHTML(c, template.HTML(text))
		return nil
	})

	req := httptest.NewRequest("GET", "http://localhost", nil)
	_, _ = app.Test(req)
}

type mockResponseWriter struct {
	body       []byte
	statusCode int
	header     http.Header
}

func newMockResponseWriter() *mockResponseWriter {
	return &mockResponseWriter{
		header: http.Header{},
	}
}

func (mrw *mockResponseWriter) Header() http.Header {
	return mrw.header
}

func (mrw *mockResponseWriter) Write(b []byte) (int, error) {
	mrw.body = append(mrw.body, b...)
	return 0, nil
}

func (mrw *mockResponseWriter) WriteHeader(statusCode int) {
	mrw.statusCode = statusCode
}
