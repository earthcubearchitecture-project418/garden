package main

import (
	"errors"
	"log"
	"strconv"

	"../../internal/framing"
	"github.com/coyove/jsonbuilder"

	// "github.com/kpawlik/geojson"
	geojson "github.com/paulmach/go.geojson"
)

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	sfr := framing.SpatialFrame(jld)
	log.Println(sfr)

	featureCollection := geojson.NewFeatureCollection()

	// geojson, err := addSchemaOrgPoint(featureCollection, "http://foo.org/id/1", geo.Longitude, geo.Latitude)
	geojson, err := addSchemaOrgPoint(featureCollection, "http://foo.org/id/1", "10", "10")
	if err != nil {
		log.Println(err)
	}

	log.Println(geojson)

	// Once we have a framed result we can
}

func addSchemaOrgPoint(fc *geojson.FeatureCollection, idToUse string, lon string, lat string) (string, error) {
	x, err := strconv.ParseFloat(lon, 64)
	if err != nil {
		return "", err
	}
	y, err := strconv.ParseFloat(lat, 64)
	if err != nil {
		return "", err
	}

	json := jsonbuilder.Object()
	json.Set("type", "Point").Set("coordinates", jsonbuilder.Array(x, y))
	geojson := json.Marshal()
	//log.Println("POINT: ", geojson)
	return addGeoJSON(fc, idToUse, []byte(geojson))
}

func addGeoJSON(fc *geojson.FeatureCollection, idToUse string, gjb []byte) (string, error) {
	geom, err := geojson.UnmarshalGeometry(gjb)
	if err != nil {
		return "", err
	}

	switch geom.Type {
	case "Point", "MultiPoint", "LineString", "MultiLineString", "Polygon", "MultiPolygon", "GeometryCollection":
		fc.AddFeature(geojson.NewFeature(geom))
		break

	case "Feature":
		feature, ferr := geojson.UnmarshalFeature(gjb)
		if ferr != nil {
			return "", err
		}
		fc.AddFeature(feature)
		break

	case "FeatureCollection":
		fcoll, fcerr := geojson.UnmarshalFeatureCollection(gjb)
		if fcerr != nil {
			return "", fcerr
		}
		for _, feature := range fcoll.Features {
			fc.AddFeature(feature)
		}
		break

	default:
		return "", errors.New("UNKNOWN GeoJSON Type: " + string(geom.Type))
	}

	return string(gjb), nil
}

const jld = `{
"@context": {
	    "@vocab": "http://schema.org/",
		    "datacite": "http://purl.org/spar/datacite/"
},
	"@type": "Dataset",
	"name": "Removal of organic carbon by natural bacterioplankton communities as a function of pCO2 from laboratory experiments between 2012 and 2016",
	"spatialCoverage": {
		"@type": "Place",
		"geo": {
			"@type": "GeoCoordinates",
			"latitude": 39.3280,
			"longitude": 120.1633
		}
	}
}
	  `
