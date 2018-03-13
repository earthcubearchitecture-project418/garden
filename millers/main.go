package main

import (
	"fmt"
	"log"
	"time"

	"earthcube.org/Project418/garden/millers/millerspatial"

	"github.com/minio/minio-go"
)

func main() {
	fmt.Println("The miller....")
	st := time.Now()
	log.Printf("Start time: %s \n", st) // Log the time at start for the record

	mc := miniConnection() // minio connection

	buckets, err := listBuckets(mc)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Bucket list...")
	for _, bucket := range buckets {
		fmt.Println(bucket.Name)
		// processBucketObjects(mc, bucket.Name)
	}

	// MOCK call
	// millersmock.MockObjects(mc, "getiedadataorg")

	// GRAPH calls
	// millersgraph.GraphMillObjects(mc, "getiedadataorg")
	// millersgraph.GraphMillObjects(mc, "opentopographyorg")
	// millersgraph.GraphMillObjects(mc, "dataneotomadborg")

	// SPATIAL calls
	millerspatial.ProcessBucketObjects(mc, "opentopographyorg")
	// processBucketObjects(mc, "dataneotomadborg")
	// processBucketObjects(mc, "getiedadataorg")
	// processBucketObjects(mc, "opencoredataorg")
	// processBucketObjects(mc, "wwwbco-dmoorg")

	et := time.Now()
	diff := et.Sub(st)
	log.Printf("End time: %s \n", et)
	log.Printf("Run time: %f \n", diff.Minutes())
}

func miniConnection() *minio.Client {
	// Set up minio and initialize client
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
		fmt.Println(err)
		return nil, err
	}

	return buckets, err
}
