package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/garyburd/redigo/redis"
	geojson "github.com/paulmach/go.geojson"
)

// LocType is there to do a first cut marshalling to just get the type before  next marshalling
type LocType struct {
	Type string `json:"type"`
}

func main() {
	c, err := redisDial()
	defer c.Close()

	if err != nil {
		log.Printf("Error on body parameter read %v \n", err)
	}

	var value1 int
	var value2 []interface{}
	reply, err := redis.Values(c.Do("SCAN", "p418", "LIMIT", "50000"))
	if err != nil {
		fmt.Println(err)
		return
	}

	if _, err := redis.Scan(reply, &value1, &value2); err != nil {
		fmt.Printf("Error in scan %v \n", err)
	}

	log.Println(value1) // the point of this logging is what?  the point of value1 is what!?
	// log.Println(value2)

	filter := ""
	results, err := redisToGeoJSON(value2, filter)
	if err != nil {
		log.Println(err)
	}

	// log.Println(results)

	// write to a file
	f, err := os.Create("features.json")
	if err != nil {
		log.Println("error making file")
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	n4, err := w.WriteString(results)
	fmt.Printf("wrote %d bytes\n", n4)
	w.Flush()
}

func redisToGeoJSON(results []interface{}, filter string) (string, error) {

	fc := geojson.NewFeatureCollection()

	for _, item := range results {
		valcast := item.([]interface{})
		val0 := fmt.Sprintf("%s", valcast[0])
		val1 := fmt.Sprintf("%s", valcast[1])
		// log.Printf("%s %s \n", val0, val1)

		if strings.Contains(val0, filter) || filter == "" {

			lt := &LocType{}
			err := json.Unmarshal([]byte(val1), lt)
			if err != nil {
				log.Print(err)
				return "", err
			}

			rawGeometryJSON := []byte(val1)

			if lt.Type == "Feature" {
				f, err := geojson.UnmarshalFeature(rawGeometryJSON)
				if err != nil {
					log.Printf("Unmarshal feature error for %s with %s\n", val0, err)
				}
				f.SetProperty("URL", val0)
				fc.AddFeature(f)
			}

			if lt.Type == "Point" || lt.Type == "Poly" {
				g, err := geojson.UnmarshalGeometry(rawGeometryJSON)
				if err != nil {
					log.Printf("Unmarshal geom error for %s with %s\n", val0, err)
				}

				switch {
				case g.IsPoint():
					nf := geojson.NewFeature(g)
					nf.SetProperty("URL", val0)
					fc.AddFeature(nf)
				case g.IsPolygon():
					nf := geojson.NewFeature(g)
					nf.SetProperty("URL", val0)
					fc.AddFeature(nf)
				default:
					log.Println(g.Type)
				}
			}
		}
	}

	rawJSON, err := fc.MarshalJSON()
	if err != nil {
		return "", err
	}

	return string(rawJSON), nil
}

func redisDial() (redis.Conn, error) {
	//c, err := redis.Dial("tcp", "tile38:9851")
	c, err := redis.Dial("tcp", "localhost:9851")
	if err != nil {
		log.Printf("Could not connect: %v\n", err)
	}
	return c, err
}
