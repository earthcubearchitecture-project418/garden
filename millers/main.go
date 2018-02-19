package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/minio/minio-go"
)

func main() {
	fmt.Println("The miller....")

	// TODO
	// read the entries of a bucket..  ListObjectsv2
	// for each item in object list  GetObject  (use for StatOnject?)
	// Send the results (copy of results) to each of the active mills  Start with spatial

	getBucketList()
}

func getBucketList() {
	fmt.Println("Get bucket list")

	// Set up minio and initialize client
	endpoint := "localhost:9000"
	accessKeyID := "AKIAIOSFODNN7EXAMPLE"
	secretAccessKey := "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"
	useSSL := false
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}

	// Create a done channel to control 'ListObjectsV2' go routine.
	doneCh := make(chan struct{})

	// Indicate to our routine to exit cleanly upon return.
	defer close(doneCh)

	isRecursive := true
	objectCh := minioClient.ListObjectsV2("opencoredataorg", "", isRecursive, doneCh)
	for object := range objectCh {
		if object.Err != nil {
			fmt.Println(object.Err)
			return
		}
		fmt.Println(object)

		fo, err := minioClient.GetObject("opencoredataorg", object.Key, minio.GetObjectOptions{})
		if err != nil {
			fmt.Println(err)
			return
		}

		// Use io.copy
		localFile, err := os.Create("./local-file.jsonld")
		if err != nil {
			fmt.Println(err)
			return
		}
		if _, err = io.Copy(localFile, fo); err != nil {
			fmt.Println(err)
			return
		}

	}

}
