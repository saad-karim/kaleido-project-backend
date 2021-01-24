package api

type StartRequest struct {
	AssetForSale string `json:"assetForSale"`
	StartingBid  string `json:"startingBid"`
}

type CloseRequest struct {
	FromAddress     string `json:"fromAddress"`
	ContractAddress string `json:"contractAddress"`
}

type StartAuctionResponse struct {
	ContractAddress string `json:"contractAddress"`
}

type CloseAuctionResponse struct{}

type AuctionDBRow struct {
	ID     string `db:"id"`
	Item   string `db:"item"`
	Price  int    `db:"price"`
	Closed bool   `db:"closed"`
}
