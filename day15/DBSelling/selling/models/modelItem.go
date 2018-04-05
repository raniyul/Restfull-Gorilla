package models

type Selling struct {
	tblSellingID  int     `json:"tblSellingID"`
	SellingCode   string  `json:"SellingCode"`
	SellingName   string  `json:"SellingName"`
	BuyingPrice   float64 `json:"BuyingPrice"`
	SellingPrice  float64 `json:"SellingPrice"`
	SellingAmount int     `json:"SellingAmount"`
	Pieces        string  `json:"Pieces"`
}
