package lifecycle_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/saad-karim/kaleido-project-backend/auctionlifecycleservice/pkg/config"
	"github.com/saad-karim/kaleido-project-backend/auctionlifecycleservice/pkg/lifecycle"
)

func TestAuctionLifecycleStart(t *testing.T) {
	gt := NewGomegaWithT(t)

	al := lifecycle.AuctionLifecycle{
		Config: config.APIGateway(),
		Client: &http.Client{},
	}

	resp, err := al.Start()
	gt.Expect(err).NotTo(HaveOccurred())

	respBody, err := ioutil.ReadAll(resp.Body)
	gt.Expect(err).NotTo(HaveOccurred())

	gt.Expect(resp.StatusCode).To(Equal(200))
	fmt.Printf("!! SK >>> respBody: %s\n", respBody)
}
