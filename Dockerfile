FROM golang:1.22-alpine as base

# Development
####################
FROM base as dev

WORKDIR /opt/app/api

RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./
RUN go mod download

CMD ["air"]

# Build
####################
FROM base as built

WORKDIR /go/app/api
COPY . .

ENV CGO_ENABLED=0

RUN go get -d -v ./...
RUN go build -o /tmp/api ./*.go

# Production
####################
FROM busybox

COPY --from=built /tmp/api /usr/bin/api
CMD ["api", "start"]
