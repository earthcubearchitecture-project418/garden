package indexer

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/kazarena/json-gold/ld"
	"github.com/rs/xid"
)

// GraphIndexer indexes data in a graph
func GraphIndexer(jsonld string) string {

	// Steps
	// 1 take the JSON-LD..  concvert to RDF with GUID bnodes if present (buitl)
	// 2 Place this into a master RDF graph like with triple store (this could get large..  do on a domain basis?)
	// 3 Make the graph available for loading into a triple store. (named graphs as run names/dates)
	// 4 If (3) never search the default graph, always a named graph.

	nq, _ := jsonLDToNQ(jsonld) // TODO replace with NQ from isValid function..  saving time..
	rdf := globalUniqueBNodes(nq)

	// TODO
	// Load these into a graph like in triplestore or write to a KV store or to the
	// triple store?
	err := writeRDF(rdf)
	if err != nil {
		log.Println("error writing out RDF")
	}

	return "done"

}

func writeRDF(rdf string) error {

	// for now just append to a file..   later I will send to a triple store
	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile("./nquads.rdf", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte(rdf)); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

	return err // always nil,  we will never get here with FATAL..   leave for test..  but later remove to log only
}

func jsonLDToNQ(jsonld string) (string, error) {
	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")
	options.Format = "application/nquads"

	var myInterface interface{}
	err := json.Unmarshal([]byte(jsonld), &myInterface)
	if err != nil {
		log.Println("Error when transforming JSON-LD document to interface:", err)
		return "", err
	}

	nq, err := proc.ToRDF(myInterface, options) // returns triples but toss them, we just want to see if this processes with no err
	if err != nil {
		log.Println("Error when transforming JSON-LD document to RDF:", err)
		return "", err
	}

	return nq.(string), err
}

func globalUniqueBNodes(nq string) string {

	scanner := bufio.NewScanner(strings.NewReader(nq))
	// make a map here to hold our old to new map
	m := make(map[string]string)

	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		// parse the line
		split := strings.Split(scanner.Text(), " ")
		sold := split[0]
		oold := split[2]

		if strings.HasPrefix(sold, "_:") { // we are a blank node
			// check map to see if we have this in our value already
			if _, ok := m[sold]; ok {
				// fmt.Printf("We had %s, already\n", sold)
			} else {
				guid := xid.New()
				snew := fmt.Sprintf("_:b%s", guid.String())
				m[sold] = snew
			}
		}

		// scan the object nodes too.. though we should find nothing here.. the above wouldn't
		// eventually find
		if strings.HasPrefix(oold, "_:") { // we are a blank node
			// check map to see if we have this in our value already
			if _, ok := m[oold]; ok {
				// fmt.Printf("We had %s, already\n", oold)
			} else {
				guid := xid.New()
				onew := fmt.Sprintf("_:b%s", guid.String())
				m[oold] = onew
			}
		}
		// triple := tripleBuilder(split[0], split[1], split[3])
		// fmt.Println(triple)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//fmt.Println(m)

	filebytes := []byte(nq)

	for k, v := range m {
		// fmt.Printf("Replace %s with %v \n", k, v)
		filebytes = bytes.Replace(filebytes, []byte(k), []byte(v), -1)
	}

	return string(filebytes)
}
