package main

import (
	"fmt"
	"log"

	"github.com/twpayne/go-geom/encoding/geojson"
	"github.com/twpayne/go-geom/encoding/wkt"
)

// https://github.com/twpayne/go-geom/blob/master/examples/postgis/main.go

func main() {
	fmt.Println("Go-geom test")

	// read in geojson
	// make a feature struct and unmarshal the json feature into it

	f := geojson.Feature{}
	fc := geojson.FeatureCollection{}

	err := f.UnmarshalJSON([]byte(ftr()))
	if err != nil {
		log.Println(err)
	}

	s, err := wkt.Marshal(f.Geometry)
	if err != nil {
		log.Println(err)
	}

	log.Println(s)

	err = fc.UnmarshalJSON([]byte(ftrcol()))
	if err != nil {
		log.Println(err)
	}

	// TODO..   loop on features
	// each feature / wkt string is a triple in the graph...
	// ref https://jena.apache.org/documentation/query/spatial-query.html#builtin-geo-predicates
	for i := range fc.Features {
		s2, err := wkt.Marshal(fc.Features[i].Geometry)
		if err != nil {
			log.Println(err)
		}
		log.Println(s2)
	}

}

func ftr() string {
	gj := `{
		  "type": "Feature",
		  "geometry": {
			      "type": "Point",
				      "coordinates": [125.6, 10.1]
		  },
		  "properties": {
			      "name": "Dinagat Islands"
		  }
	}`

	return gj
}

func simpleftrcol() string {
	sfc := `{
		"type": "FeatureCollection",
		"features": [{
				"type": "Feature",
				"geometry": {
					"type": "Point",
					"coordinates": [
						-77.32,
						39.8

					]

				},
				"properties": {
					"URL": "http://wiki.linked.earth/Ant-DomeF1993.Uemura.2014"

				}

			}
		]
	}
	`

	return sfc
}

func ftrcol() string {

	fc := `{
		"type": "FeatureCollection",
		"features": [{
				"type": "Feature",
				"geometry": {
					"type": "Point",
					"coordinates": [
						-77.32,
						39.7

					]

				},
				"properties": {
					"URL": "http://wiki.linked.earth/Ant-DomeF1993.Uemura.2014"

				}

			},
			{
				"type": "Feature",
				"geometry": {
					"type": "Point",
					"coordinates": [
						-77.32,
						39.7

					]

				},
				"properties": {
					"URL": "http://wiki.linked.earth/Ant-DomeF2001.Uemura.2008"

				}

			},
			{
				"type": "Feature",
				"geometry": {
					"type": "Point",
					"coordinates": [
						-33.62,
						45.16

					]

				},
				"properties": {
					"URL": "http://opencoredata.org/id/dataset/97c6a8b3-e764-488e-bd62-85f919f9aa0d"

				}

			},
			{
				"type": "Feature",
				"geometry": {
					"type": "Point",
					"coordinates": [
						-33.46,
						39.37

					]

				},
				"properties": {
					"URL": "http://opencoredata.org/id/dataset/87fe348b-1abe-45dd-a415-d78008b6093b"

				}

			},
			{
				"type": "Feature",
				"geometry": {
					"type": "Point",
					"coordinates": [
						-33.46,
						39.37

					]

				},
				"properties": {
					"URL": "http://opencoredata.org/id/dataset/022f8b64-4741-482a-b053-56fa7b75acca"

				}

			},
			{
				"type": "Feature",
				"geometry": {
					"type": "Point",
					"coordinates": [
						-33.62,
						45.16

					]

				},
				"properties": {
					"URL": "http://opencoredata.org/id/dataset/21d9a90c-bc82-49cb-b66b-73863655e375"

				}

			},
			{
				"type": "Feature",
				"geometry": {
					"type": "Point",
					"coordinates": [
						-33.63,
						45.01

					]

				},
				"properties": {
					"URL": "http://opencoredata.org/id/dataset/29104a7f-0df0-40ab-8466-e999d590b39f"

				}

			},
			{
				"type": "Feature",
				"geometry": {
					"type": "Point",
					"coordinates": [
						-33.46,
						39.37

					]

				},
				"properties": {
					"URL": "http://opencoredata.org/id/dataset/ae9934a2-78e6-455d-aea5-acb6d6b4927d"

				}

			},
			{
				"type": "Feature",
				"geometry": {
					"type": "Point",
					"coordinates": [
						-33.46,
						39.37

					]

				},
				"properties": {
					"URL": "http://opencoredata.org/id/dataset/fe7525c2-b8d9-4a8a-b0de-b020451d4b84"

				}

			},
			{
				"type": "Feature",
				"geometry": {
					"type": "Point",
					"coordinates": [
						-33.46,
						39.37

					]

				},
				"properties": {
					"URL": "http://opencoredata.org/id/dataset/2df05cbe-cad1-4b1f-9b58-58bb1f8062a2"

				}

			},
			{
				"type": "Feature",
				"properties": {},
				"geometry": {
					"type": "Polygon",
					"coordinates": [
						[
							[
								-59.58984374999999,
								33.94335994657882

							],
							[
								-47.02148437499999,
								33.94335994657882

							],
							[
								-47.02148437499999,
								40.38002840251183

							],
							[
								-59.58984374999999,
								40.38002840251183

							],
							[
								-59.58984374999999,
								33.94335994657882

							]

						]

					]

				}

			}

		]

	}`

	return fc
}
