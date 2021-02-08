package client

import (
	"encoding/json"
	"github.com/davidiola/serum_go_dashboard/constants"
	"github.com/davidiola/serum_go_dashboard/models/api"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Client struct {
	BaseEndpoint string
}

func New() Client {
	c := Client{constants.BONFIDA_BASE_ENDPOINT}
	return c
}

func (c Client) RetrieveFirstNPairs(n int) []string {
	var pairs api.Pairs
	body := c.GatherHTTPResponseBody(constants.PAIRS)
	json.Unmarshal([]byte(body), &pairs)
	if pairs.Success {
		return pairs.Data[:n]
	} else {
		return []string{}
	}
}

func (c Client) RetrieveVolumeForMarket(marketIdentifier string) []api.VolumeData {
	var volume api.Volume
	body := c.GatherHTTPResponseBody(constants.VOLUME + strings.ReplaceAll(marketIdentifier, "/", ""))
	json.Unmarshal([]byte(body), &volume)
	if volume.Success {
		return volume.Data
	} else {
		return []api.VolumeData{api.VolumeData{}}
	}
}

func (c Client) RetrieveOrderBookForMarket(marketIdentifier string) api.OrderBookData {
	var orderbook api.OrderBook
	body := c.GatherHTTPResponseBody(constants.ORDERBOOK + strings.ReplaceAll(marketIdentifier, "/", ""))
	json.Unmarshal([]byte(body), &orderbook)
	if orderbook.Success {
		return orderbook.Data
	} else {
		return api.OrderBookData{}
	}
}

func (c Client) GatherHTTPResponseBody(endpoint string) []byte {
	resp, err := http.Get(c.BaseEndpoint + endpoint)
	if err != nil {
		c.HandleHTTPFailure(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

func (c Client) HandleHTTPFailure(err error) {
	log.Panicln("HTTP Failure.")
}
