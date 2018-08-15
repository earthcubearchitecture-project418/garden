
#### Deprecated 
This code has NOT been used in P418 for some time.  Please reference the Gleaner repository for the currently 
used code base.  


## P418 Crawler

### About
Some work on some crawler code.  The crawling aspect is rather simple (boring even).  It will be sitemap driven
and not really innovative at all.   It should focus mostly on routing responses to various index
functions.  The goal here will be to leverage Go libraries to do all the cool work for us WRT to 
crawling.  

This really isn't even a crawler.  Given the sitemap approach, it is really just an indexer of know white listed
items.  Each item is at this stage expected to be a data landing page.  We will look in the page for a 
JSON-LD package and process it.   We will likely add in a test to see that it is of type schema.org/Dataset.

The JSON-LD is sent to three processors; text, spatial and graph.  Later we will try and add in 
temporal, but will likely be the most problematic (though not sure why I think that now).  


### Running

After building the crawler 

```
go build -o 418crawler
```

you can select from a set of command line flag options

```
418crawler -url=http://foo.org    # index only this domain
```

```
418crawler -csv=./input/indexThese.csv    # link to CSV file with 1 domain per line (really just .txt for now)
```

```
418crawler -sitemap=./input/sitemap.xml    # a local sitemap in XMl format you wish to index
```

You can also add in ``` -indexname=./path/to/index ``` where the bleve index will be recorded.
The JSON is converted to a file called nquads.rdf (no naming option yet) and 


### Deps (for data IO, not compiling..   see vendoring for those)
The only dep for running is a Tile38 geohash server that we talk to via a redis client.  

For example:
```
docker run -d -p 9851:9851 -v ~/Data/OCDDataVolumes/spatialindex/data:/data -t tile38/tile38
```

