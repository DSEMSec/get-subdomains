FROM golang:1.16.3-alpine3.13 

RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src
COPY . .

ENV GO111MODULE=auto
RUN go mod tidy

RUN CGO_ENABLED=0 go build -o main -v src/main.go

ENTRYPOINT /go/src/main