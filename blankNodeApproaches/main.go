package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/rs/xid"
)

func main() {
	manualNodeWriter()
}

func manualNodeWriter() {
	// read the file line by line
	file, err := os.Open("testData/bcodmo.nq")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// make a map here to hold our old to new map
	m := make(map[string]string)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
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

	fmt.Println(m)

	filebytes, err := ioutil.ReadFile("testData/bcodmo.nq")
	if err != nil {
		log.Fatal()
	}

	for k, v := range m {
		// fmt.Printf("Replace %s with %v \n", k, v)
		filebytes = bytes.Replace(filebytes, []byte(k), []byte(v), -1)
	}

	fmt.Println(string(filebytes))
}
