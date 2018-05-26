### blankNodeApproaches

I am reviewing some packages to see if they would address the workflow need of going
from JSON-LD (which serializes to NT with blank nodes) into a common graph.  This the
indexer can build one graph from a given crawl that can be immediately moved into use 
by the services that use the triples.  

* https://github.com/Callidon/joseki  Loads, but not sure how to test via query 
* https://github.com/deiu/rdf2go  Has an issue with blank nodes..   https://github.com/deiu/rdf2go/issues/3 
* https://github.com/dgraph-io/dgraph  https://dgraph.io/ too much?  out of scope for this work?  This might be a good server..  but not fit for purpose on the crawler side.  
* https://github.com/knakk/rdf  solid..  in use elsewhere, not tested for this work.  It's really just a triple library, not a graph library.  As such
it doesn't address reconciling duplicate blank nodes. 



* https://github.com/kazarena/json-gold  solid..  in use for JSON-LD work
* https://github.com/wallix/triplestore nice!  but doesn't do blank nodes yet

#### Format notes
I may need to process the blank nodes myself.  As such an approach like
```
_:b[UUID]
```
would be used.  The UUID just needs to be unique and the structure is not
that important for our work.  So I really just want the shortest possible 
unique string or just move a counter forward each time I need a new blank
node in a given "graph".   

I would need to read the JSON-LD, locate all blank nodes, then get a new 
globaly unique bnode ID for each existing bnode ID in the "current" graph. 
Change them and then add to the global graph.  The logic is not that complex.