package main

import (
	"flag"
	"fmt"
	"github.com/blevesearch/bleve"
	"log"
)

// i := "/home/fils/Project418/gleaner/output/20181119_115954/bleve/ssdbpro2betaiodporg"
func main() {
	// command line flags
	var i, q string
	flag.StringVar(&i, "index", "", "a bleve index")
	flag.StringVar(&q, "query", "", "a query string to look for in the index")
	flag.Parse()

	if len(i) < 1 || len(q) < 1 {
		fmt.Println("You must provide -index and -query values, place multiword querries in quotes")
	}

	// open an index
	index, err := bleve.Open(i)
	if err != nil {
		log.Println(err)
	}

	// search for some text
	query := bleve.NewMatchQuery(q)
	search := bleve.NewSearchRequest(query)
	searchResults, err := index.Search(search)
	if err != nil {
		log.Println(err)
	}

	log.Println(searchResults)
}
