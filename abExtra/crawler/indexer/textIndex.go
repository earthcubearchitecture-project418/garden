package indexer

import (
	"log"

	"github.com/blevesearch/bleve"
)

// TODO
// 1) how best to pass index pointer/reference  (thread safety issue!!!!!!!)
// 2) how best to ID these entries
// 3)

// TextIndexer indexes data in a text index
func TextIndexer(ID string, jsonld string, index bleve.Index) string {
	// index, berr := bleve.New(textindex, mapping)
	// index, berr := bleve.Open(textindex)
	// if berr != nil {
	// 	log.Printf("Bleve error making index %v \n", berr)
	// }

	// index some data
	berr := index.Index(ID, jsonld)
	// log.Printf("Blevel Indexed item with ID %s\n", ID)
	if berr != nil {
		log.Printf("Bleve error indexing %v \n", berr)
	}

	// index.Close()

	return "done"
}
