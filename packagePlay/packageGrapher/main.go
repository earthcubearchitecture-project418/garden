package main

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/piprate/json-gold/ld"
)

func main() {
	fmt.Println("Scan a directory of FDP packages, look for meta/schemaorg.json and build the triples from them")

	// read the metadata/schemaorg.json file from the zip
	// convert it to triples (well..  duh)

	file := "/media/fils/seagate/packages/c6b45e0a565900d22812f04e96b8dcc1fcdc3229d5037b69fc212cfd8d62013a.zip"

	t, err := unzip(file)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t)

	fmt.Println("-------------------------------------------")

	nq, err := jsonLDToNQ(t, file)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(nq)
}

func unzip(archive string) (string, error) {
	reader, err := zip.OpenReader(archive)
	if err != nil {
		return "", err
	}

	for _, file := range reader.File {
		if file.Name == "metadata/schemaorg.json" {
			fileReader, err := file.Open()
			if err != nil {
				return "", err
			}
			defer fileReader.Close()

			b, err := ioutil.ReadAll(fileReader)
			return string(b), nil
		}
	}

	return "", nil
}

func jsonLDToNQ(jsonld, urlval string) (string, error) {
	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")
	options.Format = "application/nquads"

	var myInterface interface{}
	err := json.Unmarshal([]byte(jsonld), &myInterface)
	if err != nil {
		log.Println("Error when transforming %s JSON-LD document to interface:", urlval, err)
		return "", err
	}

	nq, err := proc.ToRDF(myInterface, options) // returns triples but toss them, we just want to see if this processes with no err
	if err != nil {
		log.Println("Error when transforming %s  JSON-LD document to RDF:", urlval, err)
		return "", err
	}

	return nq.(string), err
}
