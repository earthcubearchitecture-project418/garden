package main

// This code needs docker image knqz/chrome-headless which you can get with
//  docker pull knqz/chrome-headless
// run with
//  docker run -d -p 9222:9222 --rm --name chrome-headless knqz/chrome-headless

import (
	"context"
	"log"

	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/client"
)

func main() {
	var err error

	// create context
	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create chrome
	c, err := chromedp.New(ctxt, chromedp.WithTargets(client.New().WatchPageTargets(ctxt)), chromedp.WithLog(log.Printf))
	if err != nil {
		log.Fatal(err)
	}

	// run task list
	var site, res string
	err = c.Run(ctxt, text("https://www2.earthref.org/MagIC/16403", &res))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Text %s --> %s", site, res)
}

func text(targeturl string, res *string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(targeturl),
		// chromedp.Text(`#react-root`, res, chromedp.NodeVisible, chromedp.ByID),
		chromedp.Text(`tagByTypeApplicationLDJSON`, res, chromedp.NodeVisible, chromedp.ByID),
	}
}
