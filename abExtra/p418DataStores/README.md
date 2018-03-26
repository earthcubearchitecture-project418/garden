### About
This is a general repo that is mostly a place
to document the various backend data storage approaches we use
to support P418 work

All the approaches have the criteria to work with docker.  So
this material will also provide information on how that is being done.

#### triplestore (blazegraph)
Docker image: https://hub.docker.com/r/nawer/blazegraph/

compose section
```
    db:
        image: nawer/blazegraph
        environment:
            JAVA_XMS: 512m
            JAVA_XMX: 1g
        volumes:
            - /var/blazegraph:/var/lib/blazegraph
        ports:
            - "9999:9999"
```
Not used: ```   - ./docker-entrypoint-initdb.d:/docker-entrypoint-initdb.d  ```

Local volume journal
```
$ docker run -p 9999:9999 --name p418triplestore -v /Users/dfils/Data/P418/dataVolumes/triplestore:/var/lib/blazegraph -d nawer/blazegraph:latest
```

If /my/custom/override.xml is the path and name of your custom configuration file, you can start your blazegraph container like this :

```
$ docker run --name some-blazegraph -v /my/custom/override.xml:/etc/blazegraph/override.xml -d nawer/blazegraph:tag
```

Example SAPRQL

```
PREFIX schemaorg: <http://schema.org/> 
SELECT DISTINCT ?repository ?name ?url ?logo ?description ?contact_name ?contact_email ?contact_url ?contact_role 
WHERE { 
    { ?repository schemaorg:url <http://wwww.bco-dmo.org> . } 
    UNION 
    { ?repository <http://schema.org/url> "http://wwww.bco-dmo.org" . } 
    ?repository rdf:type <http://schema.org/Organization> . 
    ?repository schemaorg:name ?name . 
    ?repository schemaorg:url ?url . 
    OPTIONAL { ?repository schemaorg:description ?description . } 
    OPTIONAL { ?repository schemaorg:logo [ schemaorg:url ?logo ] . } 
    OPTIONAL { ?repository schemaorg:contactPoint ?contact . 
    ?contact schemaorg:name ?contact_name . 
    ?contact schemaorg:email ?contact_email . 
    ?contact schemaorg:contactType ?contact_role .
     ?contact schemaorg:url ?contact_url . 
     } 
     } 
     LIMIT 1
     ```



#### Spatial (tile38)

```
docker run -d -p 9851:9851 -v /root/Crawler/spatialindex:/data -t tile38/tile38
```


#### S3 (Minio)


#### Text index (bleve)
Not a data store but rather a library and results in an index.  It is
noted here since we also store a copy of the index document in the underlying
KV store of blevel.  

