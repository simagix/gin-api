# Gin API Server
## Overview

Gin API server template

### Linux Build
```
# ./gradlew clean buildLinux
```

### Docker Build
```
# ./docker.sh
# docker build -t simagix/gin-api:0.1.0 -f src/docker/Dockerfile .
```

### Unit Test
```
# export MYDB_URI=mongodb://user:password@localhost:27017/MYDB?authSource=admin
# cd api; go test -run VMPostStats
```

### Usage
```
Usage of /osca-api-server:
  -logfile string
    	a string
  -version
    	a bool
```
