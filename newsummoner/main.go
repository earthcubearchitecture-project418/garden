package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/exporter"
)

func main() {
	geziyor.NewGeziyor(geziyor.Options{
		//		StartURLs: []string{"http://quotes.toscrape.com/"},
		StartURLs: []string{"http://opencoredata.org/doc/dataset/bcd15975-680c-47db-a062-ac0bb6e66816"},
		ParseFunc: quotesParse,
		Exporters: []geziyor.Exporter{exporter.JSONExporter{}},
	}).Start()

}

func quotesParse(r *geziyor.Response) {
	r.DocHTML.Find("div.quote").Each(func(i int, s *goquery.Selection) {
		r.Exports <- map[string]interface{}{
			"text":   s.Find("span.text").Text(),
			"author": s.Find("small.author").Text(),
		}

	})
	if href, ok := r.DocHTML.Find("li.next > a").Attr("href"); ok {
		go r.Geziyor.Get(r.JoinURL(href), quotesParse)

	}

}
