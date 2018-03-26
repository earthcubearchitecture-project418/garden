package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/kazarena/json-gold/ld"
)

// DataSetMD a struct to hold metadata about a data set
type DataSetMD struct {
	ID          string
	URL         string
	Description string
	Keywords    string
	Name        string
	ContentURL  string
}

// New main only needs the structs..
// marshals to the struct and then allow us to access each item..
// based on that we can then build a JSON-LD struct, populate and serialize it.
func main() {
	xmlFile, err := os.Open("./sampleData/datacitek4.xml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer xmlFile.Close()
	byteValue, _ := ioutil.ReadAll(xmlFile)

	fmt.Println(xmlToSchemaorg(byteValue))
}

func xmlToSchemaorg(byteValue []byte) string {
	var ds DataSetMD
	var q ChiResource
	xml.Unmarshal(byteValue, &q)

	// Todo  strip all strings to remove returns and other anomalies that can come from XML
	ds.ID = q.ChiIdentifier.Text
	ds.URL = q.ChiAlternateIdentifiers.ChiAlternateIdentifier.Text
	ds.Description = q.ChiDescriptions.ChiDescription.Text
	ds.Name = q.ChiTitles.ChiTitle[0].Text // title is an array, so get the first one.
	ds.Keywords = q.ChiSubjects.ChiSubject.Text
	ds.ContentURL = q.ChiAlternateIdentifiers.ChiAlternateIdentifier.Text

	jsonld, _ := dsetBuilder(ds)

	return string(jsonld)
}

func dsetBuilder(dm DataSetMD) ([]byte, error) {
	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")

	doc := map[string]interface{}{
		"@type": "Dataset",
		"@id":   dm.ID,
		"http://schema.org/url":         dm.URL,
		"http://schema.org/description": dm.Description,
		"http://schema.org/keywords":    dm.Keywords,
		"http://schema.org/name":        dm.Name,
		"http://schema.org/distribution": map[string]interface{}{
			"@type": "DataDownload",
			"http://schema.org/contentUrl": dm.ContentURL,
		},
	}

	context := map[string]interface{}{
		"@context": map[string]interface{}{
			"@vocab":  "http://schema.org/",
			"re3data": "http://example.org/re3data/0.1/",
		},
	}

	compactedDoc, err := proc.Compact(doc, context, options)
	if err != nil {
		fmt.Println("Error when compacting", err)
	}

	return json.MarshalIndent(compactedDoc, "", " ")
}

///////////////////////////
/// structs
///////////////////////////

type ChiChidleyRoot314159 struct {
	ChiResource *ChiResource `xml:"http://datacite.org/schema/kernel-4 resource,omitempty" json:"resource,omitempty"`
}

type ChiAffiliation struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://datacite.org/schema/kernel-4 affiliation,omitempty" json:"affiliation,omitempty"`
}

type ChiAlternateIdentifier struct {
	AttrAlternateIdentifierType string   `xml:" alternateIdentifierType,attr"  json:",omitempty"`
	Text                        string   `xml:",chardata" json:",omitempty"`
	XMLName                     xml.Name `xml:"http://datacite.org/schema/kernel-4 alternateIdentifier,omitempty" json:"alternateIdentifier,omitempty"`
}

type ChiAlternateIdentifiers struct {
	ChiAlternateIdentifier *ChiAlternateIdentifier `xml:"http://datacite.org/schema/kernel-4 alternateIdentifier,omitempty" json:"alternateIdentifier,omitempty"`
	XMLName                xml.Name                `xml:"http://datacite.org/schema/kernel-4 alternateIdentifiers,omitempty" json:"alternateIdentifiers,omitempty"`
}

type ChiContributor struct {
	AttrContributorType string              `xml:" contributorType,attr"  json:",omitempty"`
	ChiAffiliation      *ChiAffiliation     `xml:"http://datacite.org/schema/kernel-4 affiliation,omitempty" json:"affiliation,omitempty"`
	ChiContributorName  *ChiContributorName `xml:"http://datacite.org/schema/kernel-4 contributorName,omitempty" json:"contributorName,omitempty"`
	ChiNameIdentifier   *ChiNameIdentifier  `xml:"http://datacite.org/schema/kernel-4 nameIdentifier,omitempty" json:"nameIdentifier,omitempty"`
	XMLName             xml.Name            `xml:"http://datacite.org/schema/kernel-4 contributor,omitempty" json:"contributor,omitempty"`
}

type ChiContributorName struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://datacite.org/schema/kernel-4 contributorName,omitempty" json:"contributorName,omitempty"`
}

type ChiContributors struct {
	ChiContributor *ChiContributor `xml:"http://datacite.org/schema/kernel-4 contributor,omitempty" json:"contributor,omitempty"`
	XMLName        xml.Name        `xml:"http://datacite.org/schema/kernel-4 contributors,omitempty" json:"contributors,omitempty"`
}

type ChiCreator struct {
	ChiAffiliation    *ChiAffiliation    `xml:"http://datacite.org/schema/kernel-4 affiliation,omitempty" json:"affiliation,omitempty"`
	ChiCreatorName    *ChiCreatorName    `xml:"http://datacite.org/schema/kernel-4 creatorName,omitempty" json:"creatorName,omitempty"`
	ChiFamilyName     *ChiFamilyName     `xml:"http://datacite.org/schema/kernel-4 familyName,omitempty" json:"familyName,omitempty"`
	ChiGivenName      *ChiGivenName      `xml:"http://datacite.org/schema/kernel-4 givenName,omitempty" json:"givenName,omitempty"`
	ChiNameIdentifier *ChiNameIdentifier `xml:"http://datacite.org/schema/kernel-4 nameIdentifier,omitempty" json:"nameIdentifier,omitempty"`
	XMLName           xml.Name           `xml:"http://datacite.org/schema/kernel-4 creator,omitempty" json:"creator,omitempty"`
}

type ChiCreatorName struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://datacite.org/schema/kernel-4 creatorName,omitempty" json:"creatorName,omitempty"`
}

type ChiCreators struct {
	ChiCreator *ChiCreator `xml:"http://datacite.org/schema/kernel-4 creator,omitempty" json:"creator,omitempty"`
	XMLName    xml.Name    `xml:"http://datacite.org/schema/kernel-4 creators,omitempty" json:"creators,omitempty"`
}

type ChiDate struct {
	AttrDateType string   `xml:" dateType,attr"  json:",omitempty"`
	Text         string   `xml:",chardata" json:",omitempty"`
	XMLName      xml.Name `xml:"http://datacite.org/schema/kernel-4 date,omitempty" json:"date,omitempty"`
}

type ChiDates struct {
	ChiDate *ChiDate `xml:"http://datacite.org/schema/kernel-4 date,omitempty" json:"date,omitempty"`
	XMLName xml.Name `xml:"http://datacite.org/schema/kernel-4 dates,omitempty" json:"dates,omitempty"`
}

type ChiDescription struct {
	AttrDescriptionType string   `xml:" descriptionType,attr"  json:",omitempty"`
	Text                string   `xml:",chardata" json:",omitempty"`
	XMLName             xml.Name `xml:"http://datacite.org/schema/kernel-4 description,omitempty" json:"description,omitempty"`
}

type ChiDescriptions struct {
	ChiDescription *ChiDescription `xml:"http://datacite.org/schema/kernel-4 description,omitempty" json:"description,omitempty"`
	XMLName        xml.Name        `xml:"http://datacite.org/schema/kernel-4 descriptions,omitempty" json:"descriptions,omitempty"`
}

type ChiEastBoundLongitude struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://datacite.org/schema/kernel-4 eastBoundLongitude,omitempty" json:"eastBoundLongitude,omitempty"`
}

type ChiFamilyName struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://datacite.org/schema/kernel-4 familyName,omitempty" json:"familyName,omitempty"`
}

type ChiFormat struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://datacite.org/schema/kernel-4 format,omitempty" json:"format,omitempty"`
}

type ChiFormats struct {
	ChiFormat *ChiFormat `xml:"http://datacite.org/schema/kernel-4 format,omitempty" json:"format,omitempty"`
	XMLName   xml.Name   `xml:"http://datacite.org/schema/kernel-4 formats,omitempty" json:"formats,omitempty"`
}

type ChiGeoLocation struct {
	ChiGeoLocationBox   *ChiGeoLocationBox   `xml:"http://datacite.org/schema/kernel-4 geoLocationBox,omitempty" json:"geoLocationBox,omitempty"`
	ChiGeoLocationPlace *ChiGeoLocationPlace `xml:"http://datacite.org/schema/kernel-4 geoLocationPlace,omitempty" json:"geoLocationPlace,omitempty"`
	ChiGeoLocationPoint *ChiGeoLocationPoint `xml:"http://datacite.org/schema/kernel-4 geoLocationPoint,omitempty" json:"geoLocationPoint,omitempty"`
	XMLName             xml.Name             `xml:"http://datacite.org/schema/kernel-4 geoLocation,omitempty" json:"geoLocation,omitempty"`
}

type ChiGeoLocationBox struct {
	ChiEastBoundLongitude *ChiEastBoundLongitude `xml:"http://datacite.org/schema/kernel-4 eastBoundLongitude,omitempty" json:"eastBoundLongitude,omitempty"`
	ChiNorthBoundLatitude *ChiNorthBoundLatitude `xml:"http://datacite.org/schema/kernel-4 northBoundLatitude,omitempty" json:"northBoundLatitude,omitempty"`
	ChiSouthBoundLatitude *ChiSouthBoundLatitude `xml:"http://datacite.org/schema/kernel-4 southBoundLatitude,omitempty" json:"southBoundLatitude,omitempty"`
	ChiWestBoundLongitude *ChiWestBoundLongitude `xml:"http://datacite.org/schema/kernel-4 westBoundLongitude,omitempty" json:"westBoundLongitude,omitempty"`
	XMLName               xml.Name               `xml:"http://datacite.org/schema/kernel-4 geoLocationBox,omitempty" json:"geoLocationBox,omitempty"`
}

type ChiGeoLocationPlace struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://datacite.org/schema/kernel-4 geoLocationPlace,omitempty" json:"geoLocationPlace,omitempty"`
}

type ChiGeoLocationPoint struct {
	ChiPointLatitude  *ChiPointLatitude  `xml:"http://datacite.org/schema/kernel-4 pointLatitude,omitempty" json:"pointLatitude,omitempty"`
	ChiPointLongitude *ChiPointLongitude `xml:"http://datacite.org/schema/kernel-4 pointLongitude,omitempty" json:"pointLongitude,omitempty"`
	XMLName           xml.Name           `xml:"http://datacite.org/schema/kernel-4 geoLocationPoint,omitempty" json:"geoLocationPoint,omitempty"`
}

type ChiGeoLocations struct {
	ChiGeoLocation *ChiGeoLocation `xml:"http://datacite.org/schema/kernel-4 geoLocation,omitempty" json:"geoLocation,omitempty"`
	XMLName        xml.Name        `xml:"http://datacite.org/schema/kernel-4 geoLocations,omitempty" json:"geoLocations,omitempty"`
}

type ChiGivenName struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://datacite.org/schema/kernel-4 givenName,omitempty" json:"givenName,omitempty"`
}

type ChiIdentifier struct {
	AttrIdentifierType string   `xml:" identifierType,attr"  json:",omitempty"`
	Text               string   `xml:",chardata" json:",omitempty"`
	XMLName            xml.Name `xml:"http://datacite.org/schema/kernel-4 identifier,omitempty" json:"identifier,omitempty"`
}

type ChiLanguage struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://datacite.org/schema/kernel-4 language,omitempty" json:"language,omitempty"`
}

type ChiNameIdentifier struct {
	AttrNameIdentifierScheme string   `xml:" nameIdentifierScheme,attr"  json:",omitempty"`
	AttrSchemeURI            string   `xml:" schemeURI,attr"  json:",omitempty"`
	Text                     string   `xml:",chardata" json:",omitempty"`
	XMLName                  xml.Name `xml:"http://datacite.org/schema/kernel-4 nameIdentifier,omitempty" json:"nameIdentifier,omitempty"`
}

type ChiNorthBoundLatitude struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://datacite.org/schema/kernel-4 northBoundLatitude,omitempty" json:"northBoundLatitude,omitempty"`
}

type ChiPointLatitude struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://datacite.org/schema/kernel-4 pointLatitude,omitempty" json:"pointLatitude,omitempty"`
}

type ChiPointLongitude struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://datacite.org/schema/kernel-4 pointLongitude,omitempty" json:"pointLongitude,omitempty"`
}

type ChiPublicationYear struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://datacite.org/schema/kernel-4 publicationYear,omitempty" json:"publicationYear,omitempty"`
}

type ChiPublisher struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://datacite.org/schema/kernel-4 publisher,omitempty" json:"publisher,omitempty"`
}

type ChiRelatedIdentifier struct {
	AttrRelatedIdentifierType string   `xml:" relatedIdentifierType,attr"  json:",omitempty"`
	AttrRelatedMetadataScheme string   `xml:" relatedMetadataScheme,attr"  json:",omitempty"`
	AttrRelationType          string   `xml:" relationType,attr"  json:",omitempty"`
	AttrSchemeURI             string   `xml:" schemeURI,attr"  json:",omitempty"`
	Text                      string   `xml:",chardata" json:",omitempty"`
	XMLName                   xml.Name `xml:"http://datacite.org/schema/kernel-4 relatedIdentifier,omitempty" json:"relatedIdentifier,omitempty"`
}

type ChiRelatedIdentifiers struct {
	ChiRelatedIdentifier []*ChiRelatedIdentifier `xml:"http://datacite.org/schema/kernel-4 relatedIdentifier,omitempty" json:"relatedIdentifier,omitempty"`
	XMLName              xml.Name                `xml:"http://datacite.org/schema/kernel-4 relatedIdentifiers,omitempty" json:"relatedIdentifiers,omitempty"`
}

type ChiResource struct {
	AttrXsiSpaceSchemaLocation string                   `xml:"http://www.w3.org/2001/XMLSchema-instance schemaLocation,attr"  json:",omitempty"`
	AttrXmlns                  string                   `xml:" xmlns,attr"  json:",omitempty"`
	AttrXmlnsXsi               string                   `xml:"xmlns xsi,attr"  json:",omitempty"`
	ChiAlternateIdentifiers    *ChiAlternateIdentifiers `xml:"http://datacite.org/schema/kernel-4 alternateIdentifiers,omitempty" json:"alternateIdentifiers,omitempty"`
	ChiContributors            *ChiContributors         `xml:"http://datacite.org/schema/kernel-4 contributors,omitempty" json:"contributors,omitempty"`
	ChiCreators                *ChiCreators             `xml:"http://datacite.org/schema/kernel-4 creators,omitempty" json:"creators,omitempty"`
	ChiDates                   *ChiDates                `xml:"http://datacite.org/schema/kernel-4 dates,omitempty" json:"dates,omitempty"`
	ChiDescriptions            *ChiDescriptions         `xml:"http://datacite.org/schema/kernel-4 descriptions,omitempty" json:"descriptions,omitempty"`
	ChiFormats                 *ChiFormats              `xml:"http://datacite.org/schema/kernel-4 formats,omitempty" json:"formats,omitempty"`
	ChiGeoLocations            *ChiGeoLocations         `xml:"http://datacite.org/schema/kernel-4 geoLocations,omitempty" json:"geoLocations,omitempty"`
	ChiIdentifier              *ChiIdentifier           `xml:"http://datacite.org/schema/kernel-4 identifier,omitempty" json:"identifier,omitempty"`
	ChiLanguage                *ChiLanguage             `xml:"http://datacite.org/schema/kernel-4 language,omitempty" json:"language,omitempty"`
	ChiPublicationYear         *ChiPublicationYear      `xml:"http://datacite.org/schema/kernel-4 publicationYear,omitempty" json:"publicationYear,omitempty"`
	ChiPublisher               *ChiPublisher            `xml:"http://datacite.org/schema/kernel-4 publisher,omitempty" json:"publisher,omitempty"`
	ChiRelatedIdentifiers      *ChiRelatedIdentifiers   `xml:"http://datacite.org/schema/kernel-4 relatedIdentifiers,omitempty" json:"relatedIdentifiers,omitempty"`
	ChiResourceType            *ChiResourceType         `xml:"http://datacite.org/schema/kernel-4 resourceType,omitempty" json:"resourceType,omitempty"`
	ChiRightsList              *ChiRightsList           `xml:"http://datacite.org/schema/kernel-4 rightsList,omitempty" json:"rightsList,omitempty"`
	ChiSizes                   *ChiSizes                `xml:"http://datacite.org/schema/kernel-4 sizes,omitempty" json:"sizes,omitempty"`
	ChiSubjects                *ChiSubjects             `xml:"http://datacite.org/schema/kernel-4 subjects,omitempty" json:"subjects,omitempty"`
	ChiTitles                  *ChiTitles               `xml:"http://datacite.org/schema/kernel-4 titles,omitempty" json:"titles,omitempty"`
	ChiVersion                 *ChiVersion              `xml:"http://datacite.org/schema/kernel-4 version,omitempty" json:"version,omitempty"`
	XMLName                    xml.Name                 `xml:"http://datacite.org/schema/kernel-4 resource,omitempty" json:"resource,omitempty"`
}

type ChiResourceType struct {
	AttrResourceTypeGeneral string   `xml:" resourceTypeGeneral,attr"  json:",omitempty"`
	Text                    string   `xml:",chardata" json:",omitempty"`
	XMLName                 xml.Name `xml:"http://datacite.org/schema/kernel-4 resourceType,omitempty" json:"resourceType,omitempty"`
}

type ChiRights struct {
	AttrRightsURI string   `xml:" rightsURI,attr"  json:",omitempty"`
	Text          string   `xml:",chardata" json:",omitempty"`
	XMLName       xml.Name `xml:"http://datacite.org/schema/kernel-4 rights,omitempty" json:"rights,omitempty"`
}

type ChiRightsList struct {
	ChiRights *ChiRights `xml:"http://datacite.org/schema/kernel-4 rights,omitempty" json:"rights,omitempty"`
	XMLName   xml.Name   `xml:"http://datacite.org/schema/kernel-4 rightsList,omitempty" json:"rightsList,omitempty"`
}

type ChiSize struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://datacite.org/schema/kernel-4 size,omitempty" json:"size,omitempty"`
}

type ChiSizes struct {
	ChiSize *ChiSize `xml:"http://datacite.org/schema/kernel-4 size,omitempty" json:"size,omitempty"`
	XMLName xml.Name `xml:"http://datacite.org/schema/kernel-4 sizes,omitempty" json:"sizes,omitempty"`
}

type ChiSouthBoundLatitude struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://datacite.org/schema/kernel-4 southBoundLatitude,omitempty" json:"southBoundLatitude,omitempty"`
}

type ChiSubject struct {
	AttrSchemeURI     string   `xml:" schemeURI,attr"  json:",omitempty"`
	AttrSubjectScheme string   `xml:" subjectScheme,attr"  json:",omitempty"`
	Text              string   `xml:",chardata" json:",omitempty"`
	XMLName           xml.Name `xml:"http://datacite.org/schema/kernel-4 subject,omitempty" json:"subject,omitempty"`
}

type ChiSubjects struct {
	ChiSubject *ChiSubject `xml:"http://datacite.org/schema/kernel-4 subject,omitempty" json:"subject,omitempty"`
	XMLName    xml.Name    `xml:"http://datacite.org/schema/kernel-4 subjects,omitempty" json:"subjects,omitempty"`
}

type ChiTitle struct {
	AttrTitleType string   `xml:" titleType,attr"  json:",omitempty"`
	Text          string   `xml:",chardata" json:",omitempty"`
	XMLName       xml.Name `xml:"http://datacite.org/schema/kernel-4 title,omitempty" json:"title,omitempty"`
}

type ChiTitles struct {
	ChiTitle []*ChiTitle `xml:"http://datacite.org/schema/kernel-4 title,omitempty" json:"title,omitempty"`
	XMLName  xml.Name    `xml:"http://datacite.org/schema/kernel-4 titles,omitempty" json:"titles,omitempty"`
}

type ChiVersion struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://datacite.org/schema/kernel-4 version,omitempty" json:"version,omitempty"`
}

type ChiWestBoundLongitude struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://datacite.org/schema/kernel-4 westBoundLongitude,omitempty" json:"westBoundLongitude,omitempty"`
}

///////////////////////////
