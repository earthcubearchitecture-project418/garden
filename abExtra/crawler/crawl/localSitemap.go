package crawl

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"

	"earthcube.org/Project418/crawler/digester"
	"github.com/blevesearch/bleve"
)

func ProcessLocalSitemap(file, textindex string, spatial bool) error {

	// Check that it is XML and can be processed into the struct above
	siteArray := IngestLocalSitemapXML(file)
	count := 0
	total := len(siteArray.URL)

	index, berr := bleve.Open(textindex)
	if berr != nil {
		// should panic here..  no idex..  no reason
		log.Printf("Bleve error making index %v \n", berr)
	}

	for entry := range siteArray.URL {
		count = count + 1
		percent := (float64(count) / float64(total)) * 100.0
		fmt.Printf("%d / %d   %.2f%%  \n ", count, total, percent)
		err := digester.DigestPage(siteArray.URL[entry].Loc, index, spatial)
		if err != nil {
			log.Printf("PROCESSDOMAIN error")
		}
		log.Printf("CRAWLER report:  URL %s done with error code: %v\n", siteArray.URL[entry].Loc, err)
	}

	index.Close()

	return nil
}

func IngestLocalSitemapXML(file string) URLSet {

	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	// Process XML into a struct
	var sitemap URLSet
	xml.Unmarshal(content, &sitemap)

	return sitemap
}
