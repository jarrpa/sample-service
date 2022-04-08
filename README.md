# URL Shortening Service

This is an attempt to put together a URL shortening service. Rather than
starting from scratch, I cloned a sample repo as a base to cover a lot of the
boilerplate structure. I then copied the relevant files and modified them to
design the desired service. Check the git history for more details.

This service is intended to generate a unique short name for any URL it's
presented. It responds to GET requests with the URL corresponding to the
shortname specified in the request. It also reponds to POST requests to present
a complete URL and return its shortened name.

The project is incomplete, but the API and basic handler functions are more
or less in place.

Repo structure:

```
.
├── api
│   └── v1
│       ├── pb
│       │   └── shorty
│       │       ├── shorty_grpc.pb.go
│       │       └── shorty.proto
│       └── swagger.json
├── cmd
│   └── shorty
│       └── shorty.go
├── deploy
│   └── shorty-svc.yaml
├── go.mod
├── go.sum
├── hack
│   ├── make-help.sh
│   └── scripts
│       ├── gocoverage.sh
│       └── pb-compile.sh
├── images
│   └── shorty
│       ├── Dockerfile
│       └── Makefile
├── internal
│   ├── model
│   │   ├── constants.go
│   │   └── orm.go
│   └── util
│       ├── errors.go
│       ├── logger.go
│       └── logrusutil.go
├── LICENSE
├── Makefile
├── OWNERS
└── pkg
    └── shorty
        ├── endpoints
        │   ├── endpoints.go
        │   └── reqJSONMap.go
        ├── service.go
        ├── shorty.go
        └── transport
            ├── grpc.go
            └── http.go
```
