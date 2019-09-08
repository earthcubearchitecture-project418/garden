package main

import (
	"log"

	"../../internal/frameparser"
	"../../internal/framing"
	"../../internal/spatial"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	sfr := framing.ProjSpatial(jld2)
	lat, long, err := frameparser.ProjLatLong(sfr)
	if err != nil {
		log.Println(err)
	}

	for x, _ := range lat {
		wkt, err := spatial.LatLongWKT(lat[x], long[x])
		if err != nil {
			log.Println(err)
		}
		log.Println(wkt)
	}

	gjs, err := spatial.LatLongGJS(lat, long)
	if err != nil {
		log.Println(err)
	}
	log.Println(gjs)

	lat, long, err = frameparser.ProjLatLong(framing.ProjSpatial(jld2))
	if err != nil {
		log.Println(err)
	}
	log.Printf("JSLD2: Lat: %s  Long: %s", lat, long)

	// Some tests to remove....
	lat, long, err = frameparser.DataLatLong(framing.DataSpatial(jld))
	if err != nil {
		log.Println(err)
	}
	log.Printf("JSLD: Lat: %s  Long: %s", lat, long)

	// log.Println(framing.DataSpatial(jlda))
	lat, long, err = frameparser.DataLatLong(framing.DataSpatial(jlda))
	if err != nil {
		log.Println(err)
	}
	log.Printf("JSLD: Lat: %s  Long: %s", lat, long)

}

// test json-ld  NOTE:  datacite not really needed in this context
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
	}`

const jlda = `{
	"@context": {
		"@vocab": "http://schema.org/",
		"datacite": "http://purl.org/spar/datacite/"
	},
		"@type": "Dataset",
		"name": "Removal of organic carbon by natural bacterioplankton communities as a function of pCO2 from laboratory experiments between 2012 and 2016",
		"spatialCoverage": {
			"@type": "Place",
			"geo": [
					{
						"@type": "GeoCoordinates",
						"http://schema.org/latitude": 14.0521333333333,
						"http://schema.org/longitude": 108.0014
					},
					{
						"@type": "GeoCoordinates",
						"http://schema.org/latitude": 14.0521333333333,
						"http://schema.org/longitude": 108.0014
					},
					{
						"@type": "GeoCoordinates",
						"latitude": 39.3280,
						"longitude": 120.1633
					}
		      ]
		}
	}`

const jld2 = `{
	"@context": {
	 "@vocab": "http://schema.org/",
	 "Abstract": "http://opencoredata.org/voc/csdco/v1/abstract",
	 "Discipline": "http://opencoredata.org/voc/csdco/v1/discipline",
	 "Expedition": "http://opencoredata.org/voc/csdco/v1/expedition",
	 "Funding": "http://opencoredata.org/voc/csdco/v1/funding",
	 "Investigators": "http://opencoredata.org/voc/csdco/v1/investigators",
	 "Lab": "http://opencoredata.org/voc/csdco/v1/lab",
	 "Linktitle": "http://opencoredata.org/voc/csdco/v1/linktitle",
	 "Linkurl": "http://opencoredata.org/voc/csdco/v1/linkurl",
	 "Outreach": "http://opencoredata.org/voc/csdco/v1/outreach",
	 "Repository": "http://opencoredata.org/voc/csdco/v1/repository",
	 "Startdate": "http://opencoredata.org/voc/csdco/v1/startdate",
	 "Status": "http://opencoredata.org/voc/csdco/v1/status",
	 "Technique": "http://opencoredata.org/voc/csdco/v1/technique",
	 "csdco": "http://opencoredata.org/voc/csdco/v1/",
	 "re3data": "http://example.org/re3data/0.1/"
	},
	"@id": "http://opencoredata.org/id/do/BHM8",
	"@type": "ResearchProject",
	"csdco:abstract": "",
	"csdco:discipline": "Paleorecords",
	"csdco:expedition": "BHM8",
	"csdco:funding": "Indiana University, Visiting Graduate Student Award",
	"csdco:investigators": "Kelsey Doiron, Arndt Schimmelman, Nguyen Van Huong",
	"csdco:lab": "LacCore",
	"csdco:linktitle": "",
	"csdco:linkurl": "",
	"csdco:outreach": "",
	"csdco:repository": "LacCore",
	"csdco:startdate": "0001-01-01T00:00:00Z",
	"csdco:status": "Lab/ICD",
	"csdco:technique": "Short Coring",
	"description": "",
	"location": {
	 "@type": "Place",
	 "geo": [
	  {
	   "@type": "GeoCoordinates",
	   "http://schema.org/latitude": 14.0521333333333,
	   "http://schema.org/longitude": 108.0014
	  },
	  {
	   "@type": "GeoCoordinates",
	   "http://schema.org/latitude": 14.0521333333333,
	   "http://schema.org/longitude": 108.0014
	  },
	  {
	   "@type": "GeoCoordinates",
	   "http://schema.org/latitude": 14.0521333333333,
	   "http://schema.org/longitude": 108.0014
	  },
	  {
	   "@type": "GeoCoordinates",
	   "http://schema.org/latitude": 14.0521333333333,
	   "http://schema.org/longitude": 108.0014
	  },
	  {
	   "@type": "GeoCoordinates",
	   "http://schema.org/latitude": 14.0521333333333,
	   "http://schema.org/longitude": 108.0014
	  },
	  {
	   "@type": "GeoCoordinates",
	   "http://schema.org/latitude": 14.0521333333333,
	   "http://schema.org/longitude": 108.0014
	  },
	  {
	   "@type": "GeoCoordinates",
	   "http://schema.org/latitude": 14.0521,
	   "http://schema.org/longitude": 108.0013
	  },
	  {
	   "@type": "GeoCoordinates",
	   "http://schema.org/latitude": 14.0521,
	   "http://schema.org/longitude": 108.0013
	  },
	  {
	   "@type": "GeoCoordinates",
	   "http://schema.org/latitude": 14.05211667,
	   "http://schema.org/longitude": 108.00155
	  },
	  {
	   "@type": "GeoCoordinates",
	   "http://schema.org/latitude": 14.05205,
	   "http://schema.org/longitude": 108.0015167
	  },
	  {
	   "@type": "GeoCoordinates",
	   "http://schema.org/latitude": 14.05216667,
	   "http://schema.org/longitude": 108.0014167
	  },
	  {
	   "@type": "GeoCoordinates",
	   "http://schema.org/latitude": 14.05208333,
	   "http://schema.org/longitude": 108.0013333
	  },
	  {
	   "@type": "GeoCoordinates",
	   "http://schema.org/latitude": 14.0521,
	   "http://schema.org/longitude": 108.0135
	  }
	 ]
	},
	"name": "",
	"url": "http://opencoredata.org/id/do/BHM8"
   }`
