FROM golang:latest as build

ENV GOOS=linux
ENV GO111MODULE=on

WORKDIR $GOPATH/src/github.com/grandcolline/iap_test

COPY . .
RUN env CGO_ENABLED=0 go install

FROM gcr.io/distroless/static as run
COPY --from=build /go/bin/iap_test /iap_test
CMD ["/iap_test"]

