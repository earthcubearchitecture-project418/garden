package main

import (
"fmt"
"log"
"reflect"
"github.com/jonas-p/go-shp"
)


func main() {


// open a shapefile for reading
shape, err := shp.Open("assets/ne_10m_graticules_10.shp")
if err != nil { log.Fatal(err) }
defer shape.Close()

// fields from the attribute table (DBF)
fields := shape.Fields()

// loop through all features in the shapefile
for shape.Next() {
	n, p := shape.Shape()

	// print feature
	fmt.Println(reflect.TypeOf(p).Elem(), p.BBox())

	// print attributes
	for k, f := range fields {
		val := shape.ReadAttribute(n, k)
		fmt.Printf("\t%v: %v\n", f, val)
	}
	fmt.Println()
}
}
