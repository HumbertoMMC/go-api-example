package handlers

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"go.uber.org/zap"
)

func TestHandler(t *testing.T) {
	logger := zap.NewExample()

	h := NewHandler(logger)

	t.Run("Test post method", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodPost, "http://localhost:8000/", strings.NewReader(`{"name":"bob"}`))

		w := httptest.NewRecorder()

		h.ServeHTTP(w, req)

		// We should get a good status code
		if want, got := http.StatusOK, w.Result().StatusCode; want != got {
			t.Fatalf("expected a %d, instead got: %d", want, got)
		}

		// We should get header Content-Type as application/json
		if want, got := "application/json", w.Result().Header.Get("Content-Type"); want != got {
			t.Fatalf("expected a %s, instead got: %s", want, got)
		}

		//
		b, err := io.ReadAll(w.Result().Body)
		if err != nil {
			log.Fatalln(err)
		}

		got := strings.TrimSpace(string(b))

		// We should get response body as application/json {"message":"Hello, bob"}
		if want := `{"message":"Hello, bob"}`; want != got {
			t.Fatalf("expected a %s, instead got: %s", want, got)
		}

	})

	t.Run("Test get method", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "http://localhost:8000/?key=bob", nil)

		w := httptest.NewRecorder()

		h.ServeHTTP(w, req)

		// We should get a good status code
		if want, got := http.StatusOK, w.Result().StatusCode; want != got {
			t.Fatalf("expected a %d, instead got: %d", want, got)
		}

		// We should get header Content-Type as application/json
		if want, got := "application/json", w.Result().Header.Get("Content-Type"); want != got {
			t.Fatalf("expected a %s, instead got: %s", want, got)
		}

		//
		b, err := io.ReadAll(w.Result().Body)
		if err != nil {
			log.Fatalln(err)
		}

		got := strings.TrimSpace(string(b))

		// We should get response body as application/json {"message":"Url Param 'key' is: bob"}
		if want := `{"message":"Url Query 'key' is: bob"}`; want != got {
			t.Fatalf("expected a %s, instead got: %s", want, got)
		}

	})
}
