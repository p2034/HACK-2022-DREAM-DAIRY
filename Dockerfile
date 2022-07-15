FROM golang:alpine as builder

ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

RUN apk add --no-cache make

WORKDIR /src/app
COPY . /src/app

RUN go mod download
RUN go mod vendor

RUN go build -o bin/application .

FROM alpine:latest as runner

COPY --from=builder /src/app/bin .

ENTRYPOINT [ "./application" ]