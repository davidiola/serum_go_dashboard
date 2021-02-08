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
