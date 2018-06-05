FROM golang:1.10-alpine3.7

# Needed for health check
RUN apk update && apk add curl

ADD main.go /go/src/app/main.go
RUN go install app
# HEALTHCHECK --interval=5s --timeout=5s --start-period=10s --retries=3 CMD ["curl", "-f", "http://localhost"]
CMD ["app"]
