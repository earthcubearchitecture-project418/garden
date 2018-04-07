package main

import (
  "encoding/json"
  "fmt"
  "log"
  "strings"

  "github.com/garyburd/redigo/redis"
  geojson "github.com/paulmach/go.geojson"
)

type LocType struct {
  Type string `json:"type"`
}

func main() {
  fmt.Println("this is a simple spatial test")

  SpatialCall(testgeom(), "")

}

func testgeom() string {

  data := `{
  "type": "FeatureCollection",
  "features": [
    {
    "type": "Feature",
    "properties": {},
    "geometry": {
      "type": "Polygon",
      "coordinates": [
      [
        [
        -112.8515625,
        -29.535229562948444
        ],
        [
        85.4296875,
        -29.535229562948444
        ],
        [
        85.4296875,
        65.36683689226321
        ],
        [
        -112.8515625,
        65.36683689226321
        ],
        [
        -112.8515625,
        -29.535229562948444
        ]
      ]
      ]
    }
    }
  ]
  }
  `

  return data
}

func redisDial() (redis.Conn, error) {
  // c, err := redis.Dial("tcp", "tile38:9851")
  c, err := redis.Dial("tcp", "localhost:9851")
  if err != nil {
    log.Printf("Could not connect: %v\n", err)
  }
  return c, err
}

func SpatialCall(geowithin, filter string) {
  // geowithin := request.QueryParameter("geowithin")
  // filter := request.QueryParameter("filter")
  // log.Printf("Called with filter: %s and geojson %s \n", filter, geowithin)

  _, err := geojson.UnmarshalFeatureCollection([]byte(geowithin))
  if err != nil {
    fmt.Println(err)
  }

  c, err := redisDial()
  defer c.Close()

  var value1 int
  var value2 []interface{}
  // TODO  fix the 50K request limit, put in cursor pattern
  reply, err := redis.Values(c.Do("WITHIN", "p418", "LIMIT", "50000", "OBJECT", geowithin))
  //WITHIN p418 BOUNDS -90 -40 17 50
  // reply, err := redis.Values(c.Do("WITHIN", "p418", "BOUNDS", -90, -40, 17, 50))
  // reply, err := redis.Values(c.Do("SCAN", "p418"))  // an early test call just to get everything
  if err != nil {
    fmt.Printf("Error in reply %v \n", err)
  }
  if _, err := redis.Scan(reply, &value1, &value2); err != nil {
    fmt.Printf("Error in scan %v \n", err)
  }

  // log.Println(value1) // the point of this logging is what?  the point of value1 is what!?
  // log.Println(value2)

  results, err := redisToGeoJSON(value2, filter)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(results)
}

func redisStringToGeoJSON(m map[string]string) (string, error) {

  fc := geojson.NewFeatureCollection()

  for k, v := range m {
    fmt.Println("k:", k, "v:", v)

    rawGeometryJSON := []byte(v)
    ID := k

    g, err := geojson.UnmarshalGeometry(rawGeometryJSON)
    if err != nil {
      log.Printf("Unmarshal geom error for %s with %s\n", rawGeometryJSON, err)
    }

    switch {
    case g.IsPoint():
      nf := geojson.NewFeature(g)
      nf.SetProperty("URL", ID)
      fc.AddFeature(nf)
    case g.IsPolygon():
      nf := geojson.NewFeature(g)
      nf.SetProperty("URL", ID)
      fc.AddFeature(nf)
    default:
      log.Println(g.Type)
    }

    if g.Type == "Feature" {
      f, err := geojson.UnmarshalFeature(rawGeometryJSON)
      if err != nil {
        log.Printf("Unmarshal feature error for %s with %s\n", ID, err)
      }
      f.SetProperty("URL", ID)
      fc.AddFeature(f)
    }

  }

  rawJSON, err := fc.MarshalJSON()
  if err != nil {
    return "", err
  }

  return string(rawJSON), nil
}

func redisToGeoJSON(results []interface{}, filter string) (string, error) {

  fc := geojson.NewFeatureCollection()

  for _, item := range results {
    valcast := item.([]interface{})
    val0 := fmt.Sprintf("%s", valcast[0])
    val1 := fmt.Sprintf("%s", valcast[1])
    //log.Printf("%s %s \n", val0, val1)

    if strings.Contains(val0, filter) || filter == "" {

      lt := &LocType{}
      err := json.Unmarshal([]byte(val1), lt)
      if err != nil {
        log.Print(err)
        return "", err
      }

      rawGeometryJSON := []byte(val1)

      switch (lt.Type) {
        case "FeatureCollection":
          fcf, err := geojson.UnmarshalFeatureCollection(rawGeometryJSON)
          if err != nil {
            log.Printf("Unmarshal featurecollection error for %s with %s\n", val0, err)
          }
          for _, f := range fcf.Features {
            f.SetProperty("URL", val0)
            fc.AddFeature(f)
          }
          break

        case "Feature":
          f, err := geojson.UnmarshalFeature(rawGeometryJSON)
          if err != nil {
            log.Printf("Unmarshal feature error for %s with %s\n", val0, err)
          }
          f.SetProperty("URL", val0)
          fc.AddFeature(f)
          break

        case "Point":
        case "Poly":
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
          break
      }
    }
  }

  rawJSON, err := fc.MarshalJSON()
  if err != nil {
    return "", err
  }

  return string(rawJSON), nil
}
