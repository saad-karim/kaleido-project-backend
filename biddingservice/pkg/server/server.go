package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	gmux "github.com/gorilla/mux"
	"github.com/saad-karim/kaleido-project-backend/biddingservice/pkg/api"
)

type Service interface {
	GetOpenAuctions() ([]byte, error)
	CurrentBid(contractAddress string) (*http.Response, error)
	BidHistory(contractAddress string) ([]byte, error)
	HighestBidder(contractAddress string) (*http.Response, error)
	PlaceBid(bidderAddress, contractAddress, bidValue string) (*http.Response, error)
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
	s.Router.HandleFunc("/auctions", s.getOpenAunctions).Methods("GET")
	s.Router.HandleFunc("/currentbid/{contractAddress}", s.currentBidEndpoint).Methods("GET")
	s.Router.HandleFunc("/bidhistory/{contractAddress}", s.bidHistoryEndpoint).Methods("GET")
	s.Router.HandleFunc("/highestbidder/{contractAddress}", s.highestBidderEndpoint).Methods("GET")
	s.Router.HandleFunc("/placebid/{contractAddress}", s.placeBidEndpoint).Methods("POST")
}

func (s *Server) Start() error {
	return s.HTTPServer.ListenAndServe()
}

func (s *Server) rootEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/health check endpoint")

	w.Write([]byte("OK"))
}

func (s *Server) getOpenAunctions(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/get open auctions endpoint")

	resp, err := s.Service.GetOpenAuctions()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	}

	w.Write(resp)
}

func (s *Server) currentBidEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/currentbuild endpoint")

	contractAddress := getParam(r, "contractAddress")
	resp, err := s.Service.CurrentBid(contractAddress)
	if err != nil {
		fmt.Println("!!SK >>> ERROR: ", err)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("!!SK >>> ERROR: ", err)
	}

	w.Write(respBody)
}

func (s *Server) bidHistoryEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/bidhistory endpoint")

	// fromAddress := "0xaa3347224b6ca9098db1dcdbc799a2f876d8fdc5"
	contractAddress := getParam(r, "contractAddress")
	resp, err := s.Service.BidHistory(contractAddress)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	}

	fmt.Println("/bidhistory endpoint resp: ", string(resp))

	w.Write(resp)
}

func (s *Server) highestBidderEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/highestbidder endpoint")

	// fromAddress := "0xaa3347224b6ca9098db1dcdbc799a2f876d8fdc5"
	contractAddress := getParam(r, "contractAddress")
	resp, err := s.Service.HighestBidder(contractAddress)
	if err != nil {
		fmt.Println("!!SK >>> ERROR: ", err)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("!!SK >>> ERROR: ", err)
	}

	w.Write(respBody)
}

func (s *Server) placeBidEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/placebid endpoint")

	contractAddress := getParam(r, "contractAddress")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("!!SK >>> ERROR: ", err)
	}

	req := &api.Bid{}
	err = json.Unmarshal(body, req)
	if err != nil {
		fmt.Println("!!SK >>> ERROR: ", err)
	}

	// fromAddress := "0xaa3347224b6ca9098db1dcdbc799a2f876d8fdc5"
	resp, err := s.Service.PlaceBid(req.User, contractAddress, req.Amount)
	if err != nil {
		fmt.Println("!!SK >>> ERROR: ", err)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("!!SK >>> ERROR: ", err)
	}

	w.Write(respBody)
}

func getParam(r *http.Request, param string) string {
	return gmux.Vars(r)[param]
}
