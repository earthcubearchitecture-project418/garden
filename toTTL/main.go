package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/knakk/rdf"
	"github.com/piprate/json-gold/ld"
)

func main() {
	fmt.Println("convert JSON-ld to TTL")

	s, err := JSONLDToTTL(exampleJLD(), "http://example.org")

	if err != nil {
		fmt.Println("Things went bad")
	}

	fmt.Println(s)
}

func JSONLDToTTL(jsonld, urlval string) (string, error) {
	// Sad that rdf2go has a bug in jsonld around blank nodes...
	// I can convert to NQ above..   I guess use knakk to then convert to turtle  (since json-ld gold
	// also does not support converting to TTL.

	nq, err := JSONLDToNQ(jsonld, urlval)
	if err != nil {
		log.Println(err)
	}

	var b bytes.Buffer
	writer := bufio.NewWriter(&b)

	dec := rdf.NewTripleDecoder(strings.NewReader(nq), rdf.NTriples)
	tr, err := dec.DecodeAll()
	if err != nil {
		log.Println(err)
	}

	enc := rdf.NewTripleEncoder(writer, rdf.Turtle)
	err = enc.EncodeAll(tr)
	if err != nil {
		log.Println(err)
	}
	writer.Flush()

	return b.String(), nil
}

func JSONLDToNQ(jsonld, urlval string) (string, error) {
	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")
	options.Format = "application/nquads"

	var myInterface interface{}
	err := json.Unmarshal([]byte(jsonld), &myInterface)
	if err != nil {
		log.Printf("Error when transforming %s JSON-LD document to interface: %v", urlval, err)
		return "", err
	}

	nq, err := proc.ToRDF(myInterface, options) // returns triples but toss them, we just want to see if this processes with no err
	if err != nil {
		log.Printf("Error when transforming %s  JSON-LD document to RDF: %v", urlval, err)
		return "", err
	}

	return nq.(string), err
}

func exampleJLD() string {

	jld := `
[{
	"@context": "http://schema.org",
	"@type": "WebSite",
	"@id": "http://opencoredata.org",
	"url": "https://www.opencoredata.org/",
	"keywords": "geochemistry, paleontology, geophysics, geoscience, scientific drilling, ocean drilling, continental drilling, paleomagnetism",
	"description": "opencoredata is a data infrastructure focused on making data from JRSO and CSDCO scientific drilling data facilities available in a semantic and Linked Open Data manner.",
	"potentialAction": {
		"@type": "SearchAction",
		"@id": "http://opencoredata.org#searchAction",
		"target": "https://www.opencoredata.org/search?q={search_term_string}",
		"query-input": "required name=search_term_string"
	}
}, {
	"@context": {
		"@vocab": "http://schema.org/",
		"re3data": "http://example.org/re3data/0.1/"
	},
	"@type": "Organization",
	"@id": "http://opencoredata.org/id/facilityinfo",
	"name": "Open Core Data",
	"re3data:datalicense": "http://opencoredata.org/datapolicy.html",
	"re3data:keyword": "geochemistry, paleontology, geophysics, geoscience, scientific drilling, ocean drilling, continental drilling, paleomagnetism",
	"contactPoint": {
		"@type": "ContactPoint",
		"@id": "http://opencoredata.org/id/facilityinfo#contactPoint",
		"name": "Douglas Fils",
		"email": "dfilsAToceanleadershipDOTorg",
		"url": "http://orcid.org/0000-0002-2257-9127",
		"contactType": "technical support"
	},
	"url": "http://www.opencoredata.org",
	"sameAs": "http://www.re3data.org/repository/r3d100012071",
	"funder": {
		"@type": "Organization",
		"@id": "http://opencoredata.org/id/facilityinfo#funder",
		"name": "NSF",
		"url": "http://www.nsf.gov"
	},
	"memberOf": {
		"@type": "ProgramMembership",
		"@id": "http://opencoredata.org/id/facilityinfo#programMembership",
		"programName": "EarthCube CDF Registry",
		"hostingOrganization": {
			"@type": "Organization",
			"@id": "http://opencoredata.org/id/facilityinfo#programMembershipHost",
			"name": "RE3Data",
			"url": "http://www.re3data.org"
		}
	},
	"potentialAction": [{
		"@type": "SearchAction",
		"@id": "http://opencoredata.org/id/facilityinfo#potentialActionSwagger",
		"target": {
			"@type": "EntryPoint",
			"@id": "http://opencoredata.org/id/facilityinfo#potentialActionSwaggerEntryPoint",
			"urlTemplate": "http://opencoredata.org/apidocs.json",
			"description": "Swagger 1.2 description document",
			"httpMethod": "GET"
		}
	}, {
		"@type": "SearchAction",
		"@id": "http://opencoredata.org/id/facilityinfo#potentialActionSPARQL",
		"target": {
			"@type": "EntryPoint",
			"@id": "http://opencoredata.org/id/facilityinfo#potentialActionSPARQLEntryPpint",
			"urlTemplate": "http://opencoredata.org/sparql",
			"description": "SPARQL endpoint",
			"httpMethod": "GET"
		}
	}, {
		"@type": "SearchAction",
		"@id": "http://opencoredata.org/id/facilityinfo#potentialActionVOID",
		"target": {
			"@type": "EntryPoint",
			"@id": "http://opencoredata.org/id/facilityinfo#potentialActionVOIDEntryPoint",
			"urlTemplate": "http://opencoredata.org/rdf/void.ttl",
			"description": "VoID document",
			"httpMethod": "GET"
		}
	}],
	"mainEntityOfPage": [{
		"@type": "DigitalDocument",
		"@id": "http://opencoredata.org/id/facilityinfo#mainEntityOfPageSwaggerUI",
		"fileformat": "HTML",
		"keywords": "Swagger",
		"description": "User interface for the swagger document",
		"url": "http://opencoredata.org/common/swagger-ui/"
	},
	{
		"@type": "DigitalDocument",
		"@id": "http://opencoredata.org/id/facilityinfo#mainEntityOfPageSwaggerUI",
		"fileformat": "HTML",
		"keywords": "CatalogMap",
		"description": "A map of all catalog files available",
		"url": "http://opencoredata.org/maps/catalog/"
	}]
}]`

	return jld
}
