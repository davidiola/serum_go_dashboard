package api

type Pairs struct {
	Success bool     `json:"success"`
	Data    []string `json:"data"`
}

type VolumeData struct {
	VolumeUsd float64 `json:"volumeUsd"`
	Volume    float64 `json:"volume"`
}

type Volume struct {
	Success bool         `json:"success"`
	Data    []VolumeData `json:"data"`
}

type Order struct {
	Price float64 `json:"price"`
	Size  float64 `json:"size"`
}

type OrderBookData struct {
	Market        string  `json:"market"`
	Bids          []Order `json:"bids"`
	Asks          []Order `json:"asks"`
	MarketAddress string  `json:"marketAddress"`
}

type OrderBook struct {
	Success bool
	Data    OrderBookData
}
