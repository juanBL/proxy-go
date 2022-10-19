FROM golang:alpine AS build

RUN apk add --update git
RUN mkdir /build
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/zenrows-proxy cmd/api/main.go

# Building image with the binary
FROM scratch
COPY --from=build /go/bin/zenrows-proxy /go/bin/zenrows-proxy
ENTRYPOINT ["/go/bin/zenrows-proxy"]
