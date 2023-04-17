package main

import (
	"context"
	"encoding/json"
	"net/http"
)

type Server struct {
	svc Service
}

func NewServer(svc Service) *Server {
	return &Server{
		svc: svc,
	}
}

func (s *Server) Start(listenAddress string) error {
	http.HandleFunc("/chucknorrisfact", s.handleGetFact)
	return http.ListenAndServe(listenAddress, nil)
}

func (s *Server) handleGetFact(response http.ResponseWriter, request *http.Request) {
	fact, err := s.svc.GetFact(context.Background())
	if err != nil {
		writeJSON(response, http.StatusUnprocessableEntity, map[string]any{"error": err.Error()})
		return
	}
	writeJSON(response, http.StatusOK, fact)
}

func writeJSON(response http.ResponseWriter, status int, value any) error {
	response.WriteHeader(status)
	response.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(response).Encode(value)
}
