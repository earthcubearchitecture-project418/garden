package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/deiu/rdf2go"
)

type ProviderInfo struct {
	Name   string
	Lable  string
	MailTo string
	URLs   []string
}

// This is some test code that will go into the miller

func main() {
	urls := []string{"http://foo.org/1", "http://foo.org/2", "http://foo.org/3"}
	pi := ProviderInfo{Name: "Open Core Data", Lable: "Open Core Data", MailTo: "info@opencore.org", URLs: urls}
	buildGraph(pi)
}

func buildGraph(pi ProviderInfo) {

	// make UUID here to make the baseuri uniqie

	// Set a base URI
	baseUri := "https://provisium.io/id/001:UUID"
	g := rdf2go.NewGraph(baseUri)

	// r is of type io.Reader
	bt, ot := baseTriples(pi)
	err := g.Parse(strings.NewReader(bt), "text/turtle")
	err = g.Parse(strings.NewReader(ot), "text/turtle")
	if err != nil {
		log.Println(err)
	}

	for item := range pi.URLs {
		// Add a few triples to the graph
		noderef := fmt.Sprintf("http://provisium.io#lpref%d", item)
		triple1 := rdf2go.NewTriple(rdf2go.NewResource("http://provisium.io#dataset"), rdf2go.NewResource("prov:wasDerivedFrom"), rdf2go.NewResource(noderef))
		g.Add(triple1)
		triple2 := rdf2go.NewTriple(rdf2go.NewResource("http://provisium.io#processingActivity1"), rdf2go.NewResource("prov:used"), rdf2go.NewResource(noderef))
		g.Add(triple2)
		triple3 := rdf2go.NewTriple(rdf2go.NewResource(noderef), rdf2go.NewResource("rdf:type"), rdf2go.NewResource("eos:product"))
		g.Add(triple3)
		triple4 := rdf2go.NewTriple(rdf2go.NewResource(noderef), rdf2go.NewResource("rdf:type"), rdf2go.NewResource("prov:Entity"))
		g.Add(triple4)
		triple5 := rdf2go.NewTriple(rdf2go.NewResource(noderef), rdf2go.NewResource("dcat:url"), rdf2go.NewLiteral("URL there as a CDATA"))
		g.Add(triple5)
		triple6 := rdf2go.NewTriple(rdf2go.NewResource(noderef), rdf2go.NewResource("rdfs:label"), rdf2go.NewLiteral("JSON-LD from landing page"))
		g.Add(triple6)
		triple7 := rdf2go.NewTriple(rdf2go.NewResource(noderef), rdf2go.NewResource("prov:wasAttributedTo"), rdf2go.NewResource("http://provisium#esso"))
		g.Add(triple7)
	}

	// Dump graph contents to NTriples
	out := g.String()
	fmt.Println(out)
}

func baseTriples(pi ProviderInfo) (string, string) {

	// Would be nice to have a URL here for them too..  maybe other data as well
	orgtriples := fmt.Sprintf(`@prefix rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#> .
@prefix xsd: <http://www.w3.org/2001/XMLSchema#> .
@prefix foaf: <http://xmlns.com/foaf/0.1/> .
@prefix prov: <http://www.w3.org/ns/prov#> .
@prefix eos: <http://esipfed.org/prov/eos#> .
@prefix dcat: <https://www.w3.org/ns/dcat> .
@prefix rdfs: <http://www.w3.org/2000/01/rdf-schema#> .
@prefix : <http://provisium.io#> .

:ocd
    a prov:Agent, prov:Organization ;
    rdfs:label "%s"^^xsd:string ;
    foaf:givenName "%s" ;
    foaf:mbox <mailto:%s> .
	`, pi.Lable, pi.Name, pi.MailTo)

	bt := `@prefix rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#> .
@prefix xsd: <http://www.w3.org/2001/XMLSchema#> .
@prefix foaf: <http://xmlns.com/foaf/0.1/> .
@prefix prov: <http://www.w3.org/ns/prov#> .
@prefix eos: <http://esipfed.org/prov/eos#> .
@prefix dcat: <https://www.w3.org/ns/dcat> .
@prefix rdfs: <http://www.w3.org/2000/01/rdf-schema#> .
@prefix : <http://provisium.io#> .

# Will need to honor and deference this URI to a landing page for this prov data
<http://geodex.org/id/prov/UUID/ORG_INDEXED>
    a prov:Bundle, prov:Entity ;
    rdfs:label "A collection of provenance related to the creation of a P418 index"^^xsd:string ;
    prov:generatedAtTime "2018-02-14T00:00:00Z"^^xsd:dateTime ;
    prov:wasAttributedTo :processingActivity1 .

:esso
    a prov:Agent, prov:Organization ;
    rdfs:label "EarthCube Science Support Office"^^xsd:string ;
    foaf:givenName "USGS" ;
    # need URL
    foaf:mbox <mailto:info@earthcube.org> .

:processingCode
    a eos:software, prov:Entity ;
    rdfs:label "EarthCube Project 418 Indexer"^^xsd:string ;
    # what voc to use to link to software repo?  (other ID?)  just need a URl for now
    prov:wasAttributedTo :esso .

:dataset
    a eos:product, prov:Entity ;
    rdfs:label "Dataset included spatial, text and graph results from the activity"^^xsd:string ;
    prov:wasAttributedTo :esso ;
    # prov:wasDerivedFrom :page1 ;  # what goes here?  the collection?
    prov:wasGeneratedBy :processingActivity1 .

:processingActivity1
    a eos:processStep, prov:Activity ;
    rdfs:label "Generation of indexes (spatial, text, graph) from the processed pages"^^xsd:string ;
    prov:endedAtTime "2011-07-14T02:02:02Z"^^xsd:dateTime ;
    prov:startedAtTime "2011-07-14T01:01:01Z"^^xsd:dateTime ;
    prov:used :processingCode ;  
    # prov:used :page1 ;  
    prov:wasAssociatedWith :esso .
`

	return bt, orgtriples
}
