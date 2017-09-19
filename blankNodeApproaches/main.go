package main

import (
	"fmt"
	"log"
	"os"

	rdf "github.com/knakk/rdf"

	jgraph "github.com/Callidon/joseki/graph"
	jrdf "github.com/Callidon/joseki/rdf"
)

func main() {
	//   josekiTest()
	knaccRDF()
}

func knaccRDF() {
	fmt.Println("Read and write RDF")

	// Open the input file
	inFile, err := os.Open("./testData/bcodmo.nq")
	defer inFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Decode the existing triples
	var inoutFormat rdf.Format
	inoutFormat = rdf.NTriples // FormatTTL
	dec := rdf.NewTripleDecoder(inFile, inoutFormat)
	tr, err := dec.DecodeAll()

	fmt.Print(tr)
}

func josekiTest() {
	fmt.Println("Joseki testing")
	graph := jgraph.NewTreeGraph()

	// Datas stored in a file can be easily loaded into a graph
	graph.LoadFromFile("testData/bcodmo.nq", "nt") // the nq file is really just a valid nt file...

	// Let's fetch the titles of all the books in our graph !
	// subject := jrdf.NewVariable("title")
	predicate := jrdf.NewURI("http://schema.org/name")
	// object := jrdf.NewURI("https://schema.org/Book")
	for bindings := range graph.FilterSubset(nil, predicate, nil, 0, 0) {
		fmt.Println(bindings)
	}
}
