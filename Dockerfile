FROM golang:latest as build

ENV GOOS=linux
ENV GO111MODULE=on

WORKDIR $GOPATH/src/github.com/grandcolline/iap-test

COPY . .
RUN env CGO_ENABLED=0 go install

FROM gcr.io/distroless/static as run
COPY --from=build /go/bin/iap-test /iap-test
CMD ["/iap-test"]

