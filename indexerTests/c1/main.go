package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/anaskhan96/soup"
	"github.com/boltdb/bolt"
	"github.com/deiu/rdf2go"
	"github.com/kazarena/json-gold/ld"
)

// DataSetStruct is struct to hold information about schema.org/DataSet
type DataSetStruct struct {
	Description string
	ID          string
	Type        string
	URL         string
	// URL         string `json:"schema:url"`
}

// simple JSON-LD doc for some early testing   Will be removed later as we use simpleServer
const bodyTestOLD = `{
    "@context": "http://schema.org",
    "@type": "DataCatalog",
    "@id": "http://opencoredata.org/catalogs",
    "url": "http://opencoredata.org/catalogs",
    "description": "Can I use this approach to reference this catalog from type WebSite",
    "dataset": [{
            "@type": "Dataset",
            "description": "An example dataset 1",
            "url": "http://opencoredata.org/id/rdf/geolink1.ttl"
        },
        {
            "@type": "Dataset",
            "description": "An example dataset 2",
            "url": "http://opencoredata.org/id/rdf/cruises.ttl"
        }
    ]
}
`

const bodyTest = ` {
 "@context": {
  "@vocab": "http://schema.org/",
  "re3data": "http://example.org/re3data/0.1/"
 },
 "@id": "http://opencoredata.org/catalogs/geolink",
 "@type": "DataCatalog",
 "dataset": [
  {
   "@type": "Dataset",
   "description": "Collection of cruise data (leg level) for IODP collected for GeoLink",
   "url": "http://opencoredata.org/catalog/geolink/dataset/JRSO_cruises_gl"
  },
  {
   "@type": "Dataset",
   "description": "Collection of Science Party Deployment Information collected for GeoLink",
   "url": "http://opencoredata.org/catalog/geolink/dataset/JRSO_deployments_gl"
  },
  {
   "@type": "Dataset",
   "description": "Collection of cruise data (hole level) for IODP collected for GeoLink",
   "url": "http://opencoredata.org/catalog/geolink/dataset/JRSO_holes_gl"
  }
 ],
 "description": "A catalog of RDF graphs from Open Core Data for GeoLink that align to the GeoLink base ontology",
 "url": "http://opencoredata.org/catalogs/geolink"
}
`

// A simple crawler to go through a given web site (single domain) and starting at
// a JSON-LD document, walk through the tree of documents leveraging JSON-LD framing
func main() {
	fmt.Println("Simple crawler")
	start := time.Now()

	// setup bolt
	SetupBolt()

	// Loop and load the whitelist URLs into the DB to start with
	// registerURL("http://opencoredata.org/catalog/geolink")
	registerURL("http://127.0.0.1:9900/collections/catalogs")
	//_, count := getURLToVisit() // just grab our initial set of URLs to visit, don't worry about a URL string returned
	// TODO   need a count function that is read only...
	count := getCount()
	fmt.Println(count)

	for count > 0 {
		count = caller()
	}

	// showAllKV()
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

// caller function
// working to convert this into something thread safe for concurrent calls.
func caller() int {
	// URL, count := getURLToVisit() // this process is NOT thread safe..  we may go get a URL before it is marked visited..  // need to resolve
	// fmt.Printf("URL to work with: %s with count %d\n", URL, count)

	count := getCount()
	fmt.Println(count)

	if count == 0 {
		return 0
	}

	fmt.Printf("Back in the caller with count: %d\n", count)

	// TODO
	// now use the max routines semephor example from ocdGarden
	// If count > 10  ..  queue all 10 and wait for them to finish....

	// indexer(getURLToVisit()) // just queue everything...

	maxNbConcurrentGoroutines := 1
	if count > 20 {
		maxNbConcurrentGoroutines = 20
	}
	nbJobs := count

	// Dummy channel to coordinate the number of concurrent goroutines.
	// This channel should be buffered otherwise we will be immediately blocked
	// when trying to fill it.
	concurrentGoroutines := make(chan struct{}, maxNbConcurrentGoroutines)
	// Fill the dummy channel with maxNbConcurrentGoroutines empty struct.
	for i := 0; i < maxNbConcurrentGoroutines; i++ {
		concurrentGoroutines <- struct{}{}
	}

	// The done channel indicates when a single goroutine has
	// finished its job.
	done := make(chan bool)
	// The waitForAllJobs channel allows the main program
	// to wait until we have indeed done all the jobs.
	waitForAllJobs := make(chan bool)

	// Collect all the jobs, and since the job is finished, we can
	// release another spot for a goroutine.
	go func() {
		for i := 0; i < nbJobs; i++ {
			<-done
			// Say that another goroutine can now start.
			concurrentGoroutines <- struct{}{}
		}
		// We have collected all the jobs, the program
		// can now terminate
		waitForAllJobs <- true
	}()

	// Try to start nbJobs jobs
	for i := 1; i <= nbJobs; i++ {
		fmt.Printf("ID: %v: waiting to launch!\n", i)
		// Try to receive from the concurrentGoroutines channel. When we have something,
		// it means we can start a new goroutine because another one finished.
		// Otherwise, it will block the execution until an execution
		// spot is available.
		<-concurrentGoroutines
		fmt.Printf("ID: %v: it's my turn!\n", i)
		go func(id int) {
			// DoWork()
			indexer(getURLToVisit())
			fmt.Printf("ID: %v: all done!\n", id)
			done <- true
		}(i)
	}

	// Wait for all jobs to finish
	<-waitForAllJobs

	return count
}

func indexer(URL string) {
	// body = []byte(bodyTest) // TEST REPLACE BODY  replace with test block above....   getDoc is []byte, frameForDataCatalog is string  (review)
	body := extractJSON(URL)

	// Working with framing
	// frameresult := frameForDataCatalog(string(body))
	frameresult := frameForItemList(string(body))

	for _, v := range frameresult {
		// log.Printf("Item %d with URL: %v   \n", k, v.URL)
		// TODO Register the URL in a KV with status set to unvisited
		registerURL(v.URL)
	}

	// TODO  pass body to a function to index it... .

	fmt.Printf("Index the URL: %s\n", URL)

	visitedURL(URL) // TODO make the URL visited
}

// frameForItemList take string and JSON-LD and uses a frame call to extract
// Links from type ItemList of items type DataCatalog  This is then marshalled to a struct...
func frameForItemList(jsonld string) []DataSetStruct {
	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")

	frame := map[string]interface{}{
		"@context": "http://schema.org/",
		"@type":    "DataCatalog",
	}

	var myInterface interface{}
	err := json.Unmarshal([]byte(jsonld), &myInterface)
	if err != nil {
		log.Println("Error when transforming JSON-LD document to interface:", err)
	}

	framedDoc, err := proc.Frame(myInterface, frame, options) // do I need the options set in order to avoid the large context that seems to be generated?
	if err != nil {
		log.Println("Error when trying to frame document", err)
	}

	graph := framedDoc["@graph"]
	// ld.PrintDocument("JSON-LD graph section", graph)  // debug print....
	jsonm, err := json.MarshalIndent(graph, "", " ")
	if err != nil {
		log.Println("Error trying to marshal data", err)
	}

	dss := make([]DataSetStruct, 0)
	err = json.Unmarshal(jsonm, &dss)
	if err != nil {
		log.Println("Error trying to unmarshal data to struct", err)
	}

	// log.Printf("This is the dss:  %v\n", dss)
	return dss
}

// frameForDataCatalog take string and JSON-LD and uses a frame call to extract
// only type DataSet.  This is then marshalled to a struct...
func frameForDataCatalog(jsonld string) []DataSetStruct {
	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")

	frame := map[string]interface{}{
		"@context": "http://schema.org/",
		"@type":    "Dataset",
	}

	// frame := map[string]interface{}{
	// 	"@context": {
	// 		"@vocab":  "http://schema.org/",
	// 		"re3data": "http://example.org/re3data/0.1/",
	// 	},
	// 	"@type": "Dataset",
	// }

	var myInterface interface{}
	err := json.Unmarshal([]byte(jsonld), &myInterface)
	if err != nil {
		log.Println("Error when transforming JSON-LD document to interface:", err)
	}

	framedDoc, err := proc.Frame(myInterface, frame, options) // do I need the options set in order to avoid the large context that seems to be generated?
	if err != nil {
		log.Println("Error when trying to frame document", err)
	}

	graph := framedDoc["@graph"]
	// ld.PrintDocument("JSON-LD graph section", graph)  // debug print....
	jsonm, err := json.MarshalIndent(graph, "", " ")
	if err != nil {
		log.Println("Error trying to marshal data", err)
	}

	dss := make([]DataSetStruct, 0)
	err = json.Unmarshal(jsonm, &dss)
	if err != nil {
		log.Println("Error trying to unmarshal data to struct", err)
	}

	log.Printf("This is the dss:  %v\n", dss)
	return dss
}

func showAllKV() {

	db, err := bolt.Open("walker.db", 0600, &bolt.Options{ReadOnly: true})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte("URLBucket"))

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("key=%s, value=%s\n", k, v)
		}

		return nil
	})

	db.Close() // explicitly close
}

// getDoc DEPRECATED
// simply takes a URL and return the contents of the response body to a byte array
func getDoc(urlstring string) []byte {

	u, err := url.Parse(urlstring)
	if err != nil {
		log.Println(err)
	}

	req, _ := http.NewRequest("GET", u.String(), nil)
	req.Header.Set("Accept", "application/json") // oddly the content-type is ignored for the accept header...
	req.Header.Set("Cache-Control", "no-cache")
	res, _ := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}

	defer res.Body.Close()

	// secs := time.Since(start).Seconds()
	body, _ := ioutil.ReadAll(res.Body)
	return body
}

func extractJSON(urlstring string) []byte {
	resp, err := soup.Get(urlstring)
	if err != nil {
		log.Print(err)
	}
	doc := soup.HTMLParse(resp)

	//     <script type="application/ld+json">
	jsonld := doc.Find("script", "type", "application/ld+json").Text()

	return []byte(jsonld)

	// links := doc.Find("div", "id", "comicLinks").FindAll("a")
	// for _, link := range links {
	// 	fmt.Println(link.Text(), "| Link :", link.Attrs()["href"])
	// }
}

func visitedURL(urlstring string) {

	// open in write mode
	db, err := bolt.Open("walker.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// What should the key be?  Just a simple UID?
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("URLBucket"))
		err := b.Put([]byte(urlstring), []byte("visited"))
		return err
	})

	db.Close() // explicitly close...

	// look for key and set value to "visited"

}

// registerURL take a URL and places it into the bolt KV store.
// While doing so it first ensures that the URL has not already been placed into the KV store
// regardless of whether the URL has been marked as read.
// These URLs come from a framing call onto the JSON-LD for a particular @type
func registerURL(urlstring string) {
	// open in write mode
	db, err := bolt.Open("walker.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// TODO, check if the key is already in the db
	// db.

	// What should the key be?  Just a simple UID?
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("URLBucket"))
		err := b.Put([]byte(urlstring), []byte("unvisited"))
		return err
	})

	db.Close() // explicitly close...

	// check to see if we have been there before..  if not, load and set status unvisited
	// If it's in the KV already ignore..  this is just a register system

}

// getCount  returns the count of unvisited sites
func getCount() int {
	db, err := bolt.Open("walker.db", 0600, &bolt.Options{ReadOnly: true})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	count := 0

	db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte("URLBucket"))

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			if strings.Compare(string(v), "unvisited") == 0 {
				count = count + 1
			}
		}
		return nil
	})

	return count
}

// getURLToVisit just looks into the KV store and looks for a URL to visit...
// func getURLToVisit() (string, int) {
func getURLToVisit() string {

	//  open in read only mode so at not to block and get the first URL we find that
	// is of value "unvisited"
	// db, err := bolt.Open("walker.db", 0600, &bolt.Options{ReadOnly: true})
	db, err := bolt.Open("walker.db", 0600, nil) // trying in READ mode then setting unvisited to "checkedOut"  // at end look for checkedOut not converted to
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var uvsite []byte
	count := 0

	db.Update(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte("URLBucket"))

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			if strings.Compare(string(v), "unvisited") == 0 {
				uvsite = k
				count = count + 1
			}
			// fmt.Printf("key=%s, value=%s\n", k, v)
		}
		err := b.Put(uvsite, []byte("checkedOut")) // set the value of the one we use to checked out...
		return err
	})

	// return string(uvsite), count
	return string(uvsite)

}

// processJSONLD takes the JSONLD document (as a byte array) and processes it to ensure
// it is valid.  It then
func graphJSONLD(jsonld string) {

	baseURI := "https://earthcube.org/cdf/"

	// Create a new graph
	g := rdf2go.NewGraph(baseURI)
	g.Parse(strings.NewReader(jsonld), "application/ld+json")

	// if err != nil {
	// 	// deal with err
	// }
}

// jsonLDToRDF take a JSON-LD string and convert it to n-triples and returns it.
func jsonLDToRDF(jsonld string) string {
	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")
	options.Format = "application/nquads"

	var myInterface interface{}
	err := json.Unmarshal([]byte(jsonld), &myInterface)
	if err != nil {
		log.Println("Error when transforming JSON-LD document to interface:", err)
	}

	triples, err := proc.ToRDF(myInterface, options)
	if err != nil {
		log.Println("Error when transforming JSON-LD document to RDF:", err)
		return err.Error()
	}

	return triples.(string)
}

func SetupBolt() {

	db, err := bolt.Open("walker.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// You can also create a bucket only if it doesn't exist by using the Tx.CreateBucketIfNotExists()
	db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("URLBucket"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		log.Printf("Bucket created %v", b.FillPercent)
		return nil
	})
}
