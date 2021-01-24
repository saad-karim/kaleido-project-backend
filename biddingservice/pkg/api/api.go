package api

type Bid struct {
	User   string `json:"user"`
	Amount string `json:"amount"`
}

type AuctionDBRow struct {
	ID     string `db:"id"`
	Item   string `db:"item"`
	Price  int    `db:"price"`
	Closed bool   `db:"closed"`
}

type BidHistory struct {
	Amount string `json:"amount"`
	Bidder string `json:"bidder"`
}
