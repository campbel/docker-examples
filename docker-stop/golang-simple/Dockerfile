FROM golang:1.9.2-alpine3.6
ADD main.go /go/src/app/main.go
RUN go install app
CMD ["app"]