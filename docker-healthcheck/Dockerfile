FROM golang:1.9.2-alpine3.6
# Add curl to use in the healthcheck
RUN apk update && apk add curl
ADD main.go /go/src/app/main.go
RUN go install app
# Check the health of the app by curling the /health endpoint
HEALTHCHECK --interval=10s --timeout=5s --start-period=5s --retries=3 CMD curl -f http://localhost/health || exit 1
CMD ["app"]