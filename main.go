package main

import (
	"fmt"
	"github.com/davidiola/serum_go_dashboard/client"
	"github.com/jroimartin/gocui"
	"log"
	"strconv"
	"strings"
)

func main() {

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	c := client.New()
	pairs := c.RetrieveFirstNPairs(15)

	g.Update(func(g *gocui.Gui) error {
		v, err := g.View("toprow")
		if err != nil {
			return err
		}
		fmt.Fprintln(v)
		fmt.Fprintln(v)
		for _, pair := range pairs {
			fmt.Fprintf(v, "%v", pair)
			fmt.Fprintf(v, strings.Repeat(" ", 15-len(pair)))
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
		for _, pair := range pairs {
			volData := c.RetrieveVolumeForMarket(pair)[0]
			fmt.Fprintf(v, "%.4f", volData.VolumeUsd)
			fmt.Fprintf(v, strings.Repeat(" ", 15-len(strconv.FormatFloat(volData.VolumeUsd, 'f', 4, 64))))
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
		fmt.Fprintf(v, "    ")
		fmt.Fprintln(v, "Volume")
	}
	if _, err := g.SetView("toprow", int(0.05*float32(maxX)), -1, maxX, maxY); err != nil && err != gocui.ErrUnknownView {
		return err
	}
	if _, err := g.SetView("middle", int(0.05*float32(maxX)), 2, maxX, maxY); err != nil && err != gocui.ErrUnknownView {
		return err
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
