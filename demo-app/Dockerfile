FROM golang:latest AS builder

COPY . $GOPATH/src/github.com/hengkyawijaya/dockerize/you

WORKDIR $GOPATH/src/github.com/hengkyawijaya/dockerize/you

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

RUN dep ensure -v --vendor-only

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./build/main ./*.go

FROM ubuntu:18.04

RUN DD_API_KEY=03ed27b263661e934bad7a3ce90a89ab bash -c "$(curl -L https://raw.githubusercontent.com/DataDog/datadog-agent/master/cmd/agent/install_script.sh)"

COPY --from=builder go/src/github.com/hengkyawijaya/dockerize/you/build /build

WORKDIR /build

CMD ["./main"]

EXPOSE 2113
