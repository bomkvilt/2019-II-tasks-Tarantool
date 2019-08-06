# setup enviroment
FROM golang:alpine as builder
ENV GO111MODULE=on
RUN apk add --update --no-cache git;

WORKDIR /bootcamp
COPY ./go.mod .
RUN go mod download;

# build tests
COPY './tests' './tests'
RUN go build		\
	-o ./bin/tests  \
	./tests/cmd/main.go;

# run tests
FROM alpine:latest
WORKDIR /bootcamp
COPY 				'./tests/cmd/config.yaml' 	'./tests/cmd/config.yaml'
COPY --from=builder '/bootcamp/bin/tests' 		'./bin/tests'
CMD './bin/tests'
