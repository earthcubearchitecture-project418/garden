# BLAST

## About
Some notes about using the Blast package to perform text indexing and search

### Notes

```
docker run -i -p 5000:5000 -p 8000:8000 earthcube/blast:0.0.1
```

```

docker run -d -p 5000:5000 -v ~/Data/P418/dataVolumes/blast:/data -p 8000:8000 earthcube/blast:0.0.1

```

```
 cat ./example/search_request2.json | xargs -0 ./bin/blast search --request
```
