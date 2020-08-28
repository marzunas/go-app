# API to compute SHA-512 hash for secrets
This API module makes several end points available to compute, fetch hash and retrieve statistics of the hash ops.

The service exposes following end points :

* /hash - Accepts a parameter with name "password" to compute hash for. The hash is not immediately available, a unique
request identifier provided with the ACK response, which can be used to retrieve the hash by invoking /hash/{id}.
```bash
curl -d "password=angryMonkey5" http://localhost:8080/hash
10
```

* /stats - Returns JSON response with total count of requests and average time taken by hash compute function.
```bash
curl -d "password=angryMonkey5" http://localhost:8080/stats
{"total":3,"average":1747299}
```

* /shutdown - Shuts down the service. Any requests received after shutdown initiation will be rejected. 
```bash
curl http://localhost:8080/shutdown
```

## Build
Run build operation from source root directory to produce build artifact
```bash
go build
```

## Run
Run the application binary to start the service
```bash
./go-app
2020/08/27 16:54:10 initializing http server
starting server at port 8080
```

