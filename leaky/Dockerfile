FROM golang:1.10-alpine3.7 as builder
ADD main.go /go/src/app/main.go
RUN go install app

FROM alpine:3.7
RUN apk update && apk add curl
HEALTHCHECK --interval=5s CMD ["curl", "-f", "http://localhost/health"]
COPY --from=builder /go/bin/app app
CMD ["./app"]