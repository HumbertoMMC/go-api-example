package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

type PostRequest struct {
	Name string `json:"name"`
}

type RestResponse struct {
	Message string `json:"message"`
}

type ClientI interface {
	Greeter(name string) string
}

type Handler struct {
	Client ClientI
	Logger *zap.Logger
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.postRequest(w, r)
	case http.MethodGet:
		h.getRequest(w, r)
	}

}

// POST Request
func (h *Handler) postRequest(w http.ResponseWriter, r *http.Request) {
	var postRequest PostRequest

	err := json.NewDecoder(r.Body).Decode(&postRequest)
	if err != nil {
		ErrorHandler(h.Logger, w, ErrBadRquest)
		return
	}

	//Logging
	h.Logger.Info("POST method called", zap.String("with", postRequest.Name))

	resp := RestResponse{
		Message: h.Client.Greeter(postRequest.Name),
	}

	w.Header().Add("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(resp)

	if err != nil {
		ErrorHandler(h.Logger, w, ErrInternalServerError)
		return
	}
}

// GET Request
func (h *Handler) getRequest(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["key"]

	if !ok || len(keys[0]) < 1 {
		ErrorHandler(h.Logger, w, ErrMissingKey)
		return
	}

	// Query()["key"] will return an array of items,
	// we only want the single item.
	key := keys[0]

	//Logging
	h.Logger.Info("GET method called with", zap.String("Url Query key", key))

	resp := RestResponse{
		Message: "Url Query 'key' is: " + string(key),
	}

	w.Header().Add("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(resp)

	if err != nil {
		ErrorHandler(h.Logger, w, ErrInternalServerError)
		return
	}
}

type ClientDB struct {
	Users []string
}

func NewHandler(logger *zap.Logger) Handler {
	clientDB := ClientDB{
		Users: []string{"Jim", "Bob", "Peter"},
	}
	return Handler{
		Client: clientDB,
		Logger: logger,
	}
}

func (c ClientDB) Greeter(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}
