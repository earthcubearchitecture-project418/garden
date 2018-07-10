package main

import (
	"bytes"
	"fmt"
	"log"
	"sync"

	"github.com/go-resty/resty"
	minio "github.com/minio/minio-go"
	"github.com/tidwall/sjson"
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
		"getiedadataorg", "opencoredataorg", "opentopographyorg",
		"wikilinkedearth", "wwwbco-dmoorg", "wwwhydroshareorg", "wwwunavcoorg"}

	// for loop these...
	// f := []string{"opencoredataorg", "opentopographyorg"}

	for x := range f {
		entries := getObjects(mc, f[x])
		multiCall(entries)

	}
}

func namemapping(url string) string {
	m := map[string]string{"baltoopendaporg": "balto",
		"cdfregistry": "cdf", "csdco": "csdco", "dataneotomadborg": "neotoma",
		"dsirisedu": "iris", "earthreforg": "magic", "getiedadataorg": "ieda",
		"opencoredataorg": "opencore", "opentopographyorg": "opentopo",
		"wikilinkedearth": "linkedearth", "wwwbco-dmoorg": "bcodmo",
		"wwwhydroshareorg": "hydroshare", "wwwunavcoorg": "unavco"}

	return m[url]
}

func logomapping(url string) string {
	m := map[string]string{"baltoopendaporg": "http://geodex.org/images/logos/EarthCubeLogo.png",
		"cdfregistry":       "http://geodex.org/images/logos/EarthCubeLogo.png",
		"csdco":             "http://geodex.org/images/logos/csdco.png",
		"dataneotomadborg":  "http://geodex.org/images/logos/neotoma.png",
		"dsirisedu":         "http://geodex.org/images/logos/iris_color_screen_lrg.png",
		"earthreforg":       "http://geodex.org/images/logos/magic.png",
		"getiedadataorg":    "http://geodex.org/images/logos/ieda_maplogo.png",
		"opencoredataorg":   "http://geodex.org/images/logos/ocd_logo.jpg",
		"opentopographyorg": "http://geodex.org/images/logos/opentopo.png",
		"wikilinkedearth":   "http://geodex.org/images/logos/linkedEarth.jpeg",
		"wwwbco-dmoorg":     "http://geodex.org/images/logos/bco-dmo-words-BLUE.jpg",
		"wwwhydroshareorg":  "http://geodex.org/images/logos/cuahsiHydroshare.png",
		"wwwunavcoorg":      "http://geodex.org/images/logos/uv-logo.png"}

	return m[url]
}

// GetMillObjects
func getObjects(mc *minio.Client, bucketname string) []Entry {
	fmt.Printf("Getting object for bucket %s\n", bucketname)
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

	// fmt.Printf("%s   %s    %s  \n", bucket, key, urlval)

	// TODO..  modify the JSON-LD document to add in another p418 fields
	//
	// Use tidwalls awesome stuff...
	value, _ := sjson.Set(jld, "p418bucket", bucket)
	value, _ = sjson.Set(value, "p418url", urlval)
	value, _ = sjson.Set(value, "p418source", namemapping(bucket))
	value, _ = sjson.Set(value, "p418logo", logomapping(bucket))

	//     { "id": "1", "fields":
	jb := fmt.Sprintf("{\"document\": { \"id\": \"%s\", \"bucket\": \"%s\",  \"fields\": %s }}", key, bucket, value)

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
