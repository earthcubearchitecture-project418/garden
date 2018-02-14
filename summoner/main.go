package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	// "earthcube.org/Project418/garden/summoner/sitemaps"
	"earthcube.org/Project418/garden/summoner/sitemaps"
	"earthcube.org/Project418/garden/summoner/utils"
	"github.com/PuerkitoBio/goquery"
	"github.com/kazarena/json-gold/ld"
)

func main() {

	// read the input file and collecte the URL for the sitemaps
	// func to read the sitemaps URLs into an array:   list
	// place the arrays into a map(string[]string) where the key is the domain name with dots removed
	// pass this array to the function  summon
	// summon can currently read the URLs and place
	//     1) the json-ld into S3
	//     2) the URL and status into a KV store , named with the key from the map
	// Once the summon is done the Miller can be called
	//  The miller function will parse off the tasks to be done (text, spatial, graph)
	// All these can now be done as quickly as a wait group will allow
	// since the fiels are read from the s3 holdings.

	sources := flag.String("sources", "", "A file with a list of source URLs for sitemaps, one per line")
	flag.Parse()

	domains, err := domainList(sources)
	if err != nil {
		log.Printf("Error reading list of domains %v\n", err)
	}

	ru := resourceURLs(domains) //  map by domain name and []string of landing page URLs
	actOnURL(ru)
}

func resourceURLs(domains []string) map[string]sitemaps.URLSet {

	// make a map
	m := make(map[string]sitemaps.URLSet)

	for key := range domains {
		mapname, _, err := utils.DomainNameShort(domains[key])
		if err != nil {
			log.Println("Error in domain parsing")
		}
		log.Println(mapname)
		us := sitemaps.IngestSitemapXML(domains[key])
		m[mapname] = us
	}

	return m
}

func actOnURL(m map[string]sitemaps.URLSet) {

	// a blocking channel to keep concurrency under control
	semaphoreChan := make(chan struct{}, 20)
	defer close(semaphoreChan)

	// a wait group enables the main process a wait for goroutines to finish
	wg := sync.WaitGroup{}

	for k := range m {
		log.Printf("Act on URL's for %s", k)
		for i := range m[k].URL {

			wg.Add(1)

			// log.Printf("----> %s", m[k].URL[i].Loc)
			url := m[k].URL[i].Loc

			go func(i int) {
				// block until the semaphore channel has room
				// this could also be moved out of the goroutine
				// which would make sense if the list is huge
				semaphoreChan <- struct{}{}

				var client http.Client

				req, err := http.NewRequest("GET", url, nil)
				if err != nil {
					log.Print(err) // not even being able to make a req instance..  might be a fatal thing?
				}

				req.Header.Set("User-Agent", "EarthCube_DataBot/1.0")

				resp, err := client.Do(req)
				if err != nil {
					log.Printf("Error reading sitemap: %s", err)
				}
				defer resp.Body.Close()

				doc, err := goquery.NewDocumentFromResponse(resp)
				if err != nil {
					fmt.Printf("error %v", err)
					// return "", err
				}

				// TODO Version that just looks for script type application/ld+json
				// this will look for ALL nodes in the doc that match, there may be more than one
				var jsonld string
				doc.Find("script").Each(func(i int, s *goquery.Selection) {
					val, _ := s.Attr("type")
					if val == "application/ld+json" {
						err = isValid(s.Text())
						if err != nil {
							log.Printf("ERROR: At %s JSON-LD is NOT valid: %s", url, err)
						}
						jsonld = s.Text()
					}
				})

				// send to minio..  be sure to look for something back to ensure
				// we stay in sync
				log.Println(jsonld)

				wg.Done() // tell the wait group that we be done

				log.Printf("#%d would get %s ", i, url) // print an message containing the index (won't keep order)
				<-semaphoreChan                         // clear a spot in the semaphore channel
			}(i)

		}
	}

	// wait for all the goroutines to be done
	wg.Wait()

}

func domainList(sources *string) ([]string, error) { // return map(string[]string)

	log.Printf("Opening source list file: %s \n", *sources)

	var domains []string

	file, err := os.Open(*sources)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		domains = append(domains, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return domains, nil

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
