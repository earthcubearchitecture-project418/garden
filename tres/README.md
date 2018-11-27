# Time tester

##  Steps

* set up simple RESTful service to return second counts
* expose it following OpenSearch or GeoService URI patterns
* make a landing page for the service
* put in the schema.org with the API description and Time patterns for search
* address issue of "discoverable" vs "citable" units of time
* How to address exposing data sets in quanta of time that address the above

## Notes

This work is designed to try and provide a means to expose data mediated by services to 
expose similar attributes exposed by static data sets via traditional *landing pages*. 

To address this we are using terms like:

### Services Landing Page:  

Which is like a traditional data set landing page but exposes
information about services.  Not, while we may leverage things like schema:SearchAction the
intent is not to replace things like Swagger, Hydra or OpenAPI.   These later service description 
methods are far richer than what can be exposed in schema.org represented by JSON-LD.  

We will also use OpenSearch and GeoWS patterns (https://www.earthcube.org/document/2015/geows-all-hands-poster)
as URL template patterns.

### Virtual Data Set Landing Page:  

Extending the above concept is should be expressed the goal
is to expose the underlaying data.  As such, using methods to describe the time span of the data and 
using that in the discovery process is important.  This means that we need to leverage existing concepts
in spatial description such as:

* OWL Time (https://www.w3.org/TR/owl-time/)
* Geologic Time (https://github.com/GeoscienceAustralia/geosciml.org)
* Temporal "gazateer", which is to say terms like "Hurricane Ira", "Fukashima", etc.  

To get to this concept we need to be able to describe a landing page with concepts of time that
facilitate discovery.  This means concepts like

``` 
Data X for Oct 2011 
```

are desrired.  However, we will need to focus on data request following established time encoding (ISO 8601 and or RFC 3339 data formats) 
(ref: http://en.wikipedia.org/wiki/ISO_8601 http://tools.ietf.org/html/rfc3339 https://stackoverflow.com/questions/522251/whats-the-difference-between-iso-8601-and-rfc-3339-date-formats ).

So the above might become something more like

``` 
RFC3339 version here 
```

Initially the work will focus on explicate notation for the date, later version 
might regex the search for time.  Inferring 
native langage time bounding is outside the scope of this project.


## Dates

* Review BoltDB RFC3339 date order searching (https://github.com/boltdb/bolt).  
* Review xsd:date can also be used with SPARQL for searching.
* Note these will not address shallow or deep time.
* Note xsd:date does not have a resolution below day

## Calls

API calls will be required to use RFC3339 formated dates initially. This may be relaxed
later as we put in functions to convery other formats.  

```
http localhost:6789/api/v1/time/test?s=2011-06-05T14:10:43.678Z&e=2013-06-05T14:10:43.678Z
```

## SPARQL

Using SPARQL on the GSA Geologic timescale can provide things such as stages for 
given ages (Ma).  Exaples include:

```SPARQL
prefix gts: <http://resource.geosciml.org/ontology/timescale/gts#>
PREFIX rdfs: <http://www.w3.org/2000/01/rdf-schema#>
PREFIX time: <http://www.w3.org/2006/time#>
PREFIX thors: <http://resource.geosciml.org/ontology/timescale/thors#>
PREFIX tm: <http://def.seegrid.csiro.au/isotc211/iso19108/2002/temporal#>
SELECT *
WHERE {
              ?era time:hasBeginning/time:inTemporalPosition/time:numericPosition ?begin .
              ?era time:hasEnd/time:inTemporalPosition/time:numericPosition ?end .
              ?era rdfs:label ?label
               BIND ( "200"^^xsd:decimal AS ?targetAge )
               FILTER ( ?targetAge > xsd:decimal(?end) )
               FILTER ( ?targetAge < xsd:decimal(?begin) )
}
```


## Organic search considerations

How we work with time in terms of organic search is a consideration.   Initially we will likely 
deal only with "day" being the smallest discovery unit.  Obviously this only scopes to modern 
time data resources.   For deep time we will use geologic stage names and ranges in the Ma units.
We will also use Ka units for shallow time resources.   

Initially we will use something like 

```
siesmic data for california datestart:2011-06-05 dateend:2013-06-05 
```

```
magnetic data for geostage:Jurassic
```

In the case of the geotime search we can expand the search to incude sub-stage names 
from Jurassic such as Hettangian Age.  Note, we would not likely move up, in this case
to Era or Eon, since the search defined a scope size.   Only items of smaller scope
would seem valid extensions to the search.  For cases where a range of geologic numeric time 
is provided, such as;


```
magnetic data for geoagema:175-200
```

we use the range 175Ma an 200Ma to both search on numeric ranges exposed by a resource and 
also to arrive at potential stage name values. Note that Eon and Era are likely far to large
to be of value, though we will need to resolve how to deal with these so as not to incorrectly
exclude them.   We will need to engage the scientific community to resolve issues of scope 
for searches.  

## Defining Ma in JSON-LD

Based on the works of the GeoScience Australia timescale work, a concept for defining 
numeric age in JSON-LD might look like

```JSON
{
    "@id" : "isc:BaseAlbianTime",
    "@type" : "time:TimePosition",
    "hasTRS" : "http://resource.geosciml.org/classifier/cgi/geologicage/ma",
    "numericPosition" : "113.0"
  }, {
    "@id" : "isc:BaseAlbianUncertainty",
    "@type" : "time:Duration",
    "numericDuration" : "1.0",
    "unitType" : "http://www.opengis.net/def/uom/UCUM/0/Ma"
  },
```

There is still much to do, however and example of what a JSON-LD description of a data
resource define by relation to geologic stage names follows.  

```JSON
{
    "@id" : "isc:Bathonian",
    "@type" : [ "skos:Concept", "gts:GeochronologicEra", "time:ProperInterval", "gts:Age" ],
    "label" : {
      "@language" : "en",
      "@value" : "Bathonian Age"
    },
    "inScheme" : "ts:isc2017",
    "notation" : "a1.1.2.2.2.2",
    "prefLabel" : [ {
      "@language" : "es",
      "@value" : "Bathoniense"
    }],
    "hasBeginning" : "isc:BaseBathonian",
    "hasEnd" : "isc:BaseCallovian",
    "intervalDuring" : "isc:MiddleJurassic",
    "intervalIn" : "isc:MiddleJurassic",
    "intervalMeets" : "isc:Callovian",
    "intervalMetBy" : "isc:Bajocian"
}
```
