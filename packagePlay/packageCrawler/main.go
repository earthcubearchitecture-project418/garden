package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-resty/resty"
)

// Manifest is the struct for the manifest from the data package
// do not need the full datapackage.json, just the file manifest
type Manifest struct {
	Profile   string `json:"profile"`
	Resources []struct {
		Encoding string `json:"encoding"`
		Name     string `json:"name"`
		Path     string `json:"path"`
		Profile  string `json:"profile"`
	} `json:"resources"`
}

func main() {
	fmt.Println("Crawl a package at a URL and walk the items in the site")

	// read the packages.json file via http call
	// generate a struct with the names of the files (or URLs) to load
	// iterate and pull each individual file in the packages and read []byte length for now (and time this whole function)

	id := "8448c71edc22a06a26501a967223e5502dd4678be06c5761440167229ec9b715"
	ru := fmt.Sprintf("http://opencoredata.org/pkg/id/%s/resource", id)

	m := readManifest(ru, "datapackage.json")
	fmt.Println(m)

	ms := parsePackage(m)
	for _, v := range ms.Resources {
		// fmt.Println(v.Path)
		// fmt.Println(v.Name)
		l := getBytes(ru, v.Path)
		fmt.Printf("%d bytes from %s?key=%s \n", l, ru, v.Path)
	}

}

func getBytes(url, key string) int {
	resp, err := resty.R().
		SetQueryParams(map[string]string{
			"key": key,
		}).
		Get(url)
	if err != nil {
		log.Println(err)
	}
	return len(resp.Body())
}

func parsePackage(j string) Manifest {
	m := Manifest{}
	json.Unmarshal([]byte(j), &m)
	return m
}

func readManifest(url, key string) string {
	resp, err := resty.R().
		SetQueryParams(map[string]string{
			"key": key,
		}).
		Get(url)
	if err != nil {
		log.Println(err)
	}

	return resp.String()
}
