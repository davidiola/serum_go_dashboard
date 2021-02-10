package utils

import (
	"github.com/davidiola/serum_go_dashboard/models/api"
	"sort"
)

// min = true, retrieve min of orders
// min = false, retrieve max of orders
func RetrieveMinOrMaxFromOrders(orders []api.Order, min bool) float64 {
	if len(orders) == 0 {
		return 0.0
	}
	sort.Slice(orders, func(i, j int) bool {
		return orders[i].Price < orders[j].Price
	})
	if min {
		return orders[0].Price
	}
	return orders[len(orders)-1].Price
}

func RetrievePairToTradeMap(tradeData []api.TradeData, pairs []string) map[string]api.TradeData {
	m := make(map[string]api.TradeData)
	// sort trade data by timestamp
	sort.Slice(tradeData, func(i, j int) bool {
		return tradeData[i].Size < tradeData[j].Size
	})

	for _, pair := range pairs {
		for _, trade := range tradeData {
			if trade.Market == pair {
				m[pair] = trade
			}
		}
	}
	return m
}
