package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/MLaskun/go-workshop/storage"
	"github.com/MLaskun/go-workshop/types"
	"github.com/go-chi/chi/v5"
)

type APIServer struct {
	listenAddr string
	store      storage.Storage
}

func NewServer(listenAddr string, store storage.Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *APIServer) Run() error {
	r := chi.NewRouter()
	r.Method("GET", "/client/{id}", makeHTTPHandlerFunc(s.HandleGetClientByID))
	r.Method("POST", "/client", makeHTTPHandlerFunc(s.HandleCreateClient))
	r.Method("DELETE", "/client/{id}", makeHTTPHandlerFunc(s.HandleDeleteClient))

	return http.ListenAndServe(s.listenAddr, r)
}

func (s *APIServer) HandleGetClientByID(w http.ResponseWriter, r *http.Request) error {
	ctx := context.WithValue(r.Context(), requestKey, r)

	id, err := getID(ctx)
	if err != nil {
		return err
	}

	client, err := s.store.GetClientByID(id)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, client)
}

func (s *APIServer) HandleCreateClient(w http.ResponseWriter, r *http.Request) error {
	req := &types.CreateClientRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}

	client := types.NewClient(req.FirstName, req.LastName, req.Email, req.PhoneNO)
	if err := s.store.CreateClient(client); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusCreated, client)
}

func (s *APIServer) HandleDeleteClient(w http.ResponseWriter, r *http.Request) error {
	ctx := context.WithValue(r.Context(), requestKey, r)
	id, err := getID(ctx)
	if err != nil {
		return err
	}

	if err := s.store.DeleteClient(id); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusNoContent, map[string]int{"deleted": id})
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}

func makeHTTPHandlerFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

type contextKey string

const requestKey = contextKey("request")

func getID(ctx context.Context) (int, error) {
	r := ctx.Value(requestKey).(*http.Request)

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return id, fmt.Errorf("invalid id given %s", idStr)
	}
	return id, nil
}
