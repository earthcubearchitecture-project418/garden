package millersmock

import (
	//	"bytes"
	"bytes"
	"crypto/sha1"
	"fmt"
	"log"
	"sync"
	//	"log"

	minio "github.com/minio/minio-go"
)

// Entry is a struct holding the json-ld metadata and data (the text)
type Entry struct {
	Bucketname string
	Key        string
	Urlval     string
	Sha1val    string
	Jld        string
}

// MockObjects test a concurrent version of calling mock
func MockObjects(mc *minio.Client, bucketname string) {
	doneCh := make(chan struct{}) // Create a done channel to control 'ListObjectsV2' go routine.
	defer close(doneCh)           // Indicate to our routine to exit cleanly upon return.
	isRecursive := true
	objectCh := mc.ListObjectsV2(bucketname, "", isRecursive, doneCh)

	var entries []Entry

	for object := range objectCh {
		if object.Err != nil {
			fmt.Println(object.Err)
			return
		}

		fo, err := mc.GetObject(bucketname, object.Key, minio.GetObjectOptions{})
		if err != nil {
			fmt.Println(err)
			return
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
	multiCall(entries)

}

func multiCall(e []Entry) {

	// Set up the the semaphore and conccurancey
	semaphoreChan := make(chan struct{}, 20) // a blocking channel to keep concurrency under control
	defer close(semaphoreChan)
	wg := sync.WaitGroup{} // a wait group enables the main process a wait for goroutines to finish

	for k := range e {
		wg.Add(1)
		fmt.Printf("About to run #%d in a goroutine\n", k)
		go func(k int) {
			semaphoreChan <- struct{}{}
			status := Mock(e[k].Bucketname, e[k].Key, e[k].Urlval, e[k].Sha1val, e[k].Jld)

			wg.Done() // tell the wait group that we be done
			log.Printf("#%d done with %s", k, status)
			<-semaphoreChan
		}(k)
	}
	wg.Wait()
}

// Mock is a simple function to use as a stub for talking about millers
func Mock(bucketname, key, urlval, sha1val, jsonld string) string {
	fmt.Printf("%s:  %s %s   %s =? %s \n", bucketname, key, urlval, sha1val, getsha(jsonld))

	return "ok"
}

func getsha(jsonld string) string {
	h := sha1.New()
	h.Write([]byte(jsonld))
	bs := h.Sum(nil)
	bss := fmt.Sprintf("%x", bs)
	return bss
}
