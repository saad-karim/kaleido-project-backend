package bidding_test

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
	contractID = "0xab3480457aae46a93e235fbfb01d180c49c46131"
)

func TestBiddingServiceGet(t *testing.T) {
	gt := NewGomegaWithT(t)

	bidder := bidding.Bidding{
		Config: config.APIGateway(),
		Client: &http.Client{},
	}

	resp, err := bidder.Get(contractID)
	gt.Expect(err).NotTo(HaveOccurred())

	respBody, err := ioutil.ReadAll(resp.Body)
	gt.Expect(err).NotTo(HaveOccurred())

	gt.Expect(resp.StatusCode).To(Equal(200))
	fmt.Printf("!! SK >>> respBody: %s\n", respBody)
}

func TestBiddingServiceBid(t *testing.T) {
	gt := NewGomegaWithT(t)

	bidder := bidding.Bidding{
		Config: config.APIGateway(),
		Client: &http.Client{},
	}

	resp, err := bidder.Bid(contractID, "99")
	gt.Expect(err).NotTo(HaveOccurred())

	respBody, err := ioutil.ReadAll(resp.Body)
	gt.Expect(err).NotTo(HaveOccurred())

	gt.Expect(resp.StatusCode).To(Equal(200))
	fmt.Printf("!! SK >>> respBody: %s\n", respBody)
}
