package main

import (
	"fmt"
	"github.com/davidiola/serum_go_dashboard/client"
	"github.com/davidiola/serum_go_dashboard/constants"
	"github.com/davidiola/serum_go_dashboard/models/api"
	"github.com/davidiola/serum_go_dashboard/utils"
	"github.com/jroimartin/gocui"
	"log"
	"strconv"
	"strings"
)

func main() {

	c := client.New()
	pairs := c.RetrieveFirstNPairs(constants.NUM_PAIRS)
	fmt.Println("Retrieving Pairs...")
	volData := make([]api.VolumeData, constants.NUM_PAIRS)
	orderBookData := make([]api.OrderBookData, constants.NUM_PAIRS)

	for i, pair := range pairs {
		fmt.Printf("Retrieving Market Data for %s...\n", pair)
		volData[i] = c.RetrieveVolumeForMarket(pair)[0]
		orderBookData[i] = c.RetrieveOrderBookForMarket(pair)
	}

	fmt.Println("Retrieving last 24h of recent trades...")
	tradeData := c.RetrieveTradesPastDay()

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	g.Update(func(g *gocui.Gui) error {
		v, err := g.View("toprow")
		if err != nil {
			return err
		}
		fmt.Fprintln(v)
		fmt.Fprintln(v)
		for i, _ := range pairs {
			fmt.Fprintf(v, "%v", pairs[i])
			fmt.Fprintf(v, strings.Repeat(" ", constants.NUM_PADDING-len(pairs[i])))
		}
		return nil
	})

	g.Update(func(g *gocui.Gui) error {
		v, err := g.View("middle")
		if err != nil {
			return err
		}

		fmt.Fprintln(v)
		fmt.Fprintln(v)

		for i, _ := range pairs {
			fmt.Fprintf(v, "%.4f", volData[i].VolumeUsd)
			fmt.Fprintf(v, strings.Repeat(" ", constants.NUM_PADDING-len(strconv.FormatFloat(volData[i].VolumeUsd, 'f', constants.NUM_DECIMALS, 64))))
		}

		fmt.Fprintln(v)

		for i, _ := range pairs {
			maxBid := utils.RetrieveMinOrMaxFromOrders(orderBookData[i].Bids, false)
			fmt.Fprintf(v, "%.4f", maxBid)
			fmt.Fprintf(v, strings.Repeat(" ", constants.NUM_PADDING-len(strconv.FormatFloat(maxBid, 'f', constants.NUM_DECIMALS, 64))))
		}

		fmt.Fprintln(v)

		for i, _ := range pairs {
			minAsk := utils.RetrieveMinOrMaxFromOrders(orderBookData[i].Asks, true)
			fmt.Fprintf(v, "%.4f", minAsk)
			fmt.Fprintf(v, strings.Repeat(" ", constants.NUM_PADDING-len(strconv.FormatFloat(minAsk, 'f', constants.NUM_DECIMALS, 64))))
		}

		fmt.Fprintln(v)

		spread := make([]float64, constants.NUM_PAIRS)

		for i, _ := range pairs {
			maxBid := utils.RetrieveMinOrMaxFromOrders(orderBookData[i].Bids, false)
			minAsk := utils.RetrieveMinOrMaxFromOrders(orderBookData[i].Asks, true)
			spread[i] = minAsk - maxBid
			fmt.Fprintf(v, "%.4f", spread[i])
			fmt.Fprintf(v, strings.Repeat(" ", constants.NUM_PADDING-len(strconv.FormatFloat(spread[i], 'f', constants.NUM_DECIMALS, 64))))
		}

		fmt.Fprintln(v)

		for i, _ := range pairs {
			minAsk := utils.RetrieveMinOrMaxFromOrders(orderBookData[i].Asks, true)
			var sprPercentage float64
			if minAsk != 0 {
				sprPercentage = (spread[i] / minAsk) * 100
			} else {
				sprPercentage = 100.0
			}
			fmt.Fprintf(v, "%.4f", sprPercentage)
			fmt.Fprintf(v, strings.Repeat(" ", constants.NUM_PADDING-len(strconv.FormatFloat(sprPercentage, 'f', constants.NUM_DECIMALS, 64))))
		}

		fmt.Fprintln(v)

		pairToTradeMap := utils.RetrievePairToTradeMap(tradeData, pairs)
		for _, pair := range pairs {
			if trade, ok := pairToTradeMap[pair]; ok {
				fmt.Fprintf(v, "%.4f", trade.Price)
				fmt.Fprintf(v, strings.Repeat(" ", constants.NUM_PADDING-len(strconv.FormatFloat(trade.Price, 'f', constants.NUM_DECIMALS, 64))))
			} else {
				fmt.Fprintf(v, "%.4f", 0.0000)
				fmt.Fprintf(v, strings.Repeat(" ", constants.NUM_PADDING-len(strconv.FormatFloat(0.0000, 'f', constants.NUM_DECIMALS, 64))))
			}
		}

		fmt.Fprintln(v)

		for _, pair := range pairs {
			if trade, ok := pairToTradeMap[pair]; ok {
				fmt.Fprintf(v, "%.4f", trade.Size)
				fmt.Fprintf(v, strings.Repeat(" ", constants.NUM_PADDING-len(strconv.FormatFloat(trade.Size, 'f', constants.NUM_DECIMALS, 64))))
			} else {
				fmt.Fprintf(v, "%.4f", 0.0000)
				fmt.Fprintf(v, strings.Repeat(" ", constants.NUM_PADDING-len(strconv.FormatFloat(0.0000, 'f', constants.NUM_DECIMALS, 64))))
			}
		}

		return nil
	})

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}

}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("leftpanel", -1, 2, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v)
		fmt.Fprintln(v)
		fmt.Fprintln(v, " Volume (USD)")
		fmt.Fprintln(v, " Max Bid")
		fmt.Fprintln(v, " Min Ask")
		fmt.Fprintln(v, " B/A Spread")
		fmt.Fprintln(v, " Spread %")
		fmt.Fprintln(v, " Last Fill (24h)")
		fmt.Fprintln(v, " Last Fill Size")
	}
	if _, err := g.SetView("toprow", int(0.09*float32(maxX)), -1, maxX, maxY); err != nil && err != gocui.ErrUnknownView {
		return err
	}
	if _, err := g.SetView("middle", int(0.09*float32(maxX)), 2, maxX, maxY); err != nil && err != gocui.ErrUnknownView {
		return err
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
