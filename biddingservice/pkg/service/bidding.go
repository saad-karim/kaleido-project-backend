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
	"github.com/saad-karim/kaleido-project-backend/biddingservice/pkg/api"
	"github.com/saad-karim/kaleido-project-backend/biddingservice/pkg/config"
)

// Contract Address: 0x2a9a1cbaf2cf13a9c63fc7d9bd9c9d9f5920796b

type Client interface {
	Get(url string) (resp *http.Response, err error)
	Post(url, contentType string, body io.Reader) (resp *http.Response, err error)
}

type AuctionDB interface {
	GetOpenAuctions() ([]api.AuctionDBRow, error)
}

// Bidding bids on auctions
type Bidding struct {
	Config *config.KaleidoAPIGateway
	Client Client
	DB     AuctionDB
}

func (b *Bidding) GetOpenAuctions() ([]byte, error) {
	results, err := b.DB.GetOpenAuctions()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get completed database query for all open auctions")
	}

	respBytes, err := json.Marshal(results)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate response")
	}

	return respBytes, nil
}

// Get gets the current bid on an auction
func (b *Bidding) CurrentBid(contractAddress string) ([]byte, error) {
	// url := fmt.Sprintf("https://%s:%s@%s/gateways/%s/%s/bid?kld-from=%s&kld-sync=true", b.Config.KaleidoAuthUsername, b.Config.KaleidoAuthPassword, b.Config.KaleidoRestGatewayURL, b.Config.Gateway, contractAddress, b.Config.FromAddress)
	url := fmt.Sprintf("%s/bid?kld-from=%s&kld-sync=true", b.baseURL(contractAddress), b.Config.FromAddress)

	fmt.Println("URL: ", url)

	resp, err := b.Client.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get current highest bid")
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}

type Response struct {
	Output string `json:"output"`
}

func (b *Bidding) BidHistory(contractAddress string) ([]byte, error) {
	// url := fmt.Sprintf("https://%s:%s@%s/gateways/%s/%s/bids?input=%d&kld-from=%s&kld-sync=true", b.Config.KaleidoAuthUsername, b.Config.KaleidoAuthPassword, b.Config.KaleidoRestGatewayURL, b.Config.Gateway, contractAddress, index, b.Config.FromAddress)
	url := fmt.Sprintf("%s/bidCount?kld-from=%s&kld-sync=true", b.baseURL(contractAddress), b.Config.FromAddress)

	fmt.Println("URL: ", url)

	resp, err := b.Client.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get bid count")
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read bid count response")
	}

	fmt.Println("bid count resp: ", string(respBody))

	bidCount := &Response{}
	err = json.Unmarshal(respBody, bidCount)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal bid count")
	}

	bids, err := strconv.Atoi(bidCount.Output)
	if err != nil {
		return nil, err
	}

	bidHistory := []api.BidHistory{}
	for x := 0; x < bids; x++ {
		url := fmt.Sprintf("%s/bids?input=%d&kld-from=%s&kld-sync=true", b.baseURL(contractAddress), x, b.Config.FromAddress)
		fmt.Println("URL: ", url)

		resp, err := b.Client.Get(url)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get bid history")
		}

		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, errors.Wrap(err, "failed to read bid history response")
		}

		fmt.Println("bid history resp: ", string(respBody))

		bid := api.BidHistory{}
		err = json.Unmarshal(respBody, &bid)
		if err != nil {
			return nil, errors.Wrap(err, "failed to unmarshal bid history")
		}

		bidHistory = append(bidHistory, bid)
	}

	respBytes, err := json.Marshal(bidHistory)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal bid history response")
	}

	return respBytes, nil
}

func (b *Bidding) HighestBidder(contractAddress string) ([]byte, error) {
	// url := fmt.Sprintf("https://%s:%s@%s/gateways/%s/%s/highestBidder?kld-from=%s&kld-sync=true", b.Config.KaleidoAuthUsername, b.Config.KaleidoAuthPassword, b.Config.KaleidoRestGatewayURL, b.Config.Gateway, contractAddress, b.Config.FromAddress)
	url := fmt.Sprintf("%s/highestBidder?kld-from=%s&kld-sync=true", b.baseURL(contractAddress), b.Config.FromAddress)

	fmt.Println("URL: ", url)

	resp, err := b.Client.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get current highest bid")
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}

func (b *Bidding) PlaceBid(bidderAddress, contractAddress, bidValue string) ([]byte, error) {
	// url := fmt.Sprintf("https://%s:%s@%s/gateways/%s/%s/placeBid?kld-from=%s&kld-sync=true", b.Config.KaleidoAuthUsername, b.Config.KaleidoAuthPassword, b.Config.KaleidoRestGatewayURL, b.Config.Gateway, contractAddress, b.Config.FromAddress)
	url := fmt.Sprintf("%s/placeBid?kld-from=%s&kld-sync=true", b.baseURL(contractAddress), bidderAddress)

	fmt.Println("URL: ", url)

	bid := &api.Bid{
		Amount: bidValue,
	}

	postBody, err := json.Marshal(bid)
	if err != nil {
		return nil, errors.Wrap(err, "invalid request format")
	}

	resp, err := b.Client.Post(url, "application/json", bytes.NewReader(postBody))
	if err != nil {
		return nil, errors.Wrap(err, "failed to start/initialize chaincode")
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}

func (b *Bidding) baseURL(contractAddress string) string {
	return fmt.Sprintf("https://%s:%s@%s/gateways/%s/%s", b.Config.KaleidoAuthUsername, b.Config.KaleidoAuthPassword, b.Config.KaleidoRestGatewayURL, b.Config.Gateway, contractAddress)
}
