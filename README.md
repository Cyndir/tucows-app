# tucows-app

Implementation of the Tucows interview exercise https://github.com/tucows/interview-exercise-alt

Utilizes containerized versions of postgres and rabbitMQ running locally

How to run:

Requirements:
docker
go version 1.20+

How to run:
in the root of this repo, run `docker compose up`
in a terminal, navigate to cmd/ordermanagement
 `go run main.go`

in a second terminal, navigate to cmd/paymentprocessor
 `go run main.go`


http request examples are in http.requests