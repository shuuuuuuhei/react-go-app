FROM golang:1.17-alpine as dev
RUN apk update
RUN apk add alpine-sdk
COPY ./server/main.go /go/src/
COPY ./server/go.mod /go/src/
WORKDIR /go/src/
RUN go build -o main .
CMD [ "/go/src/main" ]