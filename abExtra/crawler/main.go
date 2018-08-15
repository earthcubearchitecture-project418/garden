package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	_ "net/http/pprof"
	"os"
	"time"

	"earthcube.org/Project418/crawler/crawl"
	"github.com/blevesearch/bleve"
)

// var cpuprofile = flag.String("cpuprofile", "", "write cpu profile `file`")
// var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func main() {
	fmt.Println("P418 Indexer")

	urlPtr := flag.String("url", "", "a url to request a sitemap from")
	csvPtr := flag.String("csv", "", "a csv file with domains to process")
	sitemapPtr := flag.String("sitemap", "", "a local sitemap to process")
	indexnamePtr := flag.String("indexname", "p418.bleve", "name for the bleve text index")
	spatialPtr := flag.Bool("spatial", true, "boolean for doing the spatial index aspect")
	flag.Parse()

	// Profile setup
	// if *cpuprofile != "" {
	// 	f, err := os.Create(*cpuprofile)
	// 	if err != nil {
	// 		log.Fatal("could not create CPU profile: ", err)
	// 	}
	// 	if err := pprof.StartCPUProfile(f); err != nil {
	// 		log.Fatal("could not start CPU profile: ", err)
	// 	}
	// 	defer pprof.StopCPUProfile()
	// }

	// Set up our log file for runs...
	f, err := os.OpenFile("logfile.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	initBleve(*indexnamePtr) // init a new Bleve index

	log.Printf("Start time: %s \n", time.Now()) // Log the time at start for the record

	// Need to have a list of domains we are going to index and the sitemap at them
	// This should be a facilityDomains.csv file.

	indexFileName := *indexnamePtr
	spatial := *spatialPtr

	if *sitemapPtr != "" {
		err := crawl.ProcessLocalSitemap(*sitemapPtr, indexFileName, spatial)
		if err != nil {
			log.Println("Error from process sitemap function")
		}
		log.Print("MAIN report: Crawler returned")
	} else {
		var domainList []string
		if *csvPtr != "" {
			domainList, _ = readListFromFile(*csvPtr)
		}

		if *urlPtr != "" {
			urlToIndex := *urlPtr // needed?
			domainList = append(domainList, urlToIndex)
		}
		count := 0
		total := len(domainList)
		for key := range domainList {
			count = count + 1
			// percent := count / total
			fmt.Printf("%d of %d  domains\n", count, total)
			err := crawl.ProcessDomain(domainList[key], indexFileName, spatial)
			if err != nil {
				log.Println("Error from process domain function")
			}
			log.Print("MAIN report: Crawler returned")
		}
	}

	// Log the time at end for the record
	log.Printf("End time: %s \n", time.Now())

	// Profiling shutdown
	// if *memprofile != "" {
	// 	f, err := os.Create(*memprofile)
	// 	if err != nil {
	// 		log.Fatal("could not create memory profile: ", err)
	// 	}
	// 	runtime.GC() // get up-to-date statistics
	// 	if err := pprof.WriteHeapProfile(f); err != nil {
	// 		log.Fatal("could not write memory profile: ", err)
	// 	}
	// 	f.Close()
	// }
}

// Initialize the text index
func initBleve(filename string) {
	// TODO resolve this hack
	// open a new index
	mapping := bleve.NewIndexMapping()
	index, berr := bleve.New(filename, mapping)
	if berr != nil {
		log.Printf("Bleve error making index %v \n", berr)
	}
	index.Close()
}

// read the CSV file contents as string array
func readListFromFile(filename string) ([]string, error) {
	var entries []string

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		// if (strings.STARTSWITH(# , scanner.Text())
		entries = append(entries, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return entries, nil
}
