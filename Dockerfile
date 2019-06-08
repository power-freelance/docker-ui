# Build web app
FROM node:8.11.3 as web-builder

WORKDIR /web
ADD web .
RUN yarn \
 && yarn build

# Build go backend
FROM golang:alpine as go-builder

# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

WORKDIR $GOPATH/src/github.com/power-freelance/docker-ui

COPY go.mod .
COPY go.sum .

ENV GO111MODULE=on
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/docker-ui ./cmd/docker-ui/main.go

# Compose vue + go and run
FROM scratch

COPY --from=go-builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=go-builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=go-builder /go/bin/docker-ui /go/bin/docker-ui

COPY --from=web-builder /web/dist /var/www

ENTRYPOINT ["/go/bin/docker-ui", "--staticRoot", "/var/www"]
