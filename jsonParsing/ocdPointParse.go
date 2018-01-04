package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("JSON parsing test")

	str := []byte(jsonLD())

	var f interface{}
	err := json.Unmarshal(str, &f)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Key: spatial Value: map[@type:Place geo:map[@type:GeoCoordinates latitude:26.03 longitude:-147.93]]
	z := f.(map[string]interface{})
	urlval := z["url"]
	spatial := z["spatial"]
	geocast := spatial.(map[string]interface{})
	geo := geocast["geo"]
	gccast := geo.(map[string]interface{})
	latval := gccast["latitude"]
	lngval := gccast["longitude"]

	fmt.Println(geo)
	fmt.Println(urlval)
	fmt.Println(latval)
	fmt.Println(lngval)

}

func jsonLDNOTDATASET() string {

	jsonldtext := `{
 "@context": {
  "@vocab": "http://schema.org/",
  "re3data": "http://example.org/re3data/0.1/"
 },
 "@id": "http://opencoredata.org/id/dataset/045deec9-94b2-445a-8fd2-43dbe90841fb",
 "@type": "Dataset",
 "description": "Janus Vcd Image for ocean drilling expedition 199 site 1215 hole A",
 "distribution": {
  "@type": "DataDownload",
  "contentUrl": "http://opencoredata.org/api/v1/documents/download/199_1215A_JanusVcdImage_JcAruSDk.csv"
 },
 "keywords": "Leg Site Hole Core Core_type Section_number Section_type Top_cm Depth_mbsf Page_id Url Janus Vcd Image DSDP, OPD, IODP, JanusVcdImage",
 "name": "199_1215A_JanusVcdImage_JcAruSDk.csv",
 "publisher": {
  "@type": "Organization",
  "description": "NSF funded International Ocean Discovery Program operated by JRSO",
  "name": "International Ocean Discovery Program",
  "url": "http://iodp.org"
 },
 "spatial": {
  "@type": "Place",
  "geo": {
   "@type": "GeoCoordinates",
   "latitude": "26.03",
   "longitude": "-147.93"
  }
 },
 "url": "http://opencoredata.org/id/dataset/045deec9-94b2-445a-8fd2-43dbe90841fb",
 "variableMeasured": "Janus Vcd Image"
} `

	return jsonldtext

}

func jsonLD() string {

	jsonldtext := `{
 "@context": {
  "@vocab": "http://schema.org/",
  "re3data": "http://example.org/re3data/0.1/"
 },
 "@id": "http://opencoredata.org/id/dataset/045deec9-94b2-445a-8fd2-43dbe90841fb",
 "@type": "Dataset",
 "description": "Janus Vcd Image for ocean drilling expedition 199 site 1215 hole A",
 "distribution": {
  "@type": "DataDownload",
  "contentUrl": "http://opencoredata.org/api/v1/documents/download/199_1215A_JanusVcdImage_JcAruSDk.csv"
 },
 "keywords": "Leg Site Hole Core Core_type Section_number Section_type Top_cm Depth_mbsf Page_id Url Janus Vcd Image DSDP, OPD, IODP, JanusVcdImage",
 "name": "199_1215A_JanusVcdImage_JcAruSDk.csv",
 "publisher": {
  "@type": "Organization",
  "description": "NSF funded International Ocean Discovery Program operated by JRSO",
  "name": "International Ocean Discovery Program",
  "url": "http://iodp.org"
 },
 "spatial": {
  "@type": "Place",
  "geo": {
   "@type": "GeoCoordinates",
   "latitude": "26.03",
   "longitude": "-147.93"
  }
 },
 "url": "http://opencoredata.org/id/dataset/045deec9-94b2-445a-8fd2-43dbe90841fb",
 "variableMeasured": "Janus Vcd Image"
} `

	return jsonldtext

}
