package bidding

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"
	"github.com/saad-karim/kaleido-project-backend/biddingservice/pkg/config"
)

// Contract Address: 0xab3480457aae46a93e235fbfb01d180c49c46131

type Client interface {
	Get(url string) (resp *http.Response, err error)
	Post(url, contentType string, body io.Reader) (resp *http.Response, err error)
}

// Bidding bids on auctions
type Bidding struct {
	Config *config.KaleidoAPIGateway
	Client Client
}

type Set struct {
	X string `json:"x"`
}

// Get gets the current bid on an auction
func (b *Bidding) Get(contractAddress string) (*http.Response, error) {
	url := fmt.Sprintf("https://%s:%s@%s/gateways/u0emv3jn8q/%s/get?kld-from=%s&kld-sync=true", b.Config.KaleidoAuthUsername, b.Config.KaleidoAuthPassword, b.Config.KaleidoRestGatewayURL, contractAddress, b.Config.FromAddress)

	fmt.Println("URL: ", url)

	resp, err := b.Client.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "failed to start/initialize chaincode")
	}

	return resp, nil
}

// Bid bids on a auction
func (b *Bidding) Bid(contractAddress string, bidValue string) (*http.Response, error) {
	url := fmt.Sprintf("https://%s:%s@%s/gateways/u0emv3jn8q/%s/set?kld-from=%s&kld-sync=true", b.Config.KaleidoAuthUsername, b.Config.KaleidoAuthPassword, b.Config.KaleidoRestGatewayURL, contractAddress, b.Config.FromAddress)

	fmt.Println("URL: ", url)

	body := &Set{
		X: bidValue,
	}

	postBody, err := json.Marshal(body)
	if err != nil {
		return nil, errors.Wrap(err, "invalid request format")
	}

	resp, err := b.Client.Post(url, "application/json", bytes.NewReader(postBody))
	if err != nil {
		return nil, errors.Wrap(err, "failed to start/initialize chaincode")
	}

	return resp, nil
}
