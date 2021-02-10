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
	Success bool          `json:"success"`
	Data    OrderBookData `json:"data"`
}

type Trades struct {
	Success bool        `json:"success"`
	Data    []TradeData `json:"data"`
}

type TradeData struct {
	Market        string  `json:"market"`
	Price         float64 `json:"price"`
	Size          float64 `json:"size"`
	Side          string  `json:"side"`
	Time          string  `json:"time"`
	OrderId       string  `json:"orderId"`
	FeeCost       float64 `json:"feeCost"`
	MarketAddress string  `json:"marketAddress"`
}
