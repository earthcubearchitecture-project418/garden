package acquire

import (
	"bufio"
	"encoding/json"
	"log"
	"os"

	"earthcube.org/Project418/garden/summoner/sitemaps"
	"earthcube.org/Project418/garden/summoner/utils"
	"github.com/kazarena/json-gold/ld"
)

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

	_, err = proc.ToRDF(myInterface, options) // returns triples but toss them, just validating
	if err != nil {
		log.Println("Error when transforming JSON-LD document to RDF:", err)
		return err
	}

	return err
}

// ResourceURLs
func ResourceURLs(domains []string) map[string]sitemaps.URLSet {
	m := make(map[string]sitemaps.URLSet) // make a map

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

// DomainList
func DomainList(sources string) ([]string, error) { // return map(string[]string)

	log.Printf("Opening source list file: %s \n", sources)

	var domains []string

	file, err := os.Open(sources)
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
