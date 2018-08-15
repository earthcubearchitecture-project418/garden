package crawl

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"earthcube.org/Project418/crawler/digester"
	"github.com/blevesearch/bleve"
)

// URLSet takes a URL to a sitemap and parses out the content.
// The entries are digested one by one (later N conncurrent)
type URLSet struct {
	XMLName xml.Name  `xml:"urlset"`
	URL     []URLnode `xml:"url"`
}

// URLnode sub node struct
type URLnode struct {
	XMLName     xml.Name `xml:"url"`
	Loc         string   `xml:"loc"`
	Description string   `xml:"description"`
}

// ProcessDomain will take a domain URL and
// go get the sitemap.xml file there...
func ProcessDomain(url, textindex string, spatial bool) error {
	// Check that it is XML and can be processed into the struct above
	siteArray := IngestSitemapXML(url)
	count := 0
	total := len(siteArray.URL)

	// TODO..   open the bleve index here once and pass by reference to text
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
			log.Printf("PROCESS URL error, %v", err)
		}
		log.Printf("CRAWLER report:  URL %s done with error code: %v\n", siteArray.URL[entry].Loc, err)
	}

	index.Close()

	return nil
}

// IngestSitemapXML validates the XMl format of the sitemap and
// reads each entry into a struct array that is sent back
func IngestSitemapXML(url string) URLSet {
	// read the sitemap into a []SiteMapEntry

	var client http.Client

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Print(err) // not even being able to make a req instance..  might be a fatal thing?
	}

	req.Header.Set("User-Agent", "EarthCube_DataBot/1.0")

	resp, err := client.Do(req)
	if err != nil {
		log.Print("Error reading sitemap: %s", err)
	}
	defer resp.Body.Close()

	// var bodyString string
	var bodyBytes []byte
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Print(err)
		}
		// bodyString = string(bodyBytes)
	}

	// Process XML into a struct
	var sitemap URLSet
	xml.Unmarshal(bodyBytes, &sitemap)

	return sitemap
}

// IngestSitemapText reads text based sitemap
func IngestSitemapText(url string) string {

	return "make this a real thing"
}
