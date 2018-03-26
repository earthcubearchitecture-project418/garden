### Tile 38

#### About
The tile38 geohash database will be used for P418.  The docker file here is 
from the Tile38 repo and will be modded as needed to address P418 needs.

#### Notes

Can also just pull from dockerhub with ```docker pull tile38/tile38```

```
docker build -f Tile38Dockerfile -t earthcube/p418spatialindex:latest  . 
```

```
docker run -d -p 9851:9851 --name p418spatial -t earthcube/p418spatialindex:latest 
```

With persistent data store
```
docker run -d -p 9851:9851 -v /Users/dfils/Data/OCDDataVolumes/spatialindex:/data --name p418spatial -t earthcube/p418spatialindex:latest 

```

```
docker run -d -p 9851:9851 -v /Users/dfils/Data/OCDDataVolumes/spatialindex:/data  -t earthcube/p418spatialindex:latest 

docker run -d -p 9851:9851 -v /Users/dfils/Data/OCDDataVolumes/spatialindex:/data  -t tile38/tile38

```