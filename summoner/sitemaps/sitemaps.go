package sitemaps

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
)

// URLSet takes a URL to a sitemap and parses out the content.
// The entries are digested one by one (later N conncurrent)
type URLSet struct {
	XMLName xml.Name  `xml:"urlset"`
	URL     []URLNode `xml:"url"`
}

// URLNode sub node struct
type URLNode struct {
	XMLName     xml.Name `xml:"url"`
	Loc         string   `xml:"loc"`
	Description string   `xml:"description"`
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
		log.Printf("Error reading sitemap: %s", err)
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
