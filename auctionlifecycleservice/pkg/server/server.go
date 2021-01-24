package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	gmux "github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/saad-karim/kaleido-project-backend/auctionlifecycleservice/pkg/api"
)

type Service interface {
	Start(*api.StartRequest) (*api.StartAuctionResponse, error)
	Close(*api.CloseRequest) (*api.CloseAuctionResponse, error)
}

type HTTPServer interface {
	ListenAndServe() error
}

type Router interface {
	HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) *gmux.Route
}

type Server struct {
	Service    Service
	HTTPServer HTTPServer
	Router     Router
}

func (s *Server) RegisterEndpoints() {
	s.Router.HandleFunc("/", s.rootEndpoint)
	s.Router.HandleFunc("/start", s.startEndpoint)
	s.Router.HandleFunc("/close", s.closeEndpoint)
}

func (s *Server) Start() error {
	return s.HTTPServer.ListenAndServe()
}

func (s *Server) rootEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/health check endpoint")

	w.Write([]byte("OK"))
}

func (s *Server) startEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/start endpoint")

	respBody, err := s.handleStart(r)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	fmt.Println("/start endpoint resp: ", string(respBody))

	w.Write(respBody)
}

func (s *Server) handleStart(r *http.Request) ([]byte, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read body")
	}

	req := &api.StartRequest{}
	err = json.Unmarshal(body, req)
	if err != nil {
		return nil, errors.Wrap(err, "invalid body in request")
	}

	resp, err := s.Service.Start(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to start auction")
	}

	respBytes, err := json.Marshal(resp)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal response")
	}

	return respBytes, nil
}

func (s *Server) closeEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/close endpoint")

	respBody, err := s.handleClose(r)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	fmt.Println("/close endpoint resp: ", string(respBody))

	w.Write(respBody)
}

func (s *Server) handleClose(r *http.Request) ([]byte, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read body")
	}

	req := &api.CloseRequest{}
	err = json.Unmarshal(body, req)
	if err != nil {
		return nil, errors.Wrap(err, "invalid body in request")
	}

	resp, err := s.Service.Close(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to stop auction")
	}

	respBytes, err := json.Marshal(resp)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal response")
	}

	return respBytes, nil
}
