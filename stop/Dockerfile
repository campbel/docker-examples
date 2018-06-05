FROM golang:1.10-alpine3.7
ADD main.go /go/src/app/main.go
RUN go install app
# STOPSIGNAL SIGTERM - optionally specify the stop signal
# CMD ["app"] - is the proper exec form to pass os signal to container process, otherwise eaten by shell
CMD app 

