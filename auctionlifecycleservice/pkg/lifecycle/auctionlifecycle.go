package lifecycle

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"
	"github.com/saad-karim/kaleido-project-backend/auctionlifecycleservice/pkg/config"
)

type Client interface {
	Post(url, contentType string, body io.Reader) (resp *http.Response, err error)
}

// AuctionLifecycle manages the start and close of auctions
type AuctionLifecycle struct {
	Config *config.KaleidoAPIGateway
	Client Client
}

type Request struct {
	InitValue string `json:"initVal"`
}

// Start starts the auction
func (al *AuctionLifecycle) Start() (*http.Response, error) {
	url := fmt.Sprintf("https://%s:%s@%s/gateways/u0emv3jn8q/?kld-from=%s&kld-sync=true", al.Config.KaleidoAuthUsername, al.Config.KaleidoAuthPassword, al.Config.KaleidoRestGatewayURL, al.Config.FromAddress)

	fmt.Println("URL: ", url)

	body := &Request{
		InitValue: "10",
	}

	postBody, err := json.Marshal(body)
	if err != nil {
		return nil, errors.Wrap(err, "invalid request format")
	}

	resp, err := al.Client.Post(url, "application/json", bytes.NewReader(postBody))
	if err != nil {
		return nil, errors.Wrap(err, "failed to start/initialize chaincode")
	}

	return resp, nil
}

// Close closes the auction
func (al *AuctionLifecycle) Close() error {
	return nil
}
