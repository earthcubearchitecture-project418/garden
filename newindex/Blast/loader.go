package main

import (
	"bytes"
	"fmt"
	"log"
	"sync"

	"github.com/go-resty/resty"
	minio "github.com/minio/minio-go"
)

type Entry struct {
	Bucketname string
	Key        string
	Urlval     string
	Sha1val    string
	Jld        string
}

func main() {
	fmt.Println("blast loader")
	mc := miniConnection() // minio connection

	// for loop these...
	f := []string{"baltoopendaporg", "cdfregistry", "csdco", "dataneotomadborg", "dsirisedu", "earthreforg",
		"getiedadataorg", "opencoredataorg", "opentopographyorg", "packages", "scientificdrillingdataorg",
		"wikilinkedearth", "wwwbco-dmoorg", "wwwhydroshareorg", "wwwunavcoorg"}

	// for loop these...
	// f := []string{"opencoredataorg", "opentopographyorg"}

	for x := range f {
		entries := getObjects(mc, f[x])
		multiCall(entries)

	}

}

// GetMillObjects
func getObjects(mc *minio.Client, bucketname string) []Entry {
	doneCh := make(chan struct{}) // Create a done channel to control 'ListObjectsV2' go routine.
	defer close(doneCh)           // Indicate to our routine to exit cleanly upon return.
	isRecursive := true
	objectCh := mc.ListObjectsV2(bucketname, "", isRecursive, doneCh)

	var entries []Entry

	for object := range objectCh {
		if object.Err != nil {
			fmt.Println(object.Err)
			return nil
		}

		fo, err := mc.GetObject(bucketname, object.Key, minio.GetObjectOptions{})
		if err != nil {
			fmt.Println(err)
			return nil
		}

		oi, err := fo.Stat()
		if err != nil {
			log.Println("Issue with reading an object..  should I just fatal on this to make sure?")
		}
		urlval := oi.Metadata["X-Amz-Meta-Url"][0] // also have  X-Amz-Meta-Sha1
		sha1val := oi.Metadata["X-Amz-Meta-Sha1"][0]
		buf := new(bytes.Buffer)
		buf.ReadFrom(fo)
		jld := buf.String() // Does a complete copy of the bytes in the buffer.

		// Mock call for some validation (and a template for other millers)
		// Mock(bucketname, object.Key, urlval, sha1val, jld)
		entry := Entry{Bucketname: bucketname, Key: object.Key, Urlval: urlval, Sha1val: sha1val, Jld: jld}
		entries = append(entries, entry)

	}

	fmt.Println(len(entries))
	// multiCall(entries)

	return entries
}

func multiCall(e []Entry) {
	// Set up the the semaphore and conccurancey
	semaphoreChan := make(chan struct{}, 1) // a blocking channel to keep concurrency under control
	defer close(semaphoreChan)
	wg := sync.WaitGroup{} // a wait group enables the main process a wait for goroutines to finish

	for k := range e {
		wg.Add(1)
		fmt.Printf("About to run #%d in a goroutine\n", k)
		go func(k int) {
			semaphoreChan <- struct{}{}
			status := wrapLoad(e[k].Bucketname, e[k].Key, e[k].Urlval, e[k].Jld)

			wg.Done() // tell the wait group that we be done
			log.Printf("#%d done with %d", k, status)
			<-semaphoreChan
		}(k)
	}
	wg.Wait()
}

func wrapLoad(bucket, key, urlval, jld string) int {
	//     { "id": "1", "fields":
	jb := fmt.Sprintf("{\"document\": { \"id\": \"%s\", \"fields\": %s }}", urlval, jld)

	puturl := fmt.Sprintf("http://0.0.0.0:8000/rest/%s", key)

	resp, err := resty.R().
		SetBody(jb).
		Put(puturl)
	if err != nil {
		log.Print(err)
	}

	return resp.StatusCode()
}

// Set up minio and initialize client
func miniConnection() *minio.Client {
	endpoint := "localhost:9000"
	accessKeyID := "AKIAIOSFODNN7EXAMPLE"
	secretAccessKey := "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"
	useSSL := false
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}
	return minioClient
}

func listBuckets(mc *minio.Client) ([]minio.BucketInfo, error) {
	buckets, err := mc.ListBuckets()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return buckets, err
}
