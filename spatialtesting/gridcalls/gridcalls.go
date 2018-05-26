package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/go-resty/resty"
)

// Feature struct holds an ID associated with a hexagon geojson feature
type Feature struct {
	Type       string `json:"type"`
	ID         int    `json:"id"`
	Properties struct {
		Dummy float64 `json:"dummy"`
	} `json:"properties"`
	Geometry struct {
		Type        string        `json:"type"`
		Coordinates [][][]float64 `json:"coordinates"`
	} `json:"geometry"`
	Count   int
	GeoJSON string
}

// CallResponse holds the results of the web services call
type CallResponse struct {
	Type     string `json:"type"`
	Features []struct {
		Type       string `json:"type"`
		Properties struct {
			URL string `json:"URL"`
		} `json:"properties"`
	} `json:"features"`
}

func main() {
	fmt.Println("Hexagon grid calls")
	// open the file
	f, err := os.Open("./featureSetfull.txt")
	if err != nil {
		fmt.Printf("Error opening file: %s\n", err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	fset := []Feature{}
	for scanner.Scan() {
		st := scanner.Text()
		ft := Feature{}
		json.Unmarshal([]byte(st), &ft)
		ft.GeoJSON = st
		fset = append(fset, ft)
	}

	for _, elem := range fset {
		c, err := countCall(elem.GeoJSON)
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Printf("%d :  %d     \t\t%v\n", elem.ID, c, err)
		fmt.Printf("%d, %d\n", elem.ID, c)
	}
}

func countCall(feature string) (int, error) {

	resp, err := resty.R().
		SetQueryParams(map[string]string{
			"geowithin": feature,
			"remove":    "iedadata.org",
		}).
		Get("http://geodex.org/api/v1/spatial/search/object")
	if err != nil {
		log.Println(err)
	}

	// fmt.Printf("%v    %d \n", resp.StatusCode(), len(resp.Body()))
	r := CallResponse{}
	err = json.Unmarshal(resp.Body(), &r)

	return len(r.Features), err
}
