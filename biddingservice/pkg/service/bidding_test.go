package service_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/saad-karim/kaleido-project-backend/biddingservice/pkg/bidding"
	"github.com/saad-karim/kaleido-project-backend/biddingservice/pkg/config"
)

const (
	contractID = "0x2a9a1cbaf2cf13a9c63fc7d9bd9c9d9f5920796b"
)

func TestBiddingServiceGet(t *testing.T) {
	gt := NewGomegaWithT(t)

	bidder := bidding.Bidding{
		Config: config.APIGateway(),
		Client: &http.Client{},
	}

	resp, err := bidder.CurrentBid(contractID)
	gt.Expect(err).NotTo(HaveOccurred())

	respBody, err := ioutil.ReadAll(resp.Body)
	gt.Expect(err).NotTo(HaveOccurred())

	fmt.Printf("!! SK >>> respBody: %s\n", respBody)
	gt.Expect(resp.StatusCode).To(Equal(200))
}

func TestBiddingServiceAllBids(t *testing.T) {
	gt := NewGomegaWithT(t)

	bidder := bidding.Bidding{
		Config: config.APIGateway(),
		Client: &http.Client{},
	}

	resp, err := bidder.BidHistory(contractID, 0)
	gt.Expect(err).NotTo(HaveOccurred())

	respBody, err := ioutil.ReadAll(resp.Body)
	gt.Expect(err).NotTo(HaveOccurred())

	fmt.Printf("!! SK >>> respBody: %s\n", respBody)
	gt.Expect(resp.StatusCode).To(Equal(200))
}

func TestBiddingServiceHighestBidder(t *testing.T) {
	gt := NewGomegaWithT(t)

	bidder := bidding.Bidding{
		Config: config.APIGateway(),
		Client: &http.Client{},
	}

	resp, err := bidder.HighestBidder(contractID)
	gt.Expect(err).NotTo(HaveOccurred())

	respBody, err := ioutil.ReadAll(resp.Body)
	gt.Expect(err).NotTo(HaveOccurred())

	fmt.Printf("!! SK >>> respBody: %s\n", respBody)
	gt.Expect(resp.StatusCode).To(Equal(200))
}

func TestBiddingServiceBid(t *testing.T) {
	gt := NewGomegaWithT(t)

	bidder := bidding.Bidding{
		Config: config.APIGateway(),
		Client: &http.Client{},
	}

	resp, err := bidder.PlaceBid("0x1718c3e8fbd6462849bbeb662c95e67f20259d7d", contractID, "110")
	gt.Expect(err).NotTo(HaveOccurred())

	respBody, err := ioutil.ReadAll(resp.Body)
	gt.Expect(err).NotTo(HaveOccurred())

	fmt.Printf("!! SK >>> respBody: %s\n", respBody)
	gt.Expect(resp.StatusCode).To(Equal(200))
}
