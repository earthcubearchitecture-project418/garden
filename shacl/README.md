## SHACL Notes

### Setup
Install the TopQuadrant SHACL (Java) library

```
export SHACLROOT=/home/fils/Semantic/SHACL/shacl-1.0.0/bin
export PATH=$SHACLROOT:$PATH 
```
The following should fail (no citation entry, a recommended item)
```
shaclvalidate.sh -datafile ocddataset.ttl -shapesfile recomendShape.ttl
```

### Network based testing
The following examples use [httpie](https://github.com/jakubroztocil/httpie/) but you could 
use curl or something visual like the nice [Postman](https://www.getpostman.com/) tool.  


```
fils@xps:~/Semantic/SHACL$ http --form POST http://localhost:6789/api/beta/shacl/eval  datagraph@ocddataset.ttl  shapesgraph@recomendShape.ttl
HTTP/1.1 200 OK
Content-Length: 1150
Content-Type: text/plain; charset=utf-8
Date: Thu, 28 Jun 2018 15:08:06 GMT

@prefix rdf:   <http://www.w3.org/1999/02/22-rdf-syntax-ns#> .
@prefix owl:   <http://www.w3.org/2002/07/owl#> .
@prefix xsd:   <http://www.w3.org/2001/XMLSchema#> .
@prefix rdfs:  <http://www.w3.org/2000/01/rdf-schema#> .

[ a       <http://www.w3.org/ns/shacl#ValidationReport> ;
  <http://www.w3.org/ns/shacl#conforms>
          false ;
  <http://www.w3.org/ns/shacl#result>
          [ a       <http://www.w3.org/ns/shacl#ValidationResult> ;
            <http://www.w3.org/ns/shacl#focusNode>
                    <http://opencoredata.org/id/dataset/bcd15975-680c-47db-a062-ac0bb6e66816> ;
            <http://www.w3.org/ns/shacl#resultMessage>
                    "Less than 1 values" ;
            <http://www.w3.org/ns/shacl#resultPath>
                    <http://schema.org/citation> ;
            <http://www.w3.org/ns/shacl#resultSeverity>
                    <http://www.w3.org/ns/shacl#Violation> ;
            <http://www.w3.org/ns/shacl#sourceConstraintComponent>
                    <http://www.w3.org/ns/shacl#MinCountConstraintComponent> ;
            <http://www.w3.org/ns/shacl#sourceShape>
                    []
          ]
] .
```




```
fils@xps:~/Semantic/SHACL$ http --form POST http://localhost:6789/api/beta/shacl/eval  datagraph@ocddataset.ttl  shapesgraph@requiredShape.ttl
HTTP/1.1 200 OK
Content-Length: 341
Content-Type: text/plain; charset=utf-8
Date: Thu, 28 Jun 2018 16:41:21 GMT

@prefix rdf:   <http://www.w3.org/1999/02/22-rdf-syntax-ns#> .
@prefix owl:   <http://www.w3.org/2002/07/owl#> .
@prefix xsd:   <http://www.w3.org/2001/XMLSchema#> .
@prefix rdfs:  <http://www.w3.org/2000/01/rdf-schema#> .

[ a       <http://www.w3.org/ns/shacl#ValidationReport> ;
  <http://www.w3.org/ns/shacl#conforms>
          true
] .
```




### Refs
Use https://developers.google.com/search/docs/data-types/dataset to define the shapes for 
required and recommended


http -f POST localhost:7000 datag@./ocddataset.ttl shapeg@./dataSetRecomendShape.ttl dataref=datarefurl  shaperef=shaperefurl
