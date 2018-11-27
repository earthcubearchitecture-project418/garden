var doc = {
  "http://schema.org/name": "Manu Sporny",
  "http://schema.org/url": {"@id": "http://manu.sporny.org/"},
  "http://schema.org/image": {"@id": "http://manu.sporny.org/images/manu.png"}
};
var context = {
  "name": "http://schema.org/name",
  "homepage": {"@id": "http://schema.org/url", "@type": "@id"},
  "image": {"@id": "http://schema.org/image", "@type": "@id"}
};

// compact a document according to a particular context
// see: http://json-ld.org/spec/latest/json-ld/#compacted-document-form
jsonld.compact(doc, context, function(err, compacted) {
  console.log(JSON.stringify(compacted, null, 2));
  /* Output:
  {
    "@context": {...},
    "name": "Manu Sporny",
    "homepage": "http://manu.sporny.org/",
    "image": "http://manu.sporny.org/images/manu.png"
  }
  */
});

// compact using URLs
jsonld.compact('http://example.org/doc', 'http://example.org/context', ...);

// expand a document, removing its context
// see: http://json-ld.org/spec/latest/json-ld/#expanded-document-form
jsonld.expand(compacted, function(err, expanded) {
  /* Output:
  {
    "http://schema.org/name": [{"@value": "Manu Sporny"}],
    "http://schema.org/url": [{"@id": "http://manu.sporny.org/"}],
    "http://schema.org/image": [{"@id": "http://manu.sporny.org/images/manu.png"}]
  }
  */
});

// expand using URLs
jsonld.expand('http://example.org/doc', ...);

// flatten a document
// see: http://json-ld.org/spec/latest/json-ld/#flattened-document-form
jsonld.flatten(doc, (err, flattened) => {
  // all deep-level trees flattened to the top-level
});

// frame a document
// see: http://json-ld.org/spec/latest/json-ld-framing/#introduction
jsonld.frame(doc, frame, (err, framed) => {
  // document transformed into a particular tree structure per the given frame
});

// canonize (normalize) a document using the RDF Dataset Normalization Algorithm
// (URDNA2015), see: http://json-ld.github.io/normalization/spec/
jsonld.canonize(doc, {
  algorithm: 'URDNA2015',
  format: 'application/n-quads'
}, (err, canonized) => {
  // canonized is a string that is a canonical representation of the document
  // that can be used for hashing, comparison, etc.
});

// serialize a document to N-Quads (RDF)
jsonld.toRDF(doc, {format: 'application/n-quads'}, (err, nquads) => {
  // nquads is a string of N-Quads
});

// deserialize N-Quads (RDF) to JSON-LD
jsonld.fromRDF(nquads, {format: 'application/n-quads'}, (err, doc) => {
  // doc is JSON-LD
});

// register a custom async-callback-based RDF parser
jsonld.registerRDFParser(contentType, (input, callback) => {
  // parse input to a jsonld.js RDF dataset object...
  callback(err, dataset);
});

// register a custom synchronous RDF parser
jsonld.registerRDFParser(contentType, input => {
  // parse input to a jsonld.js RDF dataset object... and return it
  return dataset;
});

// use the promises API:

// compaction
const compacted = await jsonld.compact(doc, context);

// expansion
const expanded = await jsonld.expand(doc);

// flattening
const flattened = await jsonld.flatten(doc);

// framing
const framed = await jsonld.frame(doc, frame);

// canonicalization (normalization)
const canonized = await jsonld.canonize(doc, {format: 'application/n-quads'});

// serialize to RDF
const rdf = await jsonld.toRDF(doc, {format: 'application/n-quads'});

// deserialize from RDF
const doc = await jsonld.fromRDF(nquads, {format: 'application/n-quads'});

// register a custom promise-based RDF parser
jsonld.registerRDFParser(contentType, async input => {
  // parse input into a jsonld.js RDF dataset object...
  return new Promise(...);
});

// how to override the default document loader with a custom one -- for
// example, one that uses pre-loaded contexts:

// define a mapping of context URL => context doc
const CONTEXTS = {
  "http://example.com": {
    "@context": ...
  }, ...
};

// grab the built-in node.js doc loader
const nodeDocumentLoader = jsonld.documentLoaders.node();
// or grab the XHR one: jsonld.documentLoaders.xhr()

// change the default document loader using the callback API
// (you can also do this using the promise-based API, return a promise instead
// of using a callback)
const customLoader = (url, callback) => {
  if(url in CONTEXTS) {
    return callback(
      null, {
        contextUrl: null, // this is for a context via a link header
        document: CONTEXTS[url], // this is the actual document that was loaded
        documentUrl: url // this is the actual context URL after redirects
      });
  }
  // call the underlining documentLoader using the callback API.
  nodeDocumentLoader(url, callback);
  /* Note: By default, the node.js document loader uses a callback, but
  browser-based document loaders (xhr or jquery) return promises if they
  are supported (or polyfilled) in the browser. This behavior can be
  controlled with the 'usePromise' option when constructing the document
  loader. For example: jsonld.documentLoaders.xhr({usePromise: false}); */
};
jsonld.documentLoader = customLoader;

// alternatively, pass the custom loader for just a specific call:
const compacted = await jsonld.compact(
doc, context, {documentLoader: customLoader});
