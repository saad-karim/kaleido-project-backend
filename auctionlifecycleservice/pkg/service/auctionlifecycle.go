package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/pkg/errors"
	"github.com/saad-karim/kaleido-project-backend/auctionlifecycleservice/pkg/api"
	"github.com/saad-karim/kaleido-project-backend/auctionlifecycleservice/pkg/config"
)

type Client interface {
	Post(url, contentType string, body io.Reader) (resp *http.Response, err error)
}

type AuctionDB interface {
	InsertAuction(*api.AuctionDBRow) error
	CloseAuction(string) error
}

// AuctionLifecycle manages the start and close of auctions
type AuctionLifecycle struct {
	Config *config.KaleidoAPIGateway
	Client Client
	DB     AuctionDB
}

// Start starts the auction
func (al *AuctionLifecycle) Start(req *api.StartRequest) (*api.StartAuctionResponse, error) {
	fmt.Printf("Service - Start: req - %+v\n", req)

	url := fmt.Sprintf("https://%s:%s@%s/gateways/%s/?kld-from=%s&kld-sync=true", al.Config.KaleidoAuthUsername, al.Config.KaleidoAuthPassword, al.Config.KaleidoRestGatewayURL, al.Config.Gateway, al.Config.FromAddress)

	fmt.Println("URL: ", url)

	postBody, err := json.Marshal(req)
	if err != nil {
		return nil, errors.Wrap(err, "invalid request format")
	}

	resp, err := al.Client.Post(url, "application/json", bytes.NewReader(postBody))
	if err != nil {
		return nil, errors.Wrap(err, "failed to start/initialize chaincode")
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read response body")
	}

	auctionResp := &api.StartAuctionResponse{}
	if err := json.Unmarshal(respBody, auctionResp); err != nil {
		return nil, err
	}

	price, err := strconv.Atoi(req.StartingBid)
	if err != nil {
		return nil, err
	}

	row := &api.AuctionDBRow{
		ID:    auctionResp.ContractAddress,
		Item:  req.AssetForSale,
		Price: price,
	}
	if err := al.DB.InsertAuction(row); err != nil {
		return nil, errors.Wrap(err, "failed to insert row in database")
	}

	return auctionResp, nil
}

// Close closes the auction
func (al *AuctionLifecycle) Close(req *api.CloseRequest) (*api.CloseAuctionResponse, error) {
	fmt.Printf("Service - Close: req - %+v\n", req)

	url := fmt.Sprintf("%s/%s/closeAuction?kld-from=%s&kld-sync=true", al.baseURL(), req.ContractAddress, al.Config.FromAddress)

	fmt.Println("URL: ", url)

	_, err := al.Client.Post(url, "application/json", nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to start/initialize chaincode")
	}

	if err := al.DB.CloseAuction(req.ContractAddress); err != nil {
		return nil, errors.Wrap(err, "failed to update row in database")
	}

	return &api.CloseAuctionResponse{}, nil
}

func (al *AuctionLifecycle) baseURL() string {
	return fmt.Sprintf("https://%s:%s@%s/gateways/%s", al.Config.KaleidoAuthUsername, al.Config.KaleidoAuthPassword, al.Config.KaleidoRestGatewayURL, al.Config.Gateway)
}
