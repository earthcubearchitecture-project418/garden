package main

import (
	"archive/zip"
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/piprate/json-gold/ld"
	"github.com/rs/xid"
)

// TODO ..  need to bring in my unique bnode code..

func main() {
	fmt.Println("Scan a directory of FDP packages, look for meta/schemaorg.json and build the triples from them")

	// file := "/media/fils/seagate/packages/c6b45e0a565900d22812f04e96b8dcc1fcdc3229d5037b69fc212cfd8d62013a.zip"

	files, err := ioutil.ReadDir("/media/fils/seagate/packages")
	if err != nil {
		log.Fatal(err)
	}

	var buffer bytes.Buffer

	for _, f := range files {
		fp := fmt.Sprintf("/media/fils/seagate/packages/%s", f.Name())

		t, err := unzip(fp)
		if err != nil {
			fmt.Println(err)
		}

		nq, err := jsonLDToNQ(t, fp)
		if err != nil {
			fmt.Println(err)
		}

		nq = globalUniqueBNodes(nq)
		buffer.WriteString(nq)
	}

	fmt.Println(buffer.String())

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
