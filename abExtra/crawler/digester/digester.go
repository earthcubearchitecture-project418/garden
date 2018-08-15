package digester

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"earthcube.org/Project418/crawler/indexer"
	"github.com/PuerkitoBio/goquery"
	"github.com/blevesearch/bleve"
	"github.com/kazarena/json-gold/ld"
)

// DigestPage takes a URL and digests its content
func DigestPage(url string, index bleve.Index, spatial bool) error {
	// Let's toss in some classic GoLang concurrency here....   so we index concurrently (not parallel)
	// Fire these three off at the same time and wait...
	// Above this we can then fire off N sites too at the same time (TODO this)

	// TODO
	// 1) go get the URL body
	// 2) pass the body (or a framed version) to the indexer along with the URL for the ID
	// 3) ALT ID:  UUID like DOI from a framing event

	jsonld, err := getJSONLD(url)
	if err != nil {
		log.Printf("URL %s has error: %s", url, err)
		return nil // we have nothing to do with invalid JSON-LD
	}

	if jsonld == "" {
		log.Printf("URL %s is nil", url)
		return nil // we have nothing to do.. with nothing
	}

	messages := make(chan string)
	var wg sync.WaitGroup

	// TODO
	// Open the index...  and pass it to the
	// bleve indexer

	wg.Add(3) // or however many tasks we have...
	go func() {
		defer wg.Done()
		if spatial {
			result := indexer.SpatialIndexer(url, jsonld) //
			//result := indexer.SpatialParse(url, jsonld) // tsting new manual JSON parse rather than the JSON-LD framing API
			messages <- fmt.Sprintf("Spatial indexer status for URL %s: %s \n", url, result)
		} else {
			messages <- fmt.Sprintf("Spatial indexer not run for URL %s: %s \n", url, "")
		}
	}()

	go func() {
		defer wg.Done()
		result := indexer.TextIndexer(url, jsonld, index) // this will want an interface (struct)
		messages <- fmt.Sprintf("Text indexer status for URL %s: %s \n", url, result)
	}()

	go func() {
		defer wg.Done()
		result := indexer.GraphIndexer(jsonld) // this will have JSON-LD (bnode resolved) or triples
		messages <- fmt.Sprintf("Graph status done for URL %s: %s \n", url, result)
	}()

	// Place holder for the time series indexer
	// go func() {
	// 	defer wg.Done()
	// 	result := indexer.TimeIndexer(jsonld) // this will have JSON-LD (bnode resolved) or triples
	// 	messages <- fmt.Sprintf("Time series indexer status done for URL %s: %s \n", url, result)
	// }()

	go func() {
		for i := range messages {
			log.Printf("DIGESTER report %s", i)
		}
	}()

	wg.Wait()

	return nil
}

// have this return the JSON-LD and the nquades (for the graph index)
func getJSONLD(url string) (string, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Print(err) // One error is an IO timeout..  can just move on from there...
		return "", err
	}

	req.Header.Set("User-Agent", "EarthCube_DataBot/1.0")

	resp, err := client.Do(req)
	if err != nil {
		log.Print(err) // TODO..  make better.. but recall these errors should NOT be fatal
		return "", err
	}
	if resp.StatusCode != 200 {
		b, _ := ioutil.ReadAll(resp.Body)
		log.Print(err)
		log.Print(string(b))
	}

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		fmt.Printf("error %v", err)
		return "", err
	}

	// Version that just looks for script type application/ld+json
	// this will look for ALL nodes in the doc that match, there may be more than one
	var jsonld string
	doc.Find("script").Each(func(i int, s *goquery.Selection) {
		val, _ := s.Attr("type")
		if val == "application/ld+json" {
			// 		// fmt.Printf("%s\n", s.Text()) //  or send off to a scheme.org parser (JSONLD parser)
			err = isValid(s.Text()) // TODO..  I do nothing with this. :)   bad me...
			if err != nil {
				log.Printf("ERROR: At %s JSON-LD is NOT valid: %s", url, err)
			}
			jsonld = s.Text()
		}
	})

	// doc.Find("script").Each(func(i int, s *goquery.Selection) {
	// 	// s.Has()
	// 	val, _ := s.Attr("type")
	// 	if val == "application/ld+json" {
	// 		// fmt.Printf("%s\n", s.Text()) //  or send off to a scheme.org parser (JSONLD parser)
	// 		err = isValid(s.Text())
	// 	}
	// })

	return jsonld, err
}

func isValid(jsonld string) error {
	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")
	options.Format = "application/nquads"

	var myInterface interface{}
	err := json.Unmarshal([]byte(jsonld), &myInterface)
	if err != nil {
		log.Println("Error when transforming JSON-LD document to interface:", err)
		return err
	}

	_, err = proc.ToRDF(myInterface, options) // returns triples but toss them, we just want to see if this processes with no err
	if err != nil {
		log.Println("Error when transforming JSON-LD document to RDF:", err)
		return err
	}

	return err
}

// mockDataEvent  returns example JSONLD for local testing if needed
// func mockDataEvent() string {

// 	data := `{
// 		"@context": {
// 		 "@vocab": "http://schema.org/",
// 		 "re3data": "http://example.org/re3data/0.1/"
// 		},
// 		"@id": "http://opencoredata.org/id/dataset/bcd15975-680c-47db-a062-ac0bb6e66816",
// 		"@type": "Dataset",
// 		"description": "Data set description",
// 		"distribution": {
// 		 "@type": "DataDownload",
// 		 "contentUrl": "http://opencoredata.org/id/dataset/bcd15975-680c-47db-a062-ac0bb6e66816"
// 		},
// 		"keywords": "DSDP, OPD, IODP, JanusThermalConductivity",
// 		"name": "208_1262A_JanusThermalConductivity_VyaMsepM.csv",
// 		"publisher": {
// 		 "@type": "Organization",
// 		 "description": "NSF funded International Ocean Discovery Program operated by JRSO",
// 		 "name": "International Ocean Discovery Program",
// 		 "url": "http://iodp.org"
// 		},
// 		"spatial": {
// 		 "@type": "Place",
// 		 "geo": {
// 		  "@type": "GeoCoordinates",
// 		  "latitude": "-27.19",
// 		  "longitude": "1.58"
// 		 }
// 		},
// 		"url": "http://opencoredata.org/id/dataset/bcd15975-680c-47db-a062-ac0bb6e66816",
// 		"variableMeasured": "JanusThermalConductivity"
// 	   } `

// 	return data

// }
