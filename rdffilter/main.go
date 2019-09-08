package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/knakk/rdf"
)

const (
	chunksize int = 1024
)

var (
	data  *os.File
	part  []byte
	err   error
	count int
)

func main() {
	// fileHandle, _ := os.Open("test.nt")
	// defer fileHandle.Close()
	// fileScanner := bufio.NewScanner(fileHandle)

	// for fileScanner.Scan() {
	// 	// TODO can I skip blank lines here?
	// 	f, e := goodTriples(fileScanner.Text())
	// 	// store the good and bad (erorr with bad) in buffers
	// 	// and record to minio
	// 	if e != nil {
	// 		fmt.Println(e)
	// 		fmt.Print(f)
	// 	}
	// }

	// var rdf bytes.Buffer
	length, rdf := openFile("test.nt")
	fmt.Println(length)

	scanner := bufio.NewScanner(rdf) // rdf is already a pointer
	good := bytes.NewBuffer(make([]byte, 0))
	bad := bytes.NewBuffer(make([]byte, 0))
	for scanner.Scan() {
		// log.Printf("rdf: %s", scanner.Text())
		f, e := goodTriples(scanner.Text())
		if e == nil {
			_, err = good.Write([]byte(f))
		}
		if e != nil {
			_, err = bad.Write([]byte(fmt.Sprintf("%s :Error: %s\n", strings.TrimSuffix(f, "\n"), e)))
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(good.String())
	fmt.Println(bad.String())

}

// TODO  convert this to use a bytes.Buffer  (or better a pointer to that)
func goodTriples(f string) (string, error) {
	dec := rdf.NewTripleDecoder(strings.NewReader(f), rdf.NTriples)
	triple, err := dec.Decode()
	return triple.Serialize(rdf.NTriples), err
}

func openFile(name string) (byteCount int, buffer *bytes.Buffer) {

	data, err = os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	reader := bufio.NewReader(data)
	buffer = bytes.NewBuffer(make([]byte, 0))
	part = make([]byte, chunksize)

	for {
		if count, err = reader.Read(part); err != nil {
			break
		}
		buffer.Write(part[:count])
	}
	if err != io.EOF {
		log.Fatal("Error Reading ", name, ": ", err)
	} else {
		err = nil
	}

	byteCount = buffer.Len()
	return
}
