# tucows-app

Implementation of the Tucows interview exercise https://github.com/tucows/interview-exercise-alt

Utilizes containerized versions of postgres and rabbitMQ running locally

## Requirements:
docker
go version 1.20+

## How to run:
in the root of this repo, run `docker compose up`
in another terminal, navigate to cmd/ordermanagement
 `go run main.go`

in another terminal, navigate to cmd/paymentprocessor
 `go run main.go`

http request examples are in http.requests

## How to Test:

unit tests: go test ./...
integration tests: run `docker compose up` if not already started, `go test --tags=integration ./...`


## Improvements from here:

More dependencies could be injected (the http client on messageProcessor being one)

Messageprocessor, once refactored, needs unit tests

Observability and logging are very minimal right now

Production hardening definitely necessary, currently lots of security problems since it was written as a minimalistic proof-of-concept style

Needs to be configurable instead of hardcoded connection strings, ports, etc.

Error handling currently exists in several different styles. This should be unified.

Middleware could be utilized to clean up some duplication

Could make docker containers for go binaries and add those to the docker compose file.

Implement some github actions for quality gates (code coverage, run tests, etc)