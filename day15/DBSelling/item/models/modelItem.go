package models

type Item struct {
	tblItemID    int     `json:"tblItemID"`
	ItemCode     string  `json:"ItemCode"`
	ItemName     string  `json:"ItemName"`
	BuyingPrice  float64 `json:"BuyingPrice"`
	SellingPrice float64 `json:"SellingPrice"`
	ItemAmount   int     `json:"ItemAmount"`
	Pieces       string  `json:"Pieces"`
}
