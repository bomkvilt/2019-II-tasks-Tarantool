# setup enviroment
FROM golang:alpine as builder
ENV GO111MODULE=on
RUN apk add --update --no-cache git;

# get packages
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
COPY --from=builder ./bin/tests ./bin/tests
CMD './bin/tests'
