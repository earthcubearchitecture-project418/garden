package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"earthcube.org/Project418/garden/summoner/acquire"
	"earthcube.org/Project418/garden/summoner/utils"
)

func main() {
	// Set up our log file for runs...
	f, err := os.OpenFile("logfile.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	log.Printf("Start time: %s \n", time.Now()) // Log start time

	cfgFileLoc := flag.String("config", "", "JSON configure file")
	flag.Parse()

	cs := loadConfiguration(cfgFileLoc)

	domains, err := acquire.DomainList(cs.Source)
	if err != nil {
		log.Printf("Error reading list of domains %v\n", err)
	}
	ru := acquire.ResourceURLs(domains) // map by domain name and []string of landing page URLs
	acquire.ResRetrieve(ru, cs)

	log.Printf("End time: %s \n", time.Now()) // Log end time
}

func loadConfiguration(file *string) utils.Config {
	var config utils.Config
	configFile, err := os.Open(*file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}
