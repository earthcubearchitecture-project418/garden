package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/blevesearch/bleve"
)

const jsonld string = `{
	"@context": {
	 "@vocab": "http://schema.org/",
	 "re3data": "http://example.org/re3data/0.1/"
	},
	"@id": "http://opencoredata.org/id/dataset/bcd15975-680c-47db-a062-ac0bb6e66816",
	"@type": "Dataset",
	"description": "Data set description",
	"distribution": {
	 "@type": "DataDownload",
	 "contentUrl": "http://opencoredata.org/id/dataset/bcd15975-680c-47db-a062-ac0bb6e66816"
	},
	"keywords": "DSDP, OPD, IODP, JanusThermalConductivity",
	"name": "208_1262A_JanusThermalConductivity_VyaMsepM.csv",
	"publisher": {
	 "@type": "Organization",
	 "description": "NSF funded International Ocean Discovery Program operated by JRSO",
	 "name": "International Ocean Discovery Program",
	 "url": "http://iodp.org"
	},
	"spatial": {
	 "@type": "Place",
	 "geo": {
	  "@type": "GeoCoordinates",
	  "latitude": "-27.19",
	  "longitude": "1.58"
	 }
	},
	"url": "http://opencoredata.org/id/dataset/bcd15975-680c-47db-a062-ac0bb6e66816",
	"variableMeasured": "JanusThermalConductivity"
   } `

func main() {
	fmt.Println("Index and Store example")

	// Take the JSON-LD and pass to a function to store
	fmt.Println("index")
	err := indexDoc()
	if err != nil {
		log.Print(err)
	}

	// Search that now
	fmt.Println("now search")
	results := searchIndex()

	fmt.Println(results)
}

func indexDoc() error {
	// load jsonld into a bleve index

	mapping := bleve.NewIndexMapping()
	index, err := bleve.New("test.bleve", mapping)
	if err != nil {
		log.Print(err)
	}

	// index the text
	err = index.Index("123", jsonld)
	if err != nil {
		log.Print(err)
	}

	// load the document in too
	err = index.SetInternal([]byte(strconv.Itoa(123)), []byte(jsonld))
	if err != nil {
		log.Fatal("Trouble doing SetInternal!")
	}

	index.Close()

	fmt.Println("Indexing done")
	return err
}

func searchIndex() string {
	index, err := bleve.Open("test.bleve")
	if err != nil {
		log.Fatal()
	}

	query := bleve.NewMatchQuery("IODP")
	search := bleve.NewSearchRequest(query)
	searchResults, err := index.Search(search)
	if err != nil {
		log.Fatalln("Trouble with search request!")
	}
	fmt.Println("Search results")
	fmt.Println(searchResults)

	// try and get the internal document now....
	raw, err := index.GetInternal([]byte(strconv.Itoa(123)))
	if err != nil {
		log.Fatal("Trouble getting internal doc:", err)
	}

	fmt.Println("try and print the original document")
	fmt.Print(string(raw))

	return "done"
}
